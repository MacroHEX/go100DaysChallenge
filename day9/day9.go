package main

import "fmt"

type car struct {
	brand, model string
}

// ::: Having both value and pointer receivers is not recommended, this is for learning purposes only
//
// ::: This will create a pointer so the values will be mutable
func (c *car) showDetails() string {
	return "Brand: " + c.brand + "\nModel: " + c.model
}

// ::: This will not create a pointer so the it wont mutate
func (c car) setBrand(brand string) {
	c.brand = brand
}

func main() {
	c := car{brand: "JETOUR", model: "X70"}

	fmt.Println(c.showDetails())

	c.setBrand("Something") // This should change the brand, but wont as is a value receiver

	fmt.Println(c.showDetails()) // Brand will still be 'Jetour'
}
