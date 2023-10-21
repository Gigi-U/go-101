package mtc4

import (
	"encoding/json"
	"fmt"
)

/*
 Productos
Algunas tiendas de e-commerce necesitan realizar una funcionalidad en Go para
administrar productos y RETORNAR el VALOR delPRECIO TOTAL.

La empresa tiene tres tipos de productos: Pequeño, Mediano y Grande. Pero se espera que sean muchos más. 👉🏼(INTERFAZ)
Los costos adicionales para cada uno son:
● Pequeño: solo tiene el costo del producto. 👉🏼 return precio
● Mediano: el precio del producto + un 3% por mantenerlo en la tienda.👉🏼 precio = precio + (precio*.03)     return precio
● Grande: el precio del producto + un 6% por mantenerlo en la tienda y un adicional
de $2500 de costo de envío. 👉🏼 precio = precio + (precio*.03) + 2500   return precio
El porcentaje de mantenerlo en la tienda se calcula sobre el precio del producto.

Se requiere una función factory que reciba el tipo de producto y el precio, y retorne una
interfaz Producto que tenga el método Precio.

Se debe poder ejecutar el método Precio y que el método devuelva el precio total
basándose en el costo del producto y los adicionales, en caso de que los tenga.

👉🏼 INTERFAZ + METODO con switch interno?
*/


type Product struct{
	Price 	float32 	`json:"price"`
	Size	string	 	`json:"size"`
}


type priceFactory interface {
	setPrice() float32
}

func (p Product) setPrice() float32 {
	switch p.Size {
	case "small":
		return p.Price
	case "medium":
		p.Price = p.Price + (p.Price * 0.03)
		return p.Price
	case "large":
		p.Price = p.Price + (p.Price * 0.03) + 2500
		return p.Price
	default:
		return 0
	}
}

// Función factory que crea un producto basado en el tipo y precio proporcionados.
func finalProductPrice(price float32, size string) Product {
	return Product{Price: price, Size: size}
}

// equivale a main()
func PrecioTotalProductos(){
	product1 := finalProductPrice(3000, "small")
	product2 := finalProductPrice(13000, "medium")
	product3 := finalProductPrice(53000, "large")

	fmt.Printf("Total Price small product: $%.2f\n", product1.setPrice())
	fmt.Printf("Precio total del producto Mediano: $%.2f\n", product2.setPrice())
	fmt.Printf("Precio total del producto Grande: $%.2f\n", product3.setPrice())

	//Structure Tags
	miJSON, err := json.Marshal(product1)
	fmt.Println("\n------------Imprimiendo JSON")
	fmt.Println(string(miJSON))
	fmt.Println("\n------------Error")
	fmt.Println(err)
}
/*
*Terminal
Total Price small product: $3000.00
Precio total del producto Mediano: $13390.00
Precio total del producto Grande: $57090.00

------------Imprimiendo JSON
{"price":3000,"size":"small"}

------------Error
<nil>
*/