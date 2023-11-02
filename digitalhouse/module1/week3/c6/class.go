package c6

import (
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

const (
	filename = "digitalhouse/module1/week3/c6/data.csv"
)
// o main()
func MesaC6() {
	defer func() {
		if err := recover(); err != nil {
			log.Fatal(err)
		}
	}()
	generarEstimacion(filename)
}

func generarEstimacion(name string) {
	file, err := os.ReadFile(name)
	if err != nil {
		panic(err)
	}

	data := strings.Split(string(file), "\n")

	// Mapa donde se totalizará cada categoria
	estimaciones := map[string]float64{}

	for i := 1; i < len(data)-1; i++ {
		line := strings.Split(string(data[i]), ",")

		// Si no existe la categoria en el mapa la creo
		_, errb := estimaciones[line[1]]
		if !errb {
			estimaciones[line[1]] = 0.0
		}

		precioActual, err  := strconv.ParseFloat(line[4], 64)
		if err != nil {
			log.Panicln("No se pudo convertir el precioActual")
		}
		cantidadActual, err  := strconv.ParseFloat(strings.TrimSpace(line[5]), 64)
		if err != nil {
			log.Panicln("No se pudo convertir el cantidadActual")
		}

		// Calculo el total del producto y lo sumo a su categoria
		estimaciones[line[1]] += precioActual * cantidadActual
	}

	// Creo al archivo donde se almacenarán las estimaciones y lo abro para poder escribirlo
	fileOpen, err := os.OpenFile("digitalhouse/module1/week3/c6/estimaciones.csv", os.O_WRONLY|os.O_CREATE, 0644)
	if err != nil {
		panic(err)
  }
  defer fileOpen.Close()

	// Agrego el titulo al archivo
	_, err = fileOpen.WriteString("Categoría\t\tEstimativoPorCategoría\n")
	if err != nil {
		panic("Error al escribir el archivo")
	}

	// Agrego cada categoria y su total al archivo
	for cat, total := range estimaciones {
		_, err = fileOpen.WriteString(fmt.Sprintf("%s\t\t%.2f\n", cat, total))
		if err != nil {
			panic("Error al escribir el archivo")
		}
	}

	//  Muestro por consola cada categoria y su total
	fmt.Print("Categoría\t\tEstimativoPorCategoría\n")
	for cat, total := range estimaciones {
		fmt.Printf("%s\t\t%.2f\n", cat, total)
	}
}