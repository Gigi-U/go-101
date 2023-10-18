package pgc4

import "fmt"

type Autor struct {
	Nombre string
	Apellido string
}

func (a Autor) nombreCompleto() string {
	return fmt.Sprintf("%s %s", a.Nombre, a.Apellido)
}

type Libro struct {
	Titulo string
	Contenido string
	Autor Autor // se embeben los datos de la estructura Autor en la estructura Libro. ESTO ES LA COMPOSICIÓN
}

func (l Libro) informacion() {
	fmt.Println("Titulo: ", l.Titulo)
	fmt.Println("Contenido: ", l.Contenido)
	fmt.Println("Autor: ", l.Autor.nombreCompleto()) // Con Composición hago uso de los MÉTODOS de las estructuras anidadas. desde libro llamo al metodo nombreCompleto() de Autor

}

func HerenciaGo() {
	autor := Autor{
		"Juan",
		"Lopez",
	}

	libro := Libro{
		Titulo:    "Herencia en Go.",
		Contenido: "Go posee composición en lugar de herencia.",
		Autor:     autor,
	}
	libro.informacion()
}

/*

-La composición es un MÉTODO creado para escribir segmentos de código reutilizables.

- en lugar de heredar características de una clase u objeto principal las estructuras se componen de otras estructura, y x esto pueden usar sus funcionalidades ( vendría a ser como la herencia de JAVA)


- Herencia es tomar las propiedades de una superclase y heredarlas a una clase base

- Como GO no admite clases, logra la herencia a través de la incorporación de estructuras
*/