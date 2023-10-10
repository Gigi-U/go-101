package structures

import (
	"fmt"
)
func If() {
	
	emoji:= "🐶"

	if emoji == "🐶"{
		fmt.Println("Es un perrito!")
	} else if emoji == "🍎" {
		fmt.Println("Es una manzana!")
	}else{
		fmt.Printf("el emoji es: %s\n", emoji)
	}

	// En Go podemos declarar una variable corta dentro del if. PERO PODREMOS USARLA SÓLO DENTRO DEl if

	if fruta := "🍎"; fruta == "🍎" {
		fmt.Println("Es una manzana!")
	} else if fruta == "🍌" {
		fmt.Println("Es una banana!")

	}else {
		fmt.Printf("la fruta es: %s\n", fruta)
	}

}

