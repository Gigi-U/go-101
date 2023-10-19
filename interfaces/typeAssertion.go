package interfaces

import "fmt"

//var i any --- es un alias a poner interface

var i interface{} = "hello"

func TypeAssertion() {
	s := i.(string)
	fmt.Println(s)

	s, ok := i.(string)
	fmt.Println(s, ok)

	f, ok := i.(float64)
	fmt.Println(f, ok)

	//f = i.(float64) // panic
	//fmt.Println(f)
}

//La aserción de tipos provee acceso al tipo de datos exacto que está abstraído por una interfaz.
