package main

import (
	"os"
	"log"
	"io"
)

func main() {
	//createFile()
	deleteFile()
}

func createFile() {
	f, err := os.Create("file.txt")
	if err != nil{
		log.Fatalln(err)
	}
	defer f.Close()
	log.Printf("Created %s\n", f.Name())
}

func deleteFile()  {
	err := os.Remove("file.txt")
	if err != nil{
		log.Fatalln(err)
	}
	log.Println("Deleted file.txt")
}

func checkExistence()  {
	fi, err := os.Stat("file.txt")
	if err != nil{
		if os.IsNotExist(err) {
			log.Fatalln("Does not exist.")
		}
	}
	log.Printf("Exists, last modified %v\n", fi.ModTime())
}

func renameFile()  {
	f, _ := os.Create("file.txt")
	err := os.Rename(f.Name(), "rename.txt")
	if err != nil{
		log.Fatalln(err)
	}
}

func copyFile()  {
	of, err := os.Open("proverbs.txt")
	if err != nil{
		log.Fatalln(err)
	}
	defer of.Close()

	nf, err := os.Create("copy.txt")
	if err != nil{
		log.Fatalln(err)
	}
	defer nf.Close()

	bw, err := io.Copy(nf, of)
	if err != nil{
		log.Fatalln(err)
	}
	log.Printf("Bytes written: %d\n", bw)
	if err := nf.Sync(); err != nil{
		log.Fatalln(err)
	}
	log.Println("Done")
}