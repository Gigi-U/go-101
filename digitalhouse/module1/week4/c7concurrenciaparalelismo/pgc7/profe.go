package pgc7

import (
	"fmt"
	"os"
	"time"
)

func main() {
	var input string

	channelDevolucion := make(chan string)
	channelPrestamo := make(chan string)

	//Se puede hacer así
	go Devolucion(channelDevolucion)

	go PrintDevolucion(channelDevolucion) 	

	//O Así
	go func(channelPrestamo chan string) {
		defer close(channelPrestamo)
		for {
			time.Sleep(time.Second/2)
			channelPrestamo <- "Prestamo procesado"
		}
	}(channelPrestamo)


	go func(channelPrestamo chan string) {
		for prestamo := range channelPrestamo{
			fmt.Println(prestamo)
		}
	}(channelPrestamo)
	
	// condición de corte
	fmt.Scan(&input)
	if input !=""{
		fmt.Println("saliendo del programa")
		os.Exit(0)
	}
}

func Devolucion(channelDevolucion chan string){
	defer close(channelDevolucion)
	for {
		time.Sleep(time.Second)
		channelDevolucion <- "Devolucion procesada"
	}
}
func Prestamo(channelPrestamo chan string){
	
}

func PrintDevolucion(channelDevolucion chan string){
	for devolucion := range channelDevolucion{
		fmt.Println(devolucion)
	}
}

func PrintPrestamo(channelDevolucion chan string){
	
}