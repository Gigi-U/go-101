package pgc4

import (
	"fmt"
)

/* Objetivo
El objetivo de esta guía práctica es que podamos afianzar los conceptos sobre estructuras e interfaces. Para esto vamos a plantear un ejercicio simple e incremental, que nos permitirá repasar los temas que estudiamos.
Consigna
Una empresa de redes sociales requiere implementar una estructura usuarios con funciones que vayan agregando información a la misma.
Para optimizar y ahorrar memoria requieren que la estructura usuarios ocupe el mismo lugar en memoria para el main del programa y para las funciones.

La estructura debe tener los campos: nombre, apellido, edad, correo y contraseña. Y deben implementarse las funciones:

cambiarNombre: permite cambiar el nombre y apellido.
cambiarEdad: permite cambiar la edad.
cambiarCorreo: permite cambiar el correo.
cambiarContraseña: permite cambiar la contraseña. 

! NO ENTIENDO SI TENGO QUE HACER UNA INTERFAZ , O SOLO CAMBIAR LOS DATOS DE LA ESTRUCTURA CON PUNTERO.


*/

// Usuario Estructura que describe objeto Usuario
type Usuario struct{
	Nombre 		string 	
	Apellido 	string	
	Edad 		int 	
	Correo		string 	
	Contrasena 	string 	
}


// or main()  
func InfoUsuario(){

	u1 := Usuario{
		Nombre:   "Gisela",
		Apellido: "Urriza",
		Edad: 43,
		Correo: "gu@gmail.com",
		Contrasena: "123456",
	}

	u2:=&u1
	// Impresión usuario1
	u1.InformacionUsuario()
	
	// Impresión usuario2
	*u2 = newUser("Cosme","Fulanito",70,"cf@gmail.com","654321")
	u2.InformacionUsuario()


}

func (u Usuario) InformacionUsuario() {
	fmt.Printf(" Usuario: %s %s %d %s %s , Dirección: %v \n", u.Nombre, u.Apellido, u.Edad, u.Correo, u.Contrasena, &u.Nombre)

}

type user interface {
	cambiarNombre() string
	cambiarApellido() string
	cambiarEdad() int
	cambiarCorreo() string
	cambiarContrasena()  string
}
func newUser(nombre string, apellido string, edad int , correo string, contraseña string) Usuario {
	return Usuario{Nombre: nombre, Apellido: apellido, Edad: edad, Correo: correo, Contrasena: contraseña}
}

func (u Usuario) cambiarNombre( nombre string){
	u.Nombre = nombre
}
func (u Usuario) cambiarApellido( apellido string){
	u.Apellido = apellido
}
func (u Usuario)cambiarEdad(edad int){
	u.Edad = edad
}
func (u Usuario) cambiarCorreo(correo string){
	u.Correo = correo
}
func (u Usuario) cambiarContrasena(contrasena string){
	u.Contrasena = contrasena
}





// sin tengo 2 metodos con el mismo nombre y con la misma firma, correrlo.
//en el caso de las funciones no se puede

//!return fmt.Sprintf("hola %s", p.Nombre)