package paqueteerrors
import (
	"errors"
	"fmt"
)
func Unwrap(err error) error

func UnWrap(){
	error_1 := fmt.Errorf("first error")
	error_2 := fmt.Errorf("second error: %w", error_1)
	err := errors.Unwrap(error_2)
	if err == error_1 {
		fmt.Println("same error")
	}
}

// Toma el último error de la cadena y comprueba si este contiene otro error. En tal caso, lo retorna; si no devuelve ni