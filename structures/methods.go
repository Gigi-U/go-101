package structures

import (
	"fmt"
	"math"
)

// LOS MÉTODOS , A DIFERENCIA DE LAS FUNCIONES, presentan argumentos de receptor. es decir

type Circulo struct {
	radio float64
}

// func (v MiEstructura) metodo(){ }
func (c Circulo) area() float64 {
	area:=math.Pi * c.radio * c.radio
	return area
}

func (c Circulo) perim() float64 {
	perimetro:=2 * math.Pi * c.radio
	return perimetro
}

func Circ(){
	c:=Circulo{
		radio: 34.44,
	}
//Here we call the 2 methods defined for our struct.

    fmt.Println("area: ", c.area())
    fmt.Println("perim:", c.perim())

/* 	Go automatically handles conversion between values and pointers for method calls. You may want to use a pointer receiver type to avoid copying on method calls or to allow the method to mutate the receiving struct. */
	fmt.Println("------------POINTERS: ")

	cp := &c
    fmt.Println("area: ", cp.area())
    fmt.Println("perim:", cp.perim())
}
