package pgc4

import "fmt"

func PointersGo() {
	// DIFERENTES FORMAS DE DECLARAR UN PUNTERO
	// 1
	var p1 *int
	// 2
	var p2 = new(int)
	// 3
	var v int
	p3 := &v
	fmt.Println(p1,p2,p3)


	var x int = 19
	fmt.Println("la dirección de memoria de x es: ", &x)
	var p *int
	p = &x
	fmt.Printf("El puntero p referencia a la dirección de memoria: %v \n",p)
	fmt.Printf("al desreferenciar el puntero p obtengo el valor: %d \n",*p)



/* PUNTEROS: - acceder y modificar variables de forma CONTROLADA
- Tipo de dato cuyo valor es la dirección de memoria de otra variable
lleva * antes del tipo de dato ej. var p *int
* = operador de desreferenciación
desreferenciar un puntero es obtener el valor almacenado en la dirección de memoria
& = operador de dirección

muchas veces necesitamos una variable para una función y que esta cambie su valor. Al pasarle la dirección de memoria, los cambios se hacen sobre ESA dirección y persistiran en cualquier parte del programa.

Si en vez de enviar el puntero, enviasemos la variable,  GO generaría una copia exclusiva en el scope de nuestra función. Con el puntero puedo cambiar el valor de la dirección de memoria

- PUNTERO: VARIABLE QUE ALMACENA DIRECCIÓN DE MEMORIA
*/

//? EJEMPLO 

	var g int = 50
	// La función Incrementar recibe un puntero
	// utilizamos el operador de dirección “&”
	// para obtener la dirección de memoria de v
	Incrementar(&g)
	fmt.Println("El valor de v ahora vale: ", g)

}

// La función incrementar recibe un puntero de tipo entero
func Incrementar(g *int) {
	// Desreferenciamos la variable v con el operador “*” 
	// para obtener su valor e incrementarlo en 1
	*g++
 }
