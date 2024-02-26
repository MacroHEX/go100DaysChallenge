package main

import "fmt"

type car struct {
	model string
	brand string
	year  int
}

func newCar(model string, brand string, year int) *car {
	car := car{model: model, brand: brand, year: year}
	return &car // ::: An & prefix yields a pointer to the struct.
}

func main() {
	c := car{"X70", "Jetour", 2024}

	fmt.Println(c)

	fmt.Println(newCar("X70 Plus", "Jetour", 2024))
}
