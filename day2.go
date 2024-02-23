package main

import "fmt"

func main() {
	var someText string                            // ::: Assign type
	var someOtherText = "This is a string"         // ::: Type is inferred from value
	someOtherOtherText := "This is also a string!" // ::: Type is inferred from value
	var someInteger int                            // ::: Integer default = 0

	fmt.Println(someText)
	fmt.Println(someOtherText)
	fmt.Println(someOtherOtherText)
	fmt.Println(someInteger)
}
