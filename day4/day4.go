package main

import (
	"fmt"
	"math"
)

func main() {
	fmt.Println(isLeap(2004))
}

func isLeap(year float64) bool {
	if math.Mod(year, 4) == 0 {
		if math.Mod(year, 100) == 0 {
			if math.Mod(year, 400) == 0 {
				return true
			}
			return false
		}
		return true
	}
	return false
}
