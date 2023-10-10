package pg
import (
	"fmt"
)
/* Ejercicio 1 - Descuento
Una tienda de ropa quiere ofrecer a sus clientes un descuento sobre sus productos. Para ello necesitan una aplicación que les permita calcular el descuento basándose en dos variables: su precio y el descuento en porcentaje. La tienda espera obtener como resultado el valor con el descuento aplicado y luego imprimirlo en consola.
Crear la aplicación de acuerdo a los requerimientos. */

func Exercise1() {

	var precio, porcentaje, precioFinal float32

	precio, porcentaje = 30000, 20
	precioFinal = precio - ((precio * porcentaje) / 100.00)
	fmt.Println("El valor final del Producto es: ", precioFinal)

	// otra forma es cambiando el valor de precio SI ES UNA VARIABLE:
	fmt.Println("---------otra forma---------")

	var precio1, porcentaje1 float32
	precio1, porcentaje1 = 30000, .20
	precio1 = precio1 -(precio1 * porcentaje1)
	fmt.Println("El valor final del Producto es: ", precio1)
}