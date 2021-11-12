package main

import (
	pk "curso_golang_platzi/src/mypackage"
	"fmt"
)

func main() {
	var myPc pk.Pc
	myPc.Brand = "Lenovo"
	myPc.Model = "ThinkPad X1 Carbon"
	myPc.Ram = 8
	myPc.Disk = 500

	// crear una variable booleana para saber si quiere mas ram

	// preguntar por la ram
	fmt.Println("Cuantas GB de ram quiere?")
	fmt.Scanf("%d", &myPc.Ram)
	// preguntar por la disk
	fmt.Println("Cuantos GB de disco quiere?")
	fmt.Scanf("%d", &myPc.Disk)
	// preguntar por la marca
	fmt.Println("Que marca quiere?")
	fmt.Scanf("%s", &myPc.Brand)
	// preguntar por el modelo
	fmt.Println("Que modelo quiere?")
	fmt.Scanf("%s", &myPc.Model)
	// preguntar por la marca
	fmt.Println("Que marca quiere?")
	fmt.Scanf("%s", &myPc.Brand)

}
