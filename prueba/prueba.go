package prueba
import (
	"errors"
	"fmt"
)

const (
	catA = "categoriaA"
	catB = "categoriaB"
	catC = "categoriaC"
)

func Prueba() {
	// sueldoA := calcularSalario(catA, 10)
	// sueldoB := calcularSalario(catB, 10)
	// sueldoC := calcularSalario(catC, 10)

	// fmt.Printf("Sueldo cat A: %.2f \n", sueldoA)
	// fmt.Printf("Sueldo cat A: %.2f \n", sueldoB)
	// fmt.Printf("Sueldo cat A: %.2f \n", sueldoC)

	calcCategoriaA, err := orquestador(catA)
	calcCategoriaB, err := orquestador(catB)
	calcCategoriaC, err := orquestador(catC)

	_, err = orquestador("asdf")

	if(err != nil) {
		fmt.Println("Error:", err.Error())
	}

	sueldoCatA := calcCategoriaA(10)
	sueldoCatB := calcCategoriaB(10)
	sueldoCatC := calcCategoriaC(10)

	fmt.Printf("sueldoA: %.2f \n", sueldoCatA)
	fmt.Printf("sueldoB: %.2f \n", sueldoCatB)
	fmt.Printf("sueldoC: %.2f \n", sueldoCatC)
}

func orquestador(categoria string) (func(int) float32, error) {
	switch categoria {
	case catA:
		return categoriaA, nil
	case catB:
		return categoriaB, nil
	case catC:
		return categoriaC, nil
	default:
		return nil, errors.New("categoria no encontrada")
	}
} 

func categoriaA(horas int) float32 {
	return float32(horas) * 3000.0 * 1.5
}

func categoriaB(horas int) float32 {
	return float32(horas) * 1500.0 * 1.2
}

func categoriaC(horas int) float32 {
	return float32(horas) * 1000
}