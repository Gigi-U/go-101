package mt6

import (
	"fmt"
	"log"
	"os"
	"strings"
)

const (
	filename = "digitalhouse/module1/week3/c6/mt6/categorias.csv"
)

// o main
func Mayorista() {
	defer func() {
		if err := recover(); err != nil {
			log.SetPrefix("main: ")
			log.Println(err)
			os.Exit(1)
		}
	}()
	readfile(filename)

}

func readfile(name string) {

	//? 1) traer el archivo csv

	// el readfile me pide el file y error
	file, err := os.ReadFile(name)
	// si no se puede acceder al file, se frena todo
	if err != nil {
		panic(err)
	}
	//? 2) limpiar los ;

	data := strings.Split(string(file), ";")
	fmt.Println(data)
	//? 3) recorrer el array de objetos y eliminar 



	//? 4) limpiar de las líneas las , y los corchetes

	//? 5) guardar en variables precio y cantidad de cada producto

	//? 6) sumar x categoria los valores resultantes

	//? 7)imprimir cada categoria con su valor promedio


	
}