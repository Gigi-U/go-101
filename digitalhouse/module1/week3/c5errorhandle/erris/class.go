package erris

import (
	"errors"
	"fmt"

)

/*
En la función main, definir una variable salary y asignarle un valor de tipo “int”.
Crear un error con Errors.New() con el mensaje “Error: el salario ingresado es muy bajo” y lanzarlo en caso de que salary sea menor o igual a 10.000. La validacion debe hacerse con la funcion Is() dentro del main.
*/
const (
	salaryConst =10000
)
var (
	// ErrSalary es un error que utiliza el paquete errors
	ErrSalary = errors.New("Salario menor o igual a 10000")
)
func ErrorSalarioBajo(){
	// Validamos el salario
	result, err := validateSalary(5000)
	// Verificamos si el error es del tipo ErrSalary
	if errors.Is(err, ErrSalary) {
		fmt.Println(err.Error())
	} else {
		fmt.Println("El salario es:", result)
	}

}
// validateSalary valida el salario
func validateSalary(salary int)(int, error){

	if salary <= salaryConst {
		return 0, ErrSalary
	}

	return salary, nil
}

//& errorIs(): cuando necesitas chequear un error ESPECÍFICO