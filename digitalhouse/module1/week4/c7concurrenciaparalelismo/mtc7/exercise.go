package mtc7

import (
	"fmt"
	"sync"
)

/* Consigna
Se requiere un programa para calcular si un número es par o impar desde una lista de int.
Para esto se deben crear dos funciones:

- par(chan int) debe recibir a través de un channel de int los valores pares enviados y mostrarlos por la terminal.

- impar(chan int) debe recibir a través de un channel los valores impares y mostrarlos por terminal.

El programa debe comenzar dos goroutines que se ejecuten concurrentemente.
En la goroutine principal se itera el slice provisto con valores y, en caso de que el valor sea par, se lo debe enviar al channel correspondiente. De igual manera para los impares.

Ayuda: Si el resto de una división por dos da cero es un número par (i%2 == 0) y si da 1 es impar (i%2 == 0 ).
Para obtener el resto de una operación, utilizar el operador %: i%2 == 1.

*/

var (
	valores = []int{1, 2, 3, 4, 5, 6, 7, 8, 9}
)

// Función para determinar si un número es par
func even(num int) bool {
	return num%2 == 0
}


// va a imprimir el slice de numeros pares
func evenNum(channelEven chan int, wg *sync.WaitGroup) {
	defer wg.Done() // Marca la finalización de la goroutine
	for evenNumber := range channelEven {
		fmt.Printf("Even: %d\n", evenNumber)
	}
}

// va a imprimir el slice de numeros impares
func oddNum(channelOdd chan int, wg *sync.WaitGroup) {
	defer wg.Done() // Marca la finalización de la goroutine
	for oddNumber := range channelOdd {
		fmt.Printf("Odd: %d\n", oddNumber)
	}
}


func EvenOrOdd() {
	var wg sync.WaitGroup
    wg.Add(2) // Solo esperamos la finalización de dos goroutines

    channelEven := make(chan int)
    channelOdd := make(chan int)

    go evenNum(channelEven, &wg)
    go oddNum(channelOdd, &wg)
	//se itera el slice provisto con valores y, en caso de que el valor sea par, se lo debe enviar al channel correspondiente
    for _, v := range valores {
        if even(v) {
            channelEven <- v
        } else {
            channelOdd <- v
        }
    }

    close(channelEven) // Cierra el canal de números pares cuando termina
    close(channelOdd)   // Cierra el canal de números impares cuando termina

    wg.Wait()
}

/* TERMINAL:

Odd: 1
Odd: 3
Even: 2
Even: 4
Even: 6
Odd: 5
Odd: 7
Odd: 9
Even: 8 
*/