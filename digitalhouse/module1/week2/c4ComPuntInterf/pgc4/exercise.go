package pgc4

import (
	"fmt"
)

/* 
Consigna

Una empresa de redes sociales requiere implementar una estructura usuarios con funciones que vayan agregando información a la misma.

Para optimizar y ahorrar memoria requieren que la estructura usuarios ocupe el mismo lugar en memoria para el main del programa y para las funciones. (Punteros)

	- La estructura debe tener los campos: 
		nombre
		apellido
		edad
		correo 
		contraseña
	- Y deben implementarse las funciones:
		cambiarNombre: permite cambiar el nombre y apellido.
		cambiarEdad: permite cambiar la edad.
		cambiarCorreo: permite cambiar el correo.
		cambiarContraseña: permite cambiar la contraseña. 
*/

// Usuario Estructura que describe objeto Usuario
type Usuario struct{
	Nombre 		string 	
	Apellido 	string	
	Edad 		int 	
	Correo		string 	
	Contrasena 	string 	
}


func (u Usuario) InformacionUsuario() {
	fmt.Printf(
		" Usuario: %s %s %d %s %s \n", 
		u.Nombre, 
		u.Apellido, 
		u.Edad, 
		u.Correo, 
		u.Contrasena,
	)
}


func cambiarNombre(u *Usuario) string{
	u.Nombre = "Cosme"
	return u.Nombre
}
 func cambiarApellido( u *Usuario)string{
	u.Apellido = "Fulanito"
	return u.Apellido
}
func cambiarEdad(u *Usuario) int{
	u.Edad = 70
	return u.Edad
}
func cambiarCorreo(u *Usuario) string{
	u.Correo = "cf@gmail.com"
	return u.Correo
}
func cambiarContrasena(u *Usuario) string{
	u.Contrasena = "1111111"
	return u.Contrasena 
} 

// or main()  
func InfoUsuario(){
	// cargando datos en el Usuario inicial y seteando el puntero
	u := &Usuario{
		Nombre:   	"Gisela",
		Apellido: 	"Urriza",
		Edad: 		43,
		Correo: 	"gu@gmail.com",
		Contrasena: "123456",
	}
	// imprimiendo valor inicial de Usuario
	fmt.Sprintln( "el valor inicial de Usuario es: ")
	u.InformacionUsuario()

	//	LLamado a funciones de cambiod e datos
	cambiarNombre(u)
	cambiarApellido(u)
	cambiarEdad(u)
	cambiarCorreo(u)
	cambiarContrasena(u)
	// imprimiendo valor final de usuario
	fmt.Sprintln( "el valor final de Usuario es: ")
	u.InformacionUsuario()

	// Estructura de control para comprobar que el puntero apunta a la variable
	// 1) Verifica si el puntero se creó correctamente
	if u != nil {
		fmt.Println("El puntero a Usuario se creó correctamente.")

		// 2) Imprime la dirección de memoria del puntero
		fmt.Printf("Dirección del puntero: %p\n", u)

		// 3) Imprimela dirección de memoria de la variable apuntada
		fmt.Printf("Dirección de la variable apuntada: %p\n", &u.Nombre)
	} else {
		fmt.Println("El puntero a Usuario no se creó correctamente.")
	}
}
/*
*CONSOLA:
	Usuario: Gisela Urriza 43 gu@gmail.com 123456
	Usuario: Cosme Fulanito 70 cf@gmail.com 1111111
	El puntero a Usuario se creó correctamente.
	Dirección del puntero: 0xc0000ce0f0
	Dirección de la variable apuntada: 0xc0000ce0f0 
*/



// sin tengo 2 metodos con el mismo nombre y con la misma firma, correrlo.
//en el caso de las funciones no se puede

//return fmt.Sprintf("hola %s", p.Nombre)