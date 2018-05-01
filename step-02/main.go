
package main

import (
	"net/http"
	"github.com/gorilla/websocket"
	"fmt"
	"log"
)

var clients = make(map[*websocket.Conn]bool) // 注册客户端
var upgrader = websocket.Upgrader{}

type Message struct {
	Type      string `json:"type"`
	Content   string `json:"content"`
}

var broadcast = make(chan Message, 100)

func main() {
	// Create a simple file server
	fs := http.FileServer(http.Dir("."))
	http.Handle("/", fs)

	http.HandleFunc("/wss", handleConnections)
	go handleMessages()
	fmt.Println("8000:")
	panic(http.ListenAndServeTLS(":8000","cert.pem", "key.pem", nil))
	//err := http.ListenAndServe(":8080",nil)
}


func handleConnections(w http.ResponseWriter, r *http.Request) {
	// Upgrade initial GET request to a websocket
	ws, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		log.Fatal(err)
	}
	defer ws.Close()
	switch len(clients) {
	case 0:
		clients[ws] = true
	case 1:
		clients[ws] = true
	default:
		ws.WriteMessage(2,[]byte("fff the room is full"))
	}

	for {
		var msg Message
		// Read in a new message as JSON and map it to a Message object
		err := ws.ReadJSON(&msg)
		if err != nil {
			log.Printf("error: %v", err)
			delete(clients, ws)
			break
		}
		// Send the newly received message to the broadcast channel
		broadcast <- msg
	}
}

func handleMessages() {
	for {
		for msg := range broadcast{
			switch msg.Type {
			case "call":
				call(msg)
			case "ice":
				ice(msg)
			}
		}
	}
}

func ice(message Message) {
	for client := range clients {
		err := client.WriteJSON(message)
		if err != nil {
			log.Printf("error: %v", err)
			client.Close()
			delete(clients, client)
		}
	}
}

func call(message Message) {
	//将两个sdp返回前端比对后拿到不同的sdp
	for client := range clients {
		err := client.WriteJSON(message)
		if err != nil {
			log.Printf("error: %v", err)
			client.Close()
			delete(clients, client)
		}
	}
}

