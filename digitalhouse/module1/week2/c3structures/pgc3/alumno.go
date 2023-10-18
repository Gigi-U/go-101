package pgc3

import (
	"encoding/json"
	"fmt"
)

type Alumno struct {
	Nombre   string `json:"Nombre"`
	Apellido string `json:"Apellido"`
	DNI      string `json:"DNI"`
	Fecha    string `json:"Fecha"`
}

func (a Alumno) detalle() string {	
	return "Nombre: "+a.Nombre+"\n"+ 
			"Apellido: "+a.Apellido+"\n"+ 
			"DNI: "+a.DNI+"\n"+ 
			"Fecha: "+a.Fecha
}

func Alumni(){
	a:=Alumno{"Gisela","Urriza","277777777","26/02/1980"}
	fmt.Printf("%+v\n",a.detalle())

	// Con TAGS JSON Y ERROR
	miJSON, err := json.Marshal(a)
	fmt.Println("------------Imprimiendo JSON")
	fmt.Println(string(miJSON))

	fmt.Println("------------Error")
	fmt.Println(err)
}

// RESOLUCIÓN 
/* type Alumno struct {
	nombre   string
	apellido string
	dni      int
	fecha    string
}

func (a Alumno) Detalle() {
	fmt.Printf("\nDetalle del alumno:")
	fmt.Printf("\nNombre: %s \nApellido: %s \nDNI: %v \nFecha: %s\n", a.nombre, a.apellido, a.dni, a.fecha)
}

func Alumni() {
	alumno := Alumno{
		nombre:   "Carlos",
		apellido: "Molina",
		dni:      95866074,
		fecha:    "25/10/2023",
	}

	alumno2 := Alumno{
		nombre:   "Pablo",
		apellido: "Gonzalez",
		dni:      36564453,
		fecha:    "29/01/2015",
	}

	Alumno.Detalle(alumno)
	Alumno.Detalle(alumno2)

} */