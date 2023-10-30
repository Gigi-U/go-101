package paqueteerrors

import (
	"errors"
	"fmt"
	"io/fs"
	"os"
)

func As(err error, target interface{}) bool

func AsError() {
	_, err := os.Open("not_exist.txt")
	var pathError *fs.PathError
	if errors.As(err, &pathError) {
		fmt.Println("Path error", pathError.Path)
	}

}

// La función As() recibe dos argumentos: una variable del tipo err (error) y un target (interfaz vacía) que nos permite pasarle un puntero a un tipo error, y retorna una variable bool.

//Lo que hace la función es comparar —dentro de la cadena de errores de err— con un puntero a un tipo de errores. Cuando encuentra una primera coincidencia, asigna a target el error y devuelve true.