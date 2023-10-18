package pgc2

import "fmt"

/* Ejercicio 2 - Calcular promedio
Un colegio necesita calcular el promedio (por estudiante) de sus calificaciones. 
Se solicita generar una función en la cual se le pueda pasar N cantidad de enteros y devuelva el promedio. 
& No se pueden introducir notas negativas.
*/
 func CalcularPromedio(){

	promedioFinal:= calculador(100,70,88,95)
	fmt.Println("El promedio obtenido es de :", promedioFinal)
 }
 
func calculador(calificaciones ...uint8)float32 {
	var promedio, suma float32
	suma = 0.00	
	contador:= 0
	for i:= range calificaciones{
		contador++
		suma += float32(calificaciones[i])
		promedio =  suma/ float32(contador)
	}
	return promedio
}