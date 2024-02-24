package main

import "fmt"

func main() {
	var m map[string]int

	m = make(map[string]int)

	m["test"] = 55
	m["test2"] = 12
	m["test3"] = 24
	m["test4"] = 567

	for key, value := range m {
		fmt.Println("Key:", key, "Value:", value)
	}

	_, prs := m["test5"] // ::: value,present as value is not needed here _ is used
	fmt.Println("Present:", prs)

}
