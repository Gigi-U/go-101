package mtc2
/* Ejercicio 1 - Calcular salario
Una empresa marinera necesita calcular el salario de sus empleados basándose en la cantidad de horas trabajadas por mes y la categoría.

1. Categoría C: su salario es de $1.000 por hora.
2. Categoría B: su salario es de $1.500 por hora, más un 20 % de su salario mensual.
3. Categoría A: su salario es de $3.000 por hora, más un 50 % de su salario mensual.

Se solicita generar una función que reciba por parámetro la cantidad de minutos trabajados por mes y la categoría, y que devuelva su salario. */

// Here we set the constants to use them as conditions at the Switch Function
import (
	"errors"
	"fmt"
)

const(
	catA = "Category A"
	catB = "Category B"
	catC = "Category C"
	//  ERROR EXAMPLE
	wrongCategory ="Category X"
)
// Here we set our orchestrating function
func Exercise1() {
	
	salaryAFunction, err := orchestrator(catA)
	salaryBFunction, err := orchestrator(catB)
	salaryCFunction, err := orchestrator(catC)
	//  ERROR EXAMPLE
	_, err = orchestrator(wrongCategory)
	
	// Here we set the error messaje	  
	if err!= nil{
		fmt.Println("An Error has ocurred:", err.Error())
	}		

	// here we set the final ammount variables with the parameter minutes. 
	salaryA := salaryAFunction(7560)
	salaryB := salaryBFunction(7560)
	salaryC := salaryCFunction(7560)

	// Here we print the final ammounts 
	fmt.Printf("Category A salary: %.2f \n", salaryA)
	fmt.Printf("Category B salary: %.2f \n", salaryB)
	fmt.Printf("Category C salary: %.2f \n", salaryC)
	
	//! %.2f me muestra .00 en decimales		
}
// Here we create the Switch Function that will recieve a string and return the total amount or the error string.
	func orchestrator(category string)(func(int)float32, error ) {
	switch category {
	case catA:
		return categoryA, nil
	case catB:
		return categoryB, nil
	case catC:
		return categoryC, nil
	default:
		return nil, errors.New("Category Not Found")
	}
}
// Common variables for Salary calculator Functions
var totalSalary float32
// Here we calculate the salary -  Category A
func categoryA(minutes int)float32{
	if minutes > 0{
		totalSalary = (float32(minutes)/60) * 3000 + 3000* .5 
	}else{fmt.Println("Please enter a valid number of minutes.")}	

	return  totalSalary
}
// Here we calculate the salary - Category B
func categoryB(minutes int)float32{
	if minutes > 0 {
		totalSalary = (float32(minutes)/60) * 1500 + 3000  * .2 
	}else{fmt.Println("Please enter a valid number of minutes.")}	

	return  totalSalary
}
// Here we calculate the salary - Category C
func categoryC(minutes int)float32 {
	if minutes > 0 {
		totalSalary = (float32(minutes)/60) * 1000
	}else{fmt.Println("Please enter a valid number of minutes.")}	
	return totalSalary
}
