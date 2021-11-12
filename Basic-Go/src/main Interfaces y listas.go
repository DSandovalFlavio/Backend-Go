package main

import (
	"fmt"
)

type figuras2D interface {
	area() float64
}

type cudrado struct {
	base float64
}

type rectangulo struct {
	base   float64
	altura float64
}

func (c cudrado) area() float64 {
	return c.base * c.base
}

func (r rectangulo) area() float64 {
	return r.base * r.altura
}

func calculaArea(f figuras2D) {
	fmt.Println("Area:", f.area())
}

func main() {
	myCuadrado := cudrado{base: 5}
	myRectangulo := rectangulo{base: 5, altura: 10}

	calculaArea(myCuadrado)
	calculaArea(myRectangulo)

	// lista de interfaces
	myInterfaces := []interface{}{"Hola", 1, 2.2}
	fmt.Println(myInterfaces...)
}
