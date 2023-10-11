package functions

func Anonima(){
	// forma 1 con variable declarada
	x()

	// forma 2 sin declarar variable
	func (){
		println("Es una función anónima autoejecutada")
	}()
	// forma 2 con parámetros y sin declarar variable
	func (text string){
		println("Es una función anónima autoejecutada", text)
	}("y lleva parámetros")
}
var x = func ()  {
	println("Es una función anónima")
}
