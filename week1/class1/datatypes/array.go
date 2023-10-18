package datatypes
import (
	"fmt"
)

func Hello(){
	var a [2]string
	a[0] = "Hello"
	a[1] = "World"
	fmt.Println(a)
}

//func ObtenerEstacion(meses []string) si el array estuviera en otra función completar los parámetros
func Estaciones(){

	meses:= []string{"Enero", "Febrero", "Junio", "Agosto", "Abril"}

	for _, mes := range meses {
		if mes == "Enero" || mes == "Febrero" || mes == "Marzo" {
			fmt.Println("Verano")
		} else if mes == "Abril" || mes == "Mayo" || mes == "Junio" {
			fmt.Println("Otoño")
		} else if mes == "Julio" || mes == "Agosto" || mes == "Septiembre" {
			fmt.Println("Invierno")
		} else if mes == "Octubre" || mes == "Noviembre" || mes == "Diciembre" {
			fmt.Println("Primavera")
		} else {
			fmt.Println("No existe este mes")
		}
	}
}
