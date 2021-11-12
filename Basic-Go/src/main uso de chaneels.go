package main

import (
	"fmt"
)

func say(text string, c chan<- string) {
	c <- text
}

func main() {
	c := make(chan string, 1)

	fmt.Println("Starting Go Routines")
	go say("Hello", c)

	fmt.Println(<-c)
}
