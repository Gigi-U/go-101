package pgc4
import "fmt"

type Autor struct {
	Nombre   string
	Apellido string
}

func (a *Autor) nombreCompleto() string {
	return fmt.Sprintf("%s %s", a.Nombre, a.Apellido)
}

type Libro struct {
	Titulo    string
	Contenido string
	Autor     Autor
}

func (b *Libro) informacion() {
	fmt.Println("Titulo: ", b.Titulo)
	fmt.Println("Contenido: ", b.Contenido)
	fmt.Println("Autor: ", b.Autor.nombreCompleto())

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

	//Ejecutar con go run CTD-EB3-C05-Codificando-juntos.go
}

/*
- Las estructuras se componen de otras estructura, y x esto pueden usar sus funcionalidades ( vendría a ser como la herencia de JAVA)
*/