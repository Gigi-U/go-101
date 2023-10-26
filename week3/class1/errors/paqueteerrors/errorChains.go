package paqueteerrors

import (
	"fmt"
)

func ErrorChains(){
	error_1 := fmt.Errorf("first error")
	error_2 := fmt.Errorf("second error: %w", error_1)
	fmt.Println(error_2)
}

/* En Go se puede envolver un error dentro de otro. Esto forma una jerarquía de errores, donde una instancia de error
envuelve a otra y esta —a su vez— puede estar envuelta en otra.
En definitiva, los errores en Go forman cadenas de errores. Para envolver un error en otro utilizamos la función
fmt.Errorf() junto con la directiva %w. */