package functions

import "fmt"

 //Por último, probamos nuestra aplicación pasándole la operación que queramos realizar.
func Calcular() {
	fmt.Println(operacionAritmetica(Suma, 2, 3, 2, 1, 2, 3, 4, 5, 6))
}
 //Declararemos las constantes con las operaciones a realizar:

const (
	Suma   = "+"
	Resta  = "-"
	Multip = "*"
	Divis  = "/"
 )

 //Luego, creamos las funciones que realizarán las operaciones.

 func opSuma(valor1, valor2 float64) float64 {
	return valor1 + valor2
 }
 
 func opResta(valor1, valor2 float64) float64 {
	return valor1 - valor2
 }
 
 func opMultip(valor1, valor2 float64) float64 {
	return valor1 * valor2
 }
 
 func opDivis(valor1, valor2 float64) float64 {
 
	if valor2 == 0 {
		return 0
	}
	return valor1 / valor2
 }
 /*También crearemos la función que se encargará de recibir la operación a realizar y los valores a los cuales se le aplicará la operación. 
 Por cada operación, llamaremos a una función que reciba los valores y la función que vamos a ejecutar por ese operador.*/

 func operacionAritmetica(operador string, valores ...float64) float64 {
	switch operador {
	case Suma:
		return orquestadorOperaciones(valores, opSuma)
	case Resta:
		return orquestadorOperaciones(valores, opResta)
	case Multip:
		return orquestadorOperaciones(valores, opMultip)
	case Divis:
		return orquestadorOperaciones(valores, opDivis)
	}
 	return 0
 }
 
 //Crearemos esa función que se encargará de orquestar las operaciones:
 func orquestadorOperaciones(valores []float64, operacion func(value1, value2 float64) float64) float64 {
	var resultado float64
	for i, valor := range valores {
		if i == 0 {
			resultado = valor
		} else {
			resultado = operacion(resultado, valor)
		}
	} 
	return resultado
 }
 

 