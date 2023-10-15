package examples

import "fmt"

//? EJEMPLO PG-

type Persona struct {
	Nombre    string
	Genero    string
	Edad      int
	Profesion string
	Peso      float64
	Gustos    Preferencias
}
type Preferencias struct {
	Comidas   string
	Peliculas string
	Series    string
	Animes    string
	Deportes  string
}
func EstructuraPersona(){
	p1 := Persona{"Celeste", "Mujer", 34, "Ingeniera", 65.5, Preferencias{"pollo", "Titanic", "", "", ""}}
	fmt.Println("----------Creando p1")
	fmt.Println(p1)

	p2 := Persona{
		Nombre:"Nahuel",
		Genero:"Hombre",
		Edad:30,
		Profesion:"Ingeniero",
		Peso:77,
		Gustos: Preferencias{
			Comidas:"asado, pollo",
			Peliculas:"Coco",
			Animes:"Shingeki no Kyojin",
		},
	}

	// ACCEDER A EDAD
	fmt.Println("----------Accediendo a Edad p2")
	fmt.Println(p2.Edad)

	// MODIFICAR EDAD
	fmt.Println("----------Modificando Edad p2")
	p2.Edad = 33
	
	var p3 Persona
	p3.Nombre = "Ulises"
	p3.Edad = 15
	p3.Gustos = Preferencias{Comidas: "verduras", Peliculas: "Entrenando a mi dragón"}
	
	fmt.Println("----------Accediendo a Gustos de p2")
	fmt.Println(p2.Gustos.Animes)

	fmt.Println("----------Agregando Gusto a p2")
	p2.Gustos.Deportes = "fútbol"

	
}

