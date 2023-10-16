package mtc1
import (
	"fmt"
)

/* 
Ejercicio 2 - Qué edad tiene...
Un empleado de una empresa quiere saber el nombre y edad de uno de sus empleados. 
Según el siguiente map, debemos imprimir la edad de Benjamin.
  var employees = map[string]int{"Benjamin": 20, "Nahuel": 26, "Brenda": 19, "Darío": 44, "Pedro": 30}

Por otro lado, también es necesario:
Saber cuántos de sus empleados son mayores de 21 años.

Agregar un empleado nuevo a la lista, llamado Federico que tiene 25 años.

Eliminar a Pedro del map.

 */
func Exercise2() {

	employees := map[string]int{"Benjamin": 20, "Nahuel": 26, "Brenda": 19, "Darío": 44, "Pedro": 30}

	listadoInicial(employees)
	edadBenjamin(employees)
	contadorMayores(employees)
	agregarEmpleado(employees)
	eliminarEmpleado(employees)


}

// LISTADO GENERAL

func listadoInicial(employees map[string]int){
	fmt.Println("----------1 all------------")

	for k, e  := range employees {
		fmt.Printf("Nombre: %s \t Edad: %d\n", k, e)
		}
}



// BENJAMÍN
func edadBenjamin(employees map[string]int){
	fmt.Println("----------2 find------------")

	for k, e  := range employees {
		if k == "Benjamin" {
			fmt.Println("2) Datos Empleado - Nombre: ",k,"| Edad: ",e)
		}
	}
}

// CONTANDO MAYORES DE EDAD
func contadorMayores(employees map[string]int){
	fmt.Println("----------3 count------------ ")

	suma:=0
	 for _ ,e  := range employees {
		if e >= 21 {
			suma ++
		}
    } 
	fmt.Println("3) Hay ", suma, " empleados mayores de 21 años")

}

// NUEVO EMPLEADO
func agregarEmpleado(employees map[string]int){

	fmt.Println("----------4 add------------")

	employees["Federico"] = 25
	//fmt.Println("4) Agregando a Federico al listado: ", employees)
	 for k, e  := range employees {
		fmt.Println("Nombre: ",k,"| Edad: ",e)
	} 
}

// FUERA EMPLEADO
func eliminarEmpleado(employees map[string]int){
	fmt.Println("----------5 delete------------")

	delete(employees,"Pedro")
	//fmt.Println("5) Eliminando a Pedro del listado: ", employees)
	for k, e  := range employees {
		fmt.Println("Nombre: ",k,"| Edad: ",e)
	}
}

/* RESULTADO: 

*/