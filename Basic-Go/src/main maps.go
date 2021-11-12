package main

import (
	"fmt"
)

func main() {
	m := make(map[string]int)
	m["Jose"] = 14
	m["Maria"] = 15
	m["Jose"] = 16

	fmt.Println(m)

	// Recorer un map

	for k, v := range m {
		fmt.Println(k, v)
	}

	// encontrar un valor en un map
	value, ok := m["Jose"]
	fmt.Println(value, ok)
}
