package paqueteerrors

import (
	"errors"
	"fmt"
)

func New(text string) error

func NewErrors() {
	err := errors.New("new error")
	fmt.Println(err)
}

//toma como argumento una cadena de caracteres (text) y retorna dicha cadena como una variable de tipo error