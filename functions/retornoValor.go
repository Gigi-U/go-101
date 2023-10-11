package functions

import "fmt"

//&  FUNCION QUE RETORNA VALOR
 
const (
	Sumar   = "+"
	Restar  = "-"
	Multipl = "*"
	Divid  = "/"
 )
func opAritmetica(valor1, valor2 float64, operador string) float64 {
	switch operador {
	case Sumar:
		return valor1 + valor2
	case Restar:
		return valor1 - valor2
	case Multipl:
		return valor1 * valor2
	case Divid:
		if valor2 != 0 {
			return valor1 / valor2
		}
	}
	return 0
 }
 func RetornoValor() {
	fmt.Println(operacionAritmetica(Sumar,6,2))
	fmt.Println(operacionAritmetica(Restar,6, 2))
	fmt.Println(operacionAritmetica(Multipl,6, 2))
	fmt.Println(operacionAritmetica(Divid,6, 2))
 }