package main

import (
	"fmt"
)

// Clases en Go son llamados structs

type car struct {
	brand string
	year  int
}

func main() {
	myCar := car{brand: "Ford", year: 2020}
	fmt.Println(myCar)

	// Otra forma de declarar un struct
	var otherCar car
	otherCar.brand = "Toyota"
	fmt.Println(otherCar)

}
