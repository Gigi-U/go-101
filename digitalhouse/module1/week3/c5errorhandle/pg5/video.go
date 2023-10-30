package pg5

import "fmt"

type Gato struct {
	Name string
}

func main() {
	fmt.Println("lA FUNCION ESTA INICIANDO")
	Ejemplo3()
	fmt.Println("lA FUNCION ESTA FINALIZANDO")
}

func Ejemplo1() {
	animales := []string{
		"vaca",
		"perro",
		"halcon",
		"oso",
		//indice 5 no existe
	}

	fmt.Println("El animal que vuela es:", animales[5])
}

func Ejemplo2() {
	panic(true)
}

func Ejemplo3() {
	g := &Gato{"Misho"}
	fmt.Println(g)
	g = nil

	g.Hablar()
}

func (s *Gato) Hablar() {
	fmt.Print(s.Name, "El gato hace miau.")
}
