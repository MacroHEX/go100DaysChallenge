package main

import (
	"fmt"
)

func someGenerator(msg string) <-chan string {
	ch := make(chan string)
	go func() {
		for {
			ch <- fmt.Sprintf("%s", msg)
		}
	}()
	return ch
}

func main() {
	cha := someGenerator("this is a msg to trigger the channel")
	chb := someGenerator("this is another msg to trigger the channel")
	for i := 0; i < 2; i++ {
		fmt.Println(<-cha)
		fmt.Println(<-chb)
		fmt.Println("---")
	}
}
