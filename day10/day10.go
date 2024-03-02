package main

import "fmt"

type Counter interface {
	Add(increment int)
	Value() int
}

type Stats struct {
	value int
}

func (s *Stats) Add(v int) {
	s.value += v
}

func (s *Stats) Value() int {
	return s.value
}

func main() {
	stats := Stats{}
	fmt.Println(stats.Value())
	stats.Add(1)
	fmt.Println(stats.Value())
	stats.Add(19)
	fmt.Println(stats.Value())
}
