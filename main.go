package main

import (
	// here you import all other packages
	"fmt"

	"github.com/Gigi-U/go-101.git/digitalhouse/module1/c1basis"
	"github.com/Gigi-U/go-101.git/digitalhouse/module1/c1basis/mtc1"
	"github.com/Gigi-U/go-101.git/digitalhouse/module1/c1basis/pgc1"
	"github.com/Gigi-U/go-101.git/digitalhouse/module1/c2functions/mtc2"
	"github.com/Gigi-U/go-101.git/digitalhouse/module1/c2functions/pgc2"
	"github.com/Gigi-U/go-101.git/datatypes"
	"github.com/Gigi-U/go-101.git/functions"
	"github.com/Gigi-U/go-101.git/loopsandcycles"
	"github.com/Gigi-U/go-101.git/operators"
	"github.com/Gigi-U/go-101.git/pointers"
	"github.com/Gigi-U/go-101.git/structures"
	"github.com/Gigi-U/go-101.git/structures/examples"
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

	fmt.Println("--------LOOPS & CYCLES--------------")
	//FOLDER: loopsandcycles 
	fmt.Println("--------FOR --------------")
	loopsandcycles.For()
	fmt.Println("--------IF --------------")
	loopsandcycles.If()
	fmt.Println("--------SWITCH --------------")
	loopsandcycles.Switch()
	fmt.Println("--------WHILE --------------")
	loopsandcycles.While()

	fmt.Println("--------VARIABLES --------------")
	//FOLDER: varandconst
	varandconst.Variables()
	fmt.Println("-------- CONSTANTES--------------")
	varandconst.Const()

	fmt.Println("--------POINTERS--------------")
	//FOLDER: pointers
	pointers.Pointers()

	fmt.Println("\n--------OPERATORS--------------")
	//FOLDER: operators
	operators.Operators()


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
	mtc2.Exercise1()

	fmt.Println("--------FUNCTIONS--------------")
	functions.Hello("Gisela","Urriza\n")
	functions.Change() 
	functions.Sum(2,3)
	functions.Convert("TodO a MINUSculA O mayUSCULa  ")
	functions.ManejoErrores()
	functions.ErrorDivision()
	functions.FiltradorDeValores()
	functions.Hi()
	functions.Sum(1,4)
	functions.Anonima()
	functions.Calcular()
	functions.RetornoValor()

	fmt.Println("--------PRUEBA--------------")

	fmt.Println("--------STRUCTURES--------------")
	structures.Estructuras()
	examples.EstructuraPersona()
	structures.Circ()
	fmt.Println("--------CLASE 03--------------")

}

