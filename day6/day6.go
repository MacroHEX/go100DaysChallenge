package main

import "fmt"

func main() {
	// ::: array is a fixed-size data structure of the same type
	myArr := [5]int{1, 2, 3, 4, 5}

	// ::: a slice can be resized
	mySlice := []int{1, 2, 3, 4, 5}

	myArr[2] = 3

	mySlice = append(mySlice, 6, 7, 8)

	fmt.Println(myArr)
	fmt.Println(mySlice)

	for i := 0; i < len(myArr); i++ {
		fmt.Println(myArr[i])
	}

	for i, v := range mySlice {
		fmt.Println("Index: ", i, " Value: ", v)
	}
}
