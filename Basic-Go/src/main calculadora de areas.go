package main

import (
	pki "curso_golang_platzi/src/packageinterfaces"
	"fmt"
)

func main() {

	// Preguntar al ususario a que figura quiere calcular su area
	fmt.Println("Actualmente puedo calcular el area de las siguientes figuras")
	fmt.Println("Cuadrado, Rectangulo, Circulo, Triangulo")
	fmt.Println("Ingrese el nombre de la figura:")
	var figura string
	fmt.Scanf("%s", &figura)

	if figura == "Cuadrado" {
		fmt.Println("Ingrese el valor de la base:")
		var base float64
		fmt.Scanf("%f", &base)
		myCuadrado := pki.Cudrado{Base: base}
		pki.CalculaArea(myCuadrado)
	} else if figura == "Rectangulo" {
		fmt.Println("Ingrese el valor de la base:")
		var base float64
		fmt.Scanf("%f", &base)
		fmt.Println("Ingrese el valor de la altura:")
		var altura float64
		fmt.Scanf("%f", &altura)
		myRectangulo := pki.Rectangulo{Base: base, Altura: altura}
		pki.CalculaArea(myRectangulo)
	} else if figura == "Circulo" {
		fmt.Println("Ingrese el valor del radio:")
		var radio float64
		fmt.Scanf("%f", &radio)
		myCirculo := pki.Circulo{Radio: radio}
		pki.CalculaArea(myCirculo)
	} else if figura == "Triangulo" {
		fmt.Println("Ingrese el valor de la base:")
		var base float64
		fmt.Scanf("%f", &base)
		fmt.Println("Ingrese el valor de la altura:")
		var altura float64
		fmt.Scanf("%f", &altura)
		myTriangulo := pki.Triangulo{Base: base, Altura: altura}
		pki.CalculaArea(myTriangulo)
	} else {
		fmt.Println("Figura no valida porfavor vuelva a intentarlo")
		fmt.Println("**********************************************")
		main()
	}
}
