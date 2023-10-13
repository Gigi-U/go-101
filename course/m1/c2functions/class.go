package c2functions

import (
	"errors"
	"fmt"
)

const (
	minimo   = "minimo"
	promedio = "promedio"
	maximo   = "maximo"
)

func Orquestador() {
	minFunc, err := operation(minimo)
	maxFunc, err := operation(maximo)
	averageFunc, err := operation(promedio)

	if err != nil{
		fmt.Println("Se produjo un error")
	}
	minValue := minFunc(2, 3, 3, 4, 10, 2, 4, 5)
	maxValue := maxFunc(2, 3, 3, 4, 1, 2, 4, 5)
	averageValue := averageFunc(2, 3, 3, 4, 1, 2, 4, 5)

	fmt.Println("Minimo valor", minValue)
	fmt.Println("Max valor", maxValue)
	fmt.Println("Valor promedio", averageValue)


}

func operation(name string) (func(...int) int, error) {
	switch name {
	case minimo:
		return Minimo, nil
	case maximo:
		return Maximo, nil
	case promedio:
		return Promedio, nil
	default:
		return nil, errors.New("No se encontró la operación")
	}
	
}

// Minimo es una funcion que retorna el valor minimo dentro de un conjunto
// parametros: ...int
// retorna int
func Minimo(valores ...int) int {
	minimo := 1000
	for _, valor := range valores {
		if valor < minimo {
			minimo = valor
		}
	}
	return minimo
}

// Máximo es una funcion que retorna el valor máximo dentro de un conjunto
func Maximo(valores ...int) int {
	var maximo int
	for _, valor := range valores {
		if valor > maximo {
			maximo = valor
		}
	}
	return maximo
}

// Promedio es una funcion que retorna el valor promedio dentro de un conjunto
func Promedio(valores ...int) int {
	suma := 0
	for _, valor := range valores {
		suma += valor
	}
	promedio := suma / len(valores)
	return promedio
}