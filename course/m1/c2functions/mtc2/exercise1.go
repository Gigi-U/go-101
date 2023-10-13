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
	wrongCategory ="Category D"
)
// Here we set the function that will be called by the main.go file
func Exercise1() {
	// here we create two variables that will be associated with the orchestrator. The first one connects it with the salary calculator Funcions, the second one with the conditional statement of the error message. On the other hand we connect the orchestrator with the category constant imput required.
	funcAMinutesParam, err := orchestrator(catA)
	funcBMinutesParam, err := orchestrator(catB)
	funcCMinutesParam, err := orchestrator(catC)
	//  ERROR EXAMPLE
	_, err = orchestrator(wrongCategory)
	
	// Here we set the error message	  
	if err!= nil{
		fmt.Println("An Error has ocurred:", err.Error())
	}		

	// here we create the variable that will be associated with te salary Function return value. On the other hand we set the minutes imput required for each function. 
	salaryA := funcAMinutesParam(7560)
	salaryB := funcBMinutesParam(7560)
	salaryC := funcCMinutesParam(0)

	// Here we print the final ammounts that the Functions return us.
	fmt.Printf("Category A salary: %.2f \n", salaryA)
	fmt.Printf("Category B salary: %.2f \n", salaryB)
	fmt.Printf("Category C salary: %.2f \n", salaryC)
				//! %.2f shows result as float + 2 decimals 		
}
// Here we create an orchestrator FUNCTION that contains a Switch statement. With it, the application will know which function must be called. It will also notify as error if the category given do not exist.
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

// common message variable for functions below
var timeToWork string="No Bee no Honey. Lets get things done! 🐝"

// Here we calculate the salary for  Category A
func categoryA(minutes int)float32{
	var totalSalary float32
	if minutes > 0{
		totalSalary = (float32(minutes)/60) * 3000 + 3000* .5 
	}else{fmt.Println("A category employee: ",timeToWork)}	

	return  totalSalary
}

// Here we calculate the salary for Category B
func categoryB(minutes int)float32{
	var totalSalary float32
	if minutes > 0 {
		totalSalary = (float32(minutes)/60) * 1500 + 3000  * .2 
	}else{fmt.Println("B category employee: ",timeToWork)}	

	return  totalSalary
}

// Here we calculate the salary for Category C
func categoryC(minutes int)float32 {
	var totalSalary float32
	if minutes > 0 {
		totalSalary = (float32(minutes)/60) * 1000
	}else{fmt.Println("C category employee: ",timeToWork)}	
	return totalSalary
}
