package goroutines

import (
	"fmt"
	"time"
)

func main() {

	go contador("Oveja")
	go contador("Pez")

	// time.sleep() le da tiempo a las goroutines de ejecutarse. si no lo agrego no se llega a imprimir nada
	time.Sleep(time.Second * 2) // 2 segundos
	//time.Sleep(time.Second) 1 segundo
	//time.Sleep(time.Second / 2 ) 1/2 segundo


}

func contador(elem string) {
	for i := 1; true; i++ {
		fmt.Println(i, elem)
		time.Sleep(time.Second / 2)
	}
}