package functions

import "fmt"


func Hello(firstName string, lastName string){
	fmt.Printf("hello %s %s", firstName,lastName)
}
func Change(){
	emoji:="🐶"
	ChangeValue(&emoji)
	fmt.Printf(emoji)
}
// Cuando hacemos funciones que pasan parámetros, lo que en realidad estamos recibiendo es una copia del valor de esa variable. Si nosotros queremos modificar el contenido de esa variable debemos utilizar una Función por Referencia. se hace un puntero o pointer de la variable inicial

func ChangeValue(value *string){
	*value = "🐕\n"
}
