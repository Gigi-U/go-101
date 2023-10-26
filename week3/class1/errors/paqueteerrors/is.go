package paqueteerrors

import (
	"errors"
	"fmt"
	"io/fs"
	"os"
)

func Is(err, target error) bool

func IsError (){
	_, err := os.Open("not_exist.txt")
	var notExist error = fs.ErrNotExist
	if errors.Is(err, notExist) {
	fmt.Println("The file does not exist")
	}

}

// La función Is() recibe como argumentos dos variables de tipo error. Estas son: err y target. Si encuentra coincidencia dentro de la cadena de errores de err con el error target, devuelve true.
//Esta función la utilizamos cuando necesitamos saber si —dentro de la cadena de errores— ocurrió uno exactamente igual a target. Con esta función comparamos error con error directamente