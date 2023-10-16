package pgc1
import (
	"fmt"
)

/* Ejercicio 2 - Préstamo
Un banco quiere otorgar préstamos a sus clientes, pero no todos pueden acceder a los mismos. El banco tiene ciertas reglas para saber a qué cliente se le puede otorgar: solo le otorga préstamos a clientes cuya edad sea mayor a 22 años, se encuentren empleados y tengan más de un año de antigüedad en su trabajo. Dentro de los préstamos que otorga, no les cobrará interés a los que su sueldo sea mejor a $100.000.
Es necesario realizar una aplicación que tenga estas variables y que imprima un mensaje de acuerdo a cada caso.
Pista: Tu código tiene que poder imprimir al menos tres mensajes diferentes. */


func Prestamo() {
	var edad, antiguedad uint8
	var empleado bool
	var sueldo float32

	edad = 21
	antiguedad = 12
	empleado = true
	sueldo= 100000
	
	if edad >= 22 && empleado == true && antiguedad >= 12 {
		if sueldo > 100000{
			fmt.Println("\n Edad: ", edad, "\n Empleado: ", empleado, "\n antiguedad: ", antiguedad, "\n sueldo: ",sueldo, "\n Puede solicitar un préstamo, y no se le cobrará interes")

		}else {
			fmt.Println("\n Edad: ", edad, "\n Empleado: ", empleado, "\n antiguedad: ", antiguedad, "\n sueldo: ",sueldo, "\n puede solicitar un préstamo, pero se le cobrará interés")
		}
	}else {
		fmt.Println("\n Edad: ", edad, "\n Empleado: ", empleado, "\n antiguedad: ", antiguedad, "\n sueldo: ",sueldo, "\n No cumple con los criterios para solicitar u préstamo")

	}

}

