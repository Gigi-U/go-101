package mt6

import (
	"fmt"
	"log"
	"os"
	//"strconv"
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
	file, err := os.ReadFile(name)
	if err != nil {
		panic(err)
	}
	data := strings.Split(string(file), ";")
	//fmt.Println(data)
	//var total float64
	for i := 0; i < len(data)-1; i++ {
		line := strings.Split(string(data[i]), ",")

		fmt.Println("\n",line)

/* 		 if i != 0 {
			listaProductos, err := strconv.ParseFloat(line[2], 64)

			for i := 0; i < len(listaProductos)-1; i++ {
				

			line2 := strings.Split(string(listaProductos[i]), "],[")


			precioActual, err := strconv.ParseFloat(line2[2], 64)

			if err != nil {
				log.Panicln("no se pudo convertir el precio actual")
			}
			cantidadActual, err := strconv.ParseFloat(line[3], 64)
			if err != nil {
				log.Panicln("no se pudo convertir la cantidad actual")
			} 		


			totalProducto := precio * cantidad

			total += totalProducto
		}
		for i:=0;  i < len(line); i++ {
			fmt.Printf("%s \t\t", line[i])
			if i == len(line)-1 {
				fmt.Print("\n")
			}
		}
	}
	fmt.Printf("\t\t%s \t\t%.2f\n", "total", total)  */
}}