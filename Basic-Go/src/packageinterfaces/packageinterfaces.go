package packageinterfaces

import (
	"fmt"
)

type Figuras2D interface {
	Area() float64
}

type Cudrado struct {
	Base float64
}

type Rectangulo struct {
	Base   float64
	Altura float64
}

type Circulo struct {
	Radio float64
}

type Triangulo struct {
	Base   float64
	Altura float64
}

func (c Cudrado) Area() float64 {
	return c.Base * c.Base
}

func (r Rectangulo) Area() float64 {
	return r.Base * r.Altura
}

func (c Circulo) Area() float64 {
	return 3.14 * c.Radio
}

func (t Triangulo) Area() float64 {
	return t.Base * t.Altura / 2
}

func CalculaArea(f Figuras2D) {
	fmt.Println("Area:", f.Area())
}
