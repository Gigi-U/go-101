package professor

import "fmt"

// 	La Real Academia Española quiere saber cuántas letras tiene una palabra y luego tener cada una de las letras por separado para deletrearla.
//  A - Crear una aplicación que tenga una variable con la palabra e imprimir la cantidad de letras que tiene la misma.
//  B - Luego, imprimir cada una de las letras.

var hello string ="hello"
// function that captures words letters ammount and prints each letter
func letters() {
	sum:=0
	for _,v := range hello{
		sum++
		fmt.Println(string(v))
 	}
	 fmt.Print("number of letters in word:", sum)

}
// equals main()
func Professor1() {
	letters()
}
