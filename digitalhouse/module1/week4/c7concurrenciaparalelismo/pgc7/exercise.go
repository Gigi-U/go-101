package pgc7

import (
	"fmt"

)

/* Calcular precio

Una empresa nacional se encarga de realizar venta de productos, servicios y mantenimiento.
Para ello requieren realizar un programa que se encargue de calcular el precio total de Productos, Servicios y Mantenimientos.
Debido a la fuerte demanda, y para optimizar la velocidad, requieren que el cálculo de la sumatoria se realice en paralelo mediante tres goroutines.

Se requieren tres estructuras:
	- Productos: nombre, precio, cantidad.
	- Servicios: nombre, precio, minutos trabajados.
	- Mantenimiento: nombre, precio.

Y se requieren tres funciones:
	- Sumar Productos: recibe un array de producto y devuelve el precio total (precio por cantidad).
	- Sumar Servicios: recibe un array de servicio y devuelve el precio total (precio por media hora trabajada. Si no llega a trabajar 30 minutos se le cobra como si hubiese trabajado media hora).
	- Sumar Mantenimiento: recibe un array de mantenimiento y devuelve el precio total.

Las tres se deben ejecutar concurrentemente y, al final, se debe mostrar por pantalla el monto final (sumando el total de las tres).

*/
// Estructura Productos
type Productos struct{
	Nombre			string 	`json:"nombre"`
	Precio 			float32 `json:"precio"`
	Cantidad 		int 	`json:"cantidad"`
} 
// Estructura Servicios
type Servicios struct{
	Nombre 			string 	`json:"nombre"`
	Precio 			float32 `json:"precio"`
	MinutoTrabajado int		`json:"minutoTrabajado"`
} 
// Estructura Mantenimiento
type Mantenimiento struct{
	Nombre			string 	`json:"nombre"`
	Precio 			float32 `json:"precio"`
} 

// función sumarProductos
func sumarProductos(productosArray []Productos)(float32, error){
	var precioTotal float32


	return float32(precioTotal),nil		
}

// función sumarServicios
func sumarServicios(serviciosArray []Servicios)(float32, error){
	var precioTotal float32


	return float32(precioTotal),nil
}
// función sumarMantenimiento
func sumarMantenimiento(mantenimientoArray []Mantenimiento)(float32, error){
	var precioTotal float32


	return float32(precioTotal),nil
}

// equivale a main() si fuera un programa aislado
func MontoFinal(){
	// goroutine sumarProductos
	go sumarProductos(
		[]Productos{
			{
				Nombre:   "Producto1",
				Precio:   10.99,
				Cantidad: 5,
			},
			{
				Nombre:   "Producto2",
				Precio:   20.49,
				Cantidad: 3,
			},
			{
				Nombre:   "Producto3",
				Precio:   15.75,
				Cantidad: 7,
			},
		},		
	)

	
	go sumarServicios(
		[]Servicios{
			{
				Nombre:           "Servicio1",
				Precio:           25.99,
				MinutoTrabajado: 30,
			},
			{
				Nombre:           "Servicio2",
				Precio:           19.49,
				MinutoTrabajado: 45,
			},
			{
				Nombre:           "Servicio3",
				Precio:           35.75,
				MinutoTrabajado: 60,
			},
		},	
	)
	go sumarMantenimiento(
		[]Mantenimiento{
			{
				Nombre: "Mantenimiento1",
				Precio: 50.0,
			},
			{
				Nombre: "Mantenimiento2",
				Precio: 45.0,
			},
		},
	)

	montoFinal:= 0

	fmt.Println(montoFinal)
}