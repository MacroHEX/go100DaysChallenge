package main

import (
	"fmt"
	"time"
)

func main() {
	go thread(1)
	go thread(2)
	_, err := fmt.Scanln()
	if err != nil {
		return
	}
}

func thread(id int) {
	for i := 1; i <= 10; i++ {
		fmt.Println("Thread", id, ":", i)
		// ::: Without the delay, the first goroutine will finish executing even before the second one starts
		time.Sleep(100 * time.Millisecond)
	}
}
