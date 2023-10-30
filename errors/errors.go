package errors
import (
	"fmt"
	"errors" //Este paquete contiene cuatro funciones: New(), Unwrap(), As() y Is()
	)

	func MyErrors() {
		statusCode := 404;
		if statusCode > 399 {
		fmt.Println(errors.New("La petición ha fallado."))
		return
		}
		fmt.Println("El programa finalizó correctamente")
}


/* 
3 formas -  nos enfocaremos en errors.New("string")  
Error()
errors.New()
fmt.Errorf()
 
en Go los errores son representados por una interfaz que contiene un único Método Error que retorna un string

esta sería la interface:
type error interface {
	Error() string
}

https://cheatography.com/kstep/cheat-sheets/http-status-codes/


https://pkg.go.dev/errors 


*/