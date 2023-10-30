package datatypes
import (
	"fmt"
)
func AnimalesFrutas(){

	animales:=  make(map[string]string)
	animales["gato"]= "🐱"
	animales["perro"]= "🐶"
	animales["pato"]= "🦆"
	animales["canario"]= "🐤"

	fmt.Println(animales)
	// OTRA MANERA:

	fruits:= map[string]string{
		"manzana": "🍎",
		"banana": "🍌",
		"limon": "🍋",
		"uvas": "🍇",
	}
	fmt.Println(fruits)

	//DELETE:
	delete(fruits,"limon")

	//GET elementos
	//! si quiero retornar un elemento que no existe G0 me retorna el valor 0 del tipo de dato del valor
	fmt.Println(animales["gato"])
	// ok es un booleano
	emoji, ok := animales["gorila"]
	fmt.Println(emoji, ok)
	// debo poner el identificador blank en vez de la variable emoji xq no la voy a llamar y eso me va a tirar error
	if _, ok := animales["gorila"]; !ok {
		animales["gorila"] = "🦍"
	}
	fmt.Println(animales)
}

func LongitudMapa(){
	myMap := map[string]int{}
	fmt.Println(len(myMap))
}

func ActualizarValores(){

	students := map[string]int{"Benjamin": 20, "Nahuel": 26}	
	fmt.Println(students)
	students["Benjamin"] = 22
	fmt.Println(students)
}