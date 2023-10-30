package pointers
import "fmt"


func Pointers(){
	// FUNCIÓN 1
	dog :=  "🐶"
	var puntero *string
	puntero = &dog	
	*puntero = "🐕" // le cambio el valor a dog desde el puntero

	fmt.Printf("Tipo: %T, Valor: %s, Dirección: %v\n", dog, dog, &dog)
	fmt.Printf("Tipo: %T, Valor: %v, Desreferenciación: %s\n", puntero, puntero, *puntero)

	/* Se imprime: 
	Tipo: string, Valor: 🐕, Dirección: 0xc000028070
	Tipo: *string, Valor: 0xc000028070, Desreferenciación: 🐕*/
} 

// Las variables son espacios de memoria que almacenan un valor. los punteros son la dirección en memoria de esa variable. recordar agregar el & a la variable
// vamos a crear una variable puntero que apunte a la dirección de la variable dog