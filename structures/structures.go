package structures

import (
	"fmt"
)

func Estructuras() {

	type curso struct {
		Nombre string 
		Profesor string 
		Pais string 
	}
	// para instanciar la estructura curso:
	db:= curso{
		Nombre : "Go",
		Profesor : "Cosme Fulanito",
		Pais : "Argentina",
	}
	fmt.Printf("%+v\n", db)

	// otra forma de crear instancias -  me sirve si quiero mostrar todas las propiedades, en caso contrario debo hacerlo de la manera previa: 
	java:= curso{"Java", "Peter Parker", "USA"}
	fmt.Printf("%+v\n", java)

	// para imprimir solo una propiedad:
	fmt.Printf("%+v\n", db.Nombre)

	//& PUNTEROS A ESTRUCTURAS:
	puntero := &db
	// cambio valores con el operador de des-referenciación:
	(*puntero).Profesor = "Beto" 
	// se puede cambiar tmb así :
	puntero.Pais = "Brasil"
	fmt.Printf("%+v\n", db)
	fmt.Printf("%+v\n", puntero)

}

/* CONCEPTOS: 
Las estructuras en Go me permiten almacenar una coleccion de campos. Esto es similas a las CLASES en otros lenguajes.

!Para declarar una estructura utilizo la palabra reservada "type", el nombre de la clase y luego la palabra reservada "struct"

PG - Estructura: colección de múltiples campos que guardan datos de diferentes tipos. Cualquier objeto de la vida real que tenga propiedades, o campos, se puede representar con una estructura.
Sus caracteristicas tienen que estar unificadas en un mismo tipo de variable.const
Las estructuras me permiten "pasar" estructuras complejas a través de los diferentes MÓDULOS

Con las estructuras nos evitamos, x ejemplo, pasar 10 parámetros por una función. En lugar de 10 parámetros, pasará una estructura y podremos acceder a las propiedades de manera sencilla.
*/

