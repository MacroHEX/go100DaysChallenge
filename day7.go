package main

import "fmt"

func main() {
	var m map[string]int

	m = make(map[string]int)

	m["test"] = 55

	fmt.Println(m["test"])

}
