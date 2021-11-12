package main

import (
	"fmt"
)

func message(text string, c chan string) {
	c <- text
}

func main() {
	c := make(chan string, 2)
	c <- "uno"
	c <- "dos"

	fmt.Println(len(c), cap(c))

	// Close cierra el canal y no permite más envíos
	close(c)

	// c <- "tres"

	// Range recorre el canal y devuelve los valores
	for v := range c {
		fmt.Println(v)
	}

	// Select es una forma de manejar canales
	email1 := make(chan string)
	email2 := make(chan string)

	go message("Hola", email1)
	go message("Adios", email2)
	for i := 0; i < 2; i++ {
		select {
		case msg1 := <-email1:
			fmt.Println("Email de variable1", msg1)
		case msg2 := <-email2:
			fmt.Println("Email de la segunda variable", msg2)
		}
	}
}
