package pgc2

import "fmt"

func Tutorial() {
	resultado := suma(4, 8)
	fmt.Println("El resultado es:", resultado)
}

func suma(a, b int) int {

	return a + b
}