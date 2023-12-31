package main

import (
	// here you import all other packages
	"fmt"

	"github.com/Gigi-U/go-101.git/composition"
	"github.com/Gigi-U/go-101.git/datatypes"
	"github.com/Gigi-U/go-101.git/digitalhouse/module1/week1/c1basis"
	"github.com/Gigi-U/go-101.git/digitalhouse/module1/week1/c1basis/mtc1"
	"github.com/Gigi-U/go-101.git/digitalhouse/module1/week1/c1basis/pgc1"
	"github.com/Gigi-U/go-101.git/digitalhouse/module1/week1/c2functions"
	"github.com/Gigi-U/go-101.git/digitalhouse/module1/week1/c2functions/mtc2"
	"github.com/Gigi-U/go-101.git/digitalhouse/module1/week1/c2functions/pgc2"
	"github.com/Gigi-U/go-101.git/digitalhouse/module1/week2/c3structures"
	"github.com/Gigi-U/go-101.git/digitalhouse/module1/week2/c3structures/mtc3"
	"github.com/Gigi-U/go-101.git/digitalhouse/module1/week2/c3structures/pgc3"
	"github.com/Gigi-U/go-101.git/digitalhouse/module1/week2/c4ComPuntInterf"
	"github.com/Gigi-U/go-101.git/digitalhouse/module1/week2/c4ComPuntInterf/mtc4"
	"github.com/Gigi-U/go-101.git/digitalhouse/module1/week2/c4ComPuntInterf/pgc4"
	"github.com/Gigi-U/go-101.git/digitalhouse/module1/week3/c6"
	"github.com/Gigi-U/go-101.git/digitalhouse/module1/week3/c6/mt6"
	"github.com/Gigi-U/go-101.git/digitalhouse/module1/week4/c7concurrenciaparalelismo/mtc7"
	"github.com/Gigi-U/go-101.git/digitalhouse/module2/week5/c9jsonapirest/pgc9"
	"github.com/Gigi-U/go-101.git/functions"
	"github.com/Gigi-U/go-101.git/interfaces"
	"github.com/Gigi-U/go-101.git/loopsandcycles"
	"github.com/Gigi-U/go-101.git/operators"
	"github.com/Gigi-U/go-101.git/pointers"
	"github.com/Gigi-U/go-101.git/prueba"
	"github.com/Gigi-U/go-101.git/structures"
	"github.com/Gigi-U/go-101.git/structures/examples"
	"github.com/Gigi-U/go-101.git/varandconst"
)

func main() {
	fmt.Println("---------CLASS 01---------------")
	fmt.Println("--------------------------------")

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

	fmt.Println("\n--------OPERATORS--------------")
	//FOLDER: operators
	operators.Operators()

	fmt.Println("--------Homework C1--------------")
	c1basis.Class1()
	mtc1.Exercise1()
	mtc1.Exercise2()
	pgc1.Descuento()
	pgc1.Prestamo()

	fmt.Println("---------CLASS 02---------------")
	fmt.Println("--------------------------------")

	fmt.Println("--------FUNCTIONS--------------")
	// funcioness con parámetro
	functions.Hello("Gisela","Urriza\n")
	// funcioness con parámetro
	functions.Change() 
	// Funciones con retorno
	functions.Sum(2,3)
	// Funciones Multi retorno
	functions.Convert("TodO a MINUSculA O mayUSCULa  ")
	// Manejo de errores
	functions.ManejoErrores()
	// Manejo de errores
	functions.ErrorDivision()
	// FUNCION QUE RECIBE FUNCIONES
	functions.FiltradorDeValores()
	// FUNCION QUE RETORNA FUNCION
	functions.Hi()
	// Funciones con retorno
	functions.Sum(1,4)
	// funciones anonimas
	functions.Anonima()
	// Función con ellipsis
	functions.Calcular()
	// FUNCION QUE RETORNA VALOR
	functions.RetornoValor()

	fmt.Println("--------Homework C2--------------")
	pgc2.CalcularPromedio()
	pgc2.ImpuestosSalario()
	pgc2.Tutorial()
	mtc2.Exercise1()
	c2functions.Class2()

	fmt.Println("---------CLASS 03---------------")
	fmt.Println("--------------------------------")

	fmt.Println("--------STRUCTURES--------------")
	// Structures  - example:
	structures.Estructuras()
	// Estructuras:
	examples.EstructuraPersona()
	// Structure TAGS:
	structures.Etiquetas()
	// METHODS:
	structures.Circ()
	// Composición o Embedding Structs
	structures.Rodados()

	fmt.Println("--------Homework C3--------------")
	pgc3.Alumni()
	pgc3.VideoC3()
	mtc3.MyProducts()
	c3structures.Class3()

	fmt.Println("--------CLASS 04--------------")
	fmt.Println("--------------------------------")
	fmt.Println("---COMPOSITION - POINTERS - INTERFACES---")
	composition.HerenciaGo()
	interfaces.EmptyInterface()
	interfaces.InterfacesGo()
	interfaces.TypeAssertion()
	c3structures.CompaniaEmpleado()
	//FOLDER: pointers
	pointers.Pointers()
	pointers.PointersGo()
	fmt.Println("--------Homework C4--------------")
	c4compuntinterf.Class4()
	pgc4.InfoUsuario()
	pgc4.PersonnelManagement()
	mtc4.PrecioTotalProductos()

	fmt.Println("--------CLASS 06--------------")
	c6.MesaC6()
	mt6.Mayorista()
	fmt.Println("--------CLASS 07--------------")
	mtc7.EvenOrOdd()


	fmt.Println("--------CLASS 09--------------")
	fmt.Println("\n----------FUNCTION MARSHAL---------------")
	pgc9.Marshal()
	fmt.Println("\n----------FUNCTION UNMARSHAL---------------")
	pgc9.UnMarshal()
	pgc9.Marshal()
	pgc9.NetHttp()

	fmt.Println("\n--------PRUEBA--------------")
	prueba.Main()

}

