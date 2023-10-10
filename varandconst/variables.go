package varandconst

import (
	"fmt"
)

func Variables(){
	// FUNCIÓN 1
	dog :=  "🐶"
	ImprimirPerrito(dog);

	// FUNCIÓN 2 
	var(
		cantidad uint = 1
		precio uint = 200
		stock uint = 3
		producto = "hueso"
		paseo = "collar"
		juegoPerro1 = "pollo chillón"
		juegoGato1 = "puntero"
		juegoGato2 = "ratón"
		juegoPerro2 = "oreja de chancho"
	)
	//! CAMBIAR VALOR DE VARIABLE USANDO := EJ: juegoGato1, juegoPerro2 := "plumitas","oreja de chancho" - VER EXPLICACIÓN AL FINAL DEL CODIGO

	PetShop(cantidad, precio, stock, producto,juegoPerro1, paseo, juegoGato1, juegoGato2, juegoPerro2)

} 

// FUNCIÓN 1 
func ImprimirPerrito(dog string){
	fmt.Println(dog)
}

// FUNCIÓN 2 
func PetShop(precio, cantidad , stock uint, producto, juegoPerro1, paseo, juegoGato1, juegoGato2, juegoPerro2 string){
	// cambio el valor de la variable
	juegoGato1 = "plumitas"

	fmt.Println(precio,cantidad, stock, producto, juegoPerro1, paseo, juegoGato1, juegoGato2, juegoPerro2)	
}

/* 
* Diferentes formas de declarar variables :

var producto string = "hueso"

var paseo = "collar"

juegoPerro := "pollo chillón"

juegoGato1, juegoGato2 := "puntero","ratón"

& Si quiero cambiar el valor de una variable y esta está sola no puedo usar := , debo usar = . PERO  si al cambiar la variable declaro una nueva variable, sí me va a permitir. me va a saltar un error amarillo, pero no pasa nada. me esta diciendo que esa variable no se esta usando, como que no se da cuenta que la esta modificando abajo. pero me permite correr el código de todas maneras.
juegoGato1, juegoPerro2 := "plumitas","oreja de chancho"



var(
	cantidad = 1
	precio = 200
	stock = 3
) */
