package structures

import "fmt"

// En Go no existe el concepto de Herencia, en su lugar existe lo que se llama Embedding Structs. Es una COMPOSICIÓN que usa la estructura PADRE como CAMPO en las estructuras HIJAS

type Vehiculo struct {
	km     float64
	tiempo float64
}

func (v Vehiculo) detalle() {
	fmt.Printf("km:\t%f\ntiempo:\t%f\n", v.km, v.tiempo)
}
// HIJO AUTO
type Auto struct {
	v Vehiculo
}

func (a *Auto) Correr(minutos int) {
	a.v.tiempo = float64(minutos) / 60
	a.v.km = a.v.tiempo * 100
}

func (a *Auto) Detalle() {
	fmt.Println("\nV:\tAuto")
	a.v.detalle()
}
// HIJA MOTO
type Moto struct {
	v Vehiculo
}

func (m *Moto) Correr(minutos int) {
	m.v.tiempo = float64(minutos) / 60
	m.v.km = m.v.tiempo * 80
}
 
func (m *Moto) Detalle() {
	fmt.Println("\nV:\tMoto")
	m.v.detalle()
}

func Rodados(){
	auto := Auto{}
	auto.Correr(360)
	auto.Detalle()

	moto := Moto{}
	moto.Correr(360)
	moto.Detalle()
}