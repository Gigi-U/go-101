package varandconst

import (
	"fmt"
)
func Const(){
	// diferentes formas de declarar constantes 
	
	const producto string = "hueso"

	const paseo = "collar"
	
	juegoGato1, juegoGato2 := "puntero","ratón"

	const(
		cantidad = 1
		precio = 200.00
		stock = 3
	)

		fmt.Println(producto,cantidad, precio,stock, paseo, juegoGato1, juegoGato2)	
}

//! en CONSTANTES  no se puede usar := xq el compilador automaticamente lo leerá como variable