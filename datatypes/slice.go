package datatypes

import (
	"fmt"
)

// LOS SLICE SON APUNTADORES A ARRAYS, y al ser apuntadores no almacenan ningún dato
func Slices(){
		//BOOLEANOS
		var boolean = []bool{true, false}
		MySlice(boolean)
	
		// STRINGS
		set := [7]string{"🐷", "🐸", "🐝", "🐛", "🦋", "⚽️", "🏀"} 
		MyAnimals(set)
	
		// TAMAÑO Y CAPACIDAD | creando apuntador
		morfi:=[5]string{"🌭", "🍔", "🍟", "🍕", "🌯"}
		MiMorfi(morfi)
	
		// MAKE
		a := make([]int, 5)
		WithMake("a", a)
		b := make([]int, 0, 5)
		WithMake("b", b)
		c := b[:2]
		WithMake("c", c)
		d := c[2:5]
		WithMake("d", d)
	
		// RANGOS
		primes := []int{2, 3, 5, 7, 11, 13}
		WithRanges(primes)
	
		// VALOR CERO - ojo, si le agrego {} estoy creando un slice literal entonces ya no es de valor nulo xq está inicializado
		var plantas []string
		ValorCero(plantas)
	
	}
	
	func MySlice(boolean []bool){
		fmt.Println(boolean[0])
	}
	
	func MyAnimals(set[7]string){
		animals := set[0:5]
		//animals := set[:5] si no le agrego el 0:5 Go entiende de todos modos que va desde la posición 0. Lo mismo pasa con el final.ej: [3:]. Si no especifico nada se imprime todo ej: [:]
		sports := set[5:7]
		sports [0] = "🏈"  // va a cambiar el array deportes, pero tambien el array set
		fmt.Println("Array",set) 
		fmt.Println("animales",animals) 
		fmt.Println("deportes",sports) 
		fmt.Println("deportes posición 0",sports[0]) 
	}
	
	func MiMorfi(morfi[5]string){
		mcDonalds := morfi[1:3]
		fmt.Println("Todo el morfi",morfi) 
		fmt.Println("Morfi McDonalds",mcDonalds) 
		fmt.Println("tamaño",len(mcDonalds)) // me muestra la cantidad de elementos
		fmt.Println("capacidad",cap(mcDonalds)) // como mcDonalds es un apuntador de morfi, me muestra la cantidad de elementos desde donde nace mcDonalds en adelante
			/* así se ve el resultado:
			Todo el morfi [🌭 🍔 🍟 🍕 🌯]
			Morfi McDonalds [🍔 🍟]
			tamaño 2
			capacidad 4 */
		//* AGREGANDO elemento a SLICE (recordando que es dinámico y a dif. de los array sí puedo hacerlo.)	
		mcDonalds = append(mcDonalds,"🍦","🥤","🍩")
		fmt.Println("Morfi McDonalds con extras",mcDonalds) 
		fmt.Println("tamaño",len(mcDonalds))
		fmt.Println("capacidad",cap(mcDonalds))
		fmt.Println("Todo el morfi",morfi) 
		/* 	así se ve el resultado:
	
			Morfi McDonalds con extras [🍔 🍟 🍦 🥤 🍩]
			tamaño 5
			capacidad 8
			Todo el morfi [🌭 🍔 🍟 🍕 🌯] 
			
			Cambia el tamaño y capacidad del slice pero no el array original, ya que cambia la referencia al array fruits al superar el tamaño del array original*/
	
		//* Creando y array de slice - SLICE LITERALS (no se basa en ningún array original)
		hands:=[]string{"👋🏼", "🤚🏼", "🖐🏼"}	
		hands = append(hands,"🖖🏼", "👌🏼")	
	
		fmt.Println("juguetes",hands) 
		fmt.Println("tamaño",len(hands))
		fmt.Println("capacidad",cap(hands))
	}
	// me permite construir arrays con una función preconstruida, en donde yo puedo incluir un tamaño y una capacidad
	func WithMake(s string, x []int){
		fmt.Printf("%s len=%d cap=%d %v\n",
			s, len(x), cap(x), x)
	}
			/* withmake() me mostrará esto:
	
			a len=5 cap=5 [0 0 0 0 0]
			b len=0 cap=5 []
			c len=2 cap=5 [0 0]
			d len=3 cap=3 [0 0 0] */
	
	func WithRanges(primes []int){
		fmt.Println(primes[1:4]) // Si no ponemos un valor después de los “:” toma hasta el fin de elementos del slice y viceversa.
	}
	
	
	func ValorCero(plantas[]string){
		fmt.Println(plantas) // se crea vacio, con capacidad y tamaños 0
		fmt.Println(plantas == nil) // nil es igual a nulo - RETORNA true
	
	}
	
	

