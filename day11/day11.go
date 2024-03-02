package main

import (
	"log"
	"os"
)

func main() {
	_, err := Open("fileThatDoesExist.txt")
	if err != nil {
		return
	}
}

func Open(name string) (file *os.File, err error) {
	f, err := os.Open(name)
	if err != nil {
		log.Fatal(err)
	}
	return f, err
}
