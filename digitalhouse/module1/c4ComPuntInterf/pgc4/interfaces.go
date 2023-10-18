package pgc4

import (
	"fmt"
	"math"
)

func InterfacesGo() {
	fmt.Println("-------Ejemplo #1-------")
	ci := circle{radius: 5}
	detailsC(ci)

	fmt.Println("-------Ejemplo #2-------")
	cir := newCircle(2)
	fmt.Println(cir.area())
	fmt.Println(cir.perim())

	fmt.Println("-------Ejemplo #3-------")
	r := newGeometry(rectType, 2, 3)
	fmt.Println(r.area())
	fmt.Println(r.perim())
	c := newGeometry(circleType, 2)
	fmt.Println(c.area())
    fmt.Println(c.perim())
 
}
//&Ejemplo #1
type circle struct {
	radius float64
}

func (c circle) area() float64 {
	return math.Pi * c.radius * c.radius
}

func (c circle) perim() float64 {
	return 2 * math.Pi * c.radius
}

func detailsC(c circle) {
	fmt.Println(c)
	fmt.Println(c.area())
	fmt.Println(c.perim())
} 

//&Ejemplo #2 
/* ¿Qué pasará si queremos generar más figuras geométricas utilizando nuestra función details?
	1 definimos nuestra interfaz geometry que contiene dos métodos que adoptarán nuestros objetos: */

type geometry interface {
	area() float64
	perim() float64
}
//	2 Generamos otro objeto geométrico. En este caso, un rectángulo que —lógicamente— tenga los mismos métodos:

type rect struct {
	width, height float64
}
func (r rect) area() float64 {
	return r.width * r.height
}
func (r rect) perim() float64 {
	return 2*r.width + 2*r.height
}
func newCircle(values float64) circle {
	return circle{radius: values}
}
//	3 Modificaremos nuestra función details para que, en lugar de recibir un círculo, reciba una figura geométrica:

func details(g geometry) {
	fmt.Println(g)
	fmt.Println(g.area())
	fmt.Println(g.perim())
}

//&Ejemplo #3
//¿Qué pasará si queremos reutilizar nuestra función para poder implementar varias figuras geométricas?  En este caso tendremos que crear una función que retorne una interfaz que pueda implementar todos nuestros objetos geométricos

const (
	rectType   = "RECT"
	circleType = "CIRCLE"
)

 func newGeometry(geoType string, values ...float64) geometry {
	switch geoType {
	case rectType:
		return rect{width: values[0], height: values[1]}
	case circleType:
		return circle{radius: values[0]}
	}
	return nil
}
/*
INTERFACES:
- Una interfaz es una forma de definir métodos que deben ser utilizados, pero sin definirlos.
- Las interfaces son utilizadas para brindar modularidad al lenguaje.

*/