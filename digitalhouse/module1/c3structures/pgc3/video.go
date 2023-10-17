package pgc3

import "fmt"

type Persona struct {
	Nombre string
	Edad   int
}

func VideoC3() {
	p := Persona{
		Nombre: "Gabriel",
		Edad:   22,
	}

	p.Descripcion()
}

func (p *Persona) Descripcion() {
	fmt.Printf("%s tiene %d a~nos de edad\n", p.Nombre, p.Edad)
}
