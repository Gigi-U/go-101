package channels

import (
	"fmt"
)

func main() {
	n := 3 

	ch := make(chan int) // definiendo el canal

	go multiplicarPorDos(n, ch)

	fmt.Println(<-ch) // la lectura de un canal es bloqueante. por lo que va a esperar a que se escriba algo en el canal para continuar con la ejecución
}

func multiplicarPorDos(num int, ch chan int) {
	res := num * 2
	ch <- res
}