package pgc2

import "fmt"
/* Ejercicio 1 - Impuestos de salario
Una empresa de chocolates necesita calcular el impuesto de sus empleados al momento de depositar el sueldo.

Para cumplir el objetivo es necesario crear una función que devuelva el impuesto de un salario, teniendo en cuenta que si la persona gana más de $50.000 se le descontará un 17 % del sueldo y si gana más de $150.000 se le descontará, además, un 10 % (27 % en total).
 */
var salario, impuestoEnMoneda float32;

func ImpuestosSalario(){

	impuesto:=calculadorImpuesto(500000)
	fmt.Printf("El impuesto en moneda es de:$%.2f \n", impuesto)

}
func calculadorImpuesto(salario float32)float32{
	if salario >= 50000 && salario < 150000{
		impuestoEnMoneda = 17 * salario / 100

	}else if salario >= 150000{
		impuestoEnMoneda = 27 * salario / 100
	}else {
		fmt.Println("No debe pagar impuestos")
	}
	return impuestoEnMoneda
}