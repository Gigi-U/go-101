package functions

//import "fmt"
//si quiero que se corra desde acá debo darle nombre con mayuscula a la funcion main (otro nombre) y dejar vacio los parametros de main.go
func main() {
	//sum(1,4,6,8)
	//fmt.Println(sum(1, 4))

}

// func sum(num1, num2, num3, num4 int)int{
// 	return num1+num2+num3+num4
// }

// ? si estoy recibiendo n cantidad de valores que pueden diferir de los parametros tenemos las FUNCIONES VARIATICAS. que nos permiten recibir un número variable de parámetros de un mismo tipPrintln(a)
func sum(nums ...int) int {
	total := 0
	for _, v := range nums {
		total += v
	}
	return total
}