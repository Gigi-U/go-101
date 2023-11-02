package prueba

import (
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

func Main() {
	writeCsvFile()
	categorias := readCsvFile()
	calcularEstimaciones(&categorias)
	leerEstimaciones()


}

type Producto struct {
	codigo   int32
	nombre   string
	precio   float32
	cantidad int32
}

func (p *Producto) stringify() string {
	return fmt.Sprintf("%d!%s!%.2f!%d", p.codigo, p.nombre, p.precio, p.cantidad)
}

type Categoria struct {
	codigo         int32
	nombre         string
	listaProductos []Producto
}

func (p *Categoria) stringify() string {
	productosString := ""
	for i, v := range p.listaProductos {
		if i != 0 {
			productosString += "/"
		}
		productosString += v.stringify()
	}

	return fmt.Sprintf("%d,%s,'%s'", p.codigo, p.nombre, productosString)
}

var producto1 = Producto{1, "Algo1", 45000, 5}
var producto2 = Producto{2, "Algo2", 45000, 5}
var producto3 = Producto{3, "Algo3", 45000, 5}
var producto4 = Producto{4, "Algo4", 45000, 5}

var categoria1 = Categoria{1, "Electrodomesticos", []Producto{producto1, producto2}}
var categoria2 = Categoria{2, "Muebles", []Producto{producto3, producto4}}

func writeCsvFile() {
	categoriasSlices := []Categoria{categoria1, categoria2}
	var categoriasString string
	for i, v := range categoriasSlices {
		if i != 0 {
			categoriasString += ";"
		}
		categoriasString += v.stringify()
	}

	err := os.WriteFile("prueba/file.csv", []byte(categoriasString), 0666)
	if err != nil {
		log.Fatal("Error al escribir el archivo: ", err)
	}
}

func readCsvFile() []Categoria {
	categoriasReturn := []Categoria{}
	fileBytes, err := os.ReadFile("prueba/file.csv")
	if err != nil {
		log.Fatal("Error al leer el archivo: ", err)
	}
	fileString := fmt.Sprintf("%s", fileBytes)
	categoriaSlice := strings.Split(fileString, ";")
	for _, v := range categoriaSlice {
		categoriaAtributos := strings.Split(v, ",")
		categoriaCodigo, _ := strconv.Atoi(categoriaAtributos[0])
		categoria := Categoria{int32(categoriaCodigo), categoriaAtributos[1], []Producto{}}
		productoSlice := strings.Split(strings.Trim(categoriaAtributos[2], "'"), "/")
		for _, v := range productoSlice {
			productoAtributos := strings.Split(v, "!")
			productoCodigo, _ := strconv.Atoi(productoAtributos[0])
			productoPrecio, _ := strconv.ParseFloat(productoAtributos[2], 32)
			productoCantidad, _ := strconv.Atoi(productoAtributos[3])
			productoCompleto := Producto{int32(productoCodigo), productoAtributos[1], float32(productoPrecio), int32(productoCantidad)}
			categoria.listaProductos = append(categoria.listaProductos, productoCompleto)
		}
		categoriasReturn = append(categoriasReturn, categoria)
	}
	fmt.Println(categoriasReturn)
	return categoriasReturn
}

func calcularEstimaciones(categorias *[]Categoria) {
	estimacionesString := "Categoria,EstimativoPorCategoria"
	for _, categoria := range *categorias {
		totalEstimacion := 0.0
		for _, producto := range categoria.listaProductos {
			totalEstimacion += float64(producto.precio) * float64(producto.cantidad)
		}
		estimacionesString += ";" + fmt.Sprintf("%s,%.2f", categoria.nombre, totalEstimacion)
	}
	err := os.WriteFile("prueba/estimaciones.csv", []byte(estimacionesString), 0666)
	if err != nil {
		log.Fatal("Error al escribir el archivo de estimaciones: ", err)
	}
}

func leerEstimaciones() {
	estimacionesString, err := os.ReadFile("prueba/estimaciones.csv")
	if err != nil {
		log.Fatal("Error al leer el archivo de estimaciones: ", err)
	}
	estimacionesSlice := strings.Split(string(estimacionesString), ";")
	for i, v := range estimacionesSlice {
		estimacionesValues := strings.Split(v, ",")
		if i != 0 {
			estimacionFloat, _ := strconv.ParseFloat(estimacionesValues[1], 32)
			fmt.Printf("%s\t\t%.2f\n", estimacionesValues[0], estimacionFloat)
		} else {
			fmt.Printf("%s\t\t%s\n", estimacionesValues[0], estimacionesValues[1])
		}
	}
}


