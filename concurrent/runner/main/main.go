package main

import (
	"time"
	"log"
	"github.com/tengrommel/go_game/concurrent/runner"
	"os"
)

const timeout = 3 * time.Second

func main() {
	log.Println("Starting work.")
	r := runner.New(timeout)
	r.Add(createTask(), createTask(), createTask())

	if err := r.Start(); err != nil{
		switch err {
		case runner.ErrTimeout:
			log.Println("超时")
			os.Exit(1)
		case runner.ErrInterrupt:
			log.Println("终端")
			os.Exit(2)
		}
	}
	log.Println("程序退出.")
}

func createTask() func(int) {
	return func(id int) {
		log.Printf("程序 任务id #%d.", id)
		time.Sleep(time.Duration(id) * time.Second)
	}
}


