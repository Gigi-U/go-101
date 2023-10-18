package interfaces

import "fmt"

type ListaHeterogenea struct {
	Data []interface{}
}

func EmptyInterface() {
	l := ListaHeterogenea{}
	l.Data = append(l.Data, 1)
	l.Data = append(l.Data, "hola")
	l.Data = append(l.Data, true)

	fmt.Printf("%v\n", l.Data)

}

/* Interfaces vacías
Son aquellas interfaces que no tienen métodos declarados. Se usan de "comodín". Es decir, almacenar valores que sean de un tipo de datos desconocido, o que pueda variar dependiendo el flujo del programa.

var miVariable interface{} */