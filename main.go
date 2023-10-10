package main

import (
	// here you import all other packages
	"fmt"

	"github.com/Gigi-U/go-101.git/course/m1/c1basis"
	"github.com/Gigi-U/go-101.git/course/m1/c1basis/mtc1"
	"github.com/Gigi-U/go-101.git/course/m1/c1basis/pgc1"
	"github.com/Gigi-U/go-101.git/course/m1/c2functions/pgc2"
	"github.com/Gigi-U/go-101.git/datatypes"
	"github.com/Gigi-U/go-101.git/operators"
	"github.com/Gigi-U/go-101.git/pointers"
	"github.com/Gigi-U/go-101.git/structures"
	"github.com/Gigi-U/go-101.git/varandconst"
)

func main() {
	fmt.Println("---------DATATYPES-------------")
	//FOLDER: datatypes FILE: types
	datatypes.DataTypes()
	datatypes.ConPrintf()
	datatypes.DifTiposDatos()
	fmt.Println("--------ARRAY--------------")
	//FOLDER: datatypes FILE: array
	datatypes.Hello()
	datatypes.Estaciones()
	fmt.Println("--------MAPS --------------")
	//FOLDER: datatypes FILE: map
	datatypes.AnimalesFrutas()
	datatypes.LongitudMapa()
	datatypes.ActualizarValores()
	fmt.Println("----------SLICES------------")
	//FOLDER: datatypes FILE: slice
	datatypes.Slices()

	fmt.Println("--------STRUCTURES--------------")
	structures.Estructuras()
	fmt.Println("--------FOR --------------")
	structures.For()
	fmt.Println("--------IF --------------")
	structures.If()
	fmt.Println("--------SWITCH --------------")
	structures.Switch()
	fmt.Println("--------WHILE --------------")
	structures.While()

	fmt.Println("--------VARIABLES --------------")
	varandconst.Variables()
	fmt.Println("-------- CONSTANTES--------------")
	varandconst.Const()

	fmt.Println("--------POINTERS--------------")
	pointers.Pointers()

	fmt.Println("\n--------OPERATORS--------------")
	operators.Operators()

	fmt.Println("--------FUNCTIONS--------------")

	fmt.Println("--------CLASES--------------")
	fmt.Println("--------CLASE 01--------------")
	c1basis.Class()
	mtc1.Exercise1()
	mtc1.Exercise2()
	pgc1.Descuento()
	pgc1.Prestamo()

	fmt.Println("--------CLASE 02--------------")
	pgc2.CalcularPromedio()
	pgc2.ImpuestosSalario()
	pgc2.Tutorial()
	
}

