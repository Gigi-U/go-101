package pgc4

import (
	"bufio"
	"encoding/json"
	"fmt"
	"os"
	"strings"
	"time"
)

/*
Employee

Una empresa necesita realizar una buena gestión de sus empleados,
para esto realizaremos un pequeño programa nos ayudará a gestionar correctamente
dichos empleados.
Los objetivos son:

	- Crear una estructura Person con los campos ID, Name, DateOfBirth.
	- Crear una estructura Employee con los campos: ID, Position y una
	composición con la estructura Person.

	- Realizar un método a la estructura Employee que se llame PrintEmployee(),
	lo que hará es realizar la impresión de los campos de un empleado.

	- Instanciar en la función main() tanto una Person como un Employee cargando
	sus respectivos campos y, por último, ejecutar el método PrintEmployee().

Si logramos realizar este pequeño programa, pudimos ayudar a la empresa
a solucionar la gestión de los empleados.

Persona es una estructura que descibre el modelo de una persona.
*/

type Person struct{
	ID			int		`json:"id"`
	Name 		string	`json:"name"`
	//Age 		int		`json:"age"`
	DOB time.Time		`json:"dob"`
}
type Employee struct{
	ID			int 	`json:"id"`
	Position 	string	`json:"position"`
	Person Person		`json:"person"`
}
var reader = bufio.NewReader(os.Stdin)

var position, name, dob string
var personId int

// Generando Id autoincremental
var currentID int
func getNextID() int {
    currentID++
    return currentID
}
/* //NO SE PIDE -  Función para calcular la edad en base a los datos ingresados por consola
func CalculateAge(name, dob string) (int, error) {
	dobTime, err := time.Parse("02/01/2006", dob)
	if err != nil {
		return 0, err
	}

	age := time.Since(dobTime).Hours() / 24 / 365.25
	return int(age), nil
} */

// Function to print Employee with details
func(e Employee)PrintEmployee(){
	fmt.Printf("\nID: %d\n", e.ID)
	fmt.Printf("\tPosition: %s \n", e.Position)
	fmt.Printf("\tID: %d\n", e.Person.ID)
	fmt.Printf("\tName: %s\n", e.Person.Name)
	fmt.Printf("\tDate of Birth: %s\n", dob)
} 


func PersonnelManagement() {

	id := getNextID()

	// ingreso de datos por consola:
	fmt.Print("Ingresar puesto en la empresa: ")
	// Con BUFIO  permite ingresar strings compuestos. Scanln no los reconoce y rompe el ingreso de datos
	position, _ = reader.ReadString('\n')
    position = strings.TrimSpace(position)

	fmt.Print("Ingresa documento de identidad sin puntos: ")
	fmt.Scanln(&personId) // Lee la fecha de nacimiento como una cadena

	// ingreso de datos por consola:
	fmt.Print("Ingresar nombre del Empleado: ")
	name, _ = reader.ReadString('\n')
    name = strings.TrimSpace(name)

	fmt.Print("Ingresa fecha de nacimiento en formato (dd/mm/yyyy): ")
	fmt.Scanln(&dob) // Lee la fecha de nacimiento como una cadena


	
	/* age, err := CalculateAge(name, dob)
	if err != nil {
		fmt.Println("Error al analizar la fecha de nacimiento:", err)
		return
	} */
	//  esto se asegura de que si ingreso mal la fecha de nacimiento me notifica el error. FUNCIONA CON DÍAS Y MESES no con años posteriores al actual
		parsedDOB, err := time.Parse("02/01/2006", dob)
	if err != nil {
		fmt.Printf("Error al parsear la fecha de nacimiento: %v\n", err)
		return
	}
   // s := Person{ID: id, Name: name, DOB: time.Time{}, Age: age}
    s := Employee{ID: id, Position: position, Person: Person{ ID: personId, Name: name, DOB:parsedDOB}}

   // fmt.Printf("{%d %s %d %s}\n", s.ID, s.Name, s.Age, dob)
    //fmt.Printf("%d %s %s\n", s.ID, s.Name, dob)
	s.PrintEmployee()	
	//Structure Tags
	miJSON, err := json.Marshal(s)
	fmt.Println("------------Imprimiendo JSON")
	fmt.Println(string(miJSON))
	fmt.Println("------------Error")
	fmt.Println(err)
}
	

/* 
*TERMINAL:

Ingresar puesto en la empresa: Desarrolladora Backend
Ingresa documento de identidad sin puntos: 27111111
Ingresar nombre del Empleado: Sonia Gisela Urriza y Spreafichi
Ingresa fecha de nacimiento en formato (dd/mm/yyyy): 26/02/1980

ID: 1
        Position: Desarrollador Backend
        ID: 27111111 
        Name: Sonia Gisela Urriza y Spreafichi
        Date of Birth: 26/02/1980

------------Imprimiendo JSON
{"id":1,"position":"Desarrolladora Backend","person":{"id":27111111,"name":"Sonia Gisela Urriza y Spreafichi","dob":"1980-02-26T00:00:00Z"}}	
!El dob queda horrible en el JSON. aun no logré cambiarlo	
*/
	
	
	
	