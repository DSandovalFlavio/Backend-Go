package main

import (
	"fmt"
	"strings"
)

// una funcion que compruebe si una palabra es un palindromo
func esPalindromo(palabra string) bool {
	// creo una variable para guardar la palabra
	var palabraInvertida string
	// convierto la palabra en minuscula
	palabra = strings.ToLower(palabra)
	// recorro la palabra
	for _, letra := range palabra {
		// la invertimos
		palabraInvertida = string(letra) + palabraInvertida
	}
	// comparamos la palabra con la palabra invertida
	if palabra == palabraInvertida {
		return true
	}
	return false
}

func main() {
	slice := []string{"Hola", "que", "hace"}

	for _, v := range slice {
		fmt.Println(v)
	}

	fmt.Println(esPalindromo("AmOr a ROma"))
}
