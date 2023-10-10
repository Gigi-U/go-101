package structures

import (
	"fmt"
)
func For() {
	fmt.Println("----------FRUTAS-----------")

	frutas :=	[]string{"manzana",	"banana", "pera"}

	for i, fruta := range
	frutas {
	fmt.Println(i, fruta)
	}

	fmt.Println("--------ES IMPAR- Saltar a la siguiente iteración-------------")	
	for i := 0; i < 10; i++{
	if i % 2 == 0 {
	continue
	}
	fmt.Println(i, "es impar")
	}


	fmt.Println("--------ROMPER BUCLE-------------")	

	sum := 0
	for {
	sum++
	if sum >= 1000 {
	break
	}
	}
	fmt.Println(sum)

/* 	fmt.Println("--------BUCLES INFINITOS-------------")	

	sum := 0
	for true {
	sum++
	}

	sum := 0
	for true {
	sum++
	} */
}

