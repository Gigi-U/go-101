package mtc3

import (
	"fmt"
)

// Product Structure
type Product struct{
	ID	int
	Name string
	Price float32
	Description string
	Category string
} 

// Global Slice from Product 
var products = []Product {
	{1,"Soccer Ball", 4000,"Champion Sports Extreme Series Composite Soccer Ball: Sizes 3","Sports"},
	{2,"Barbie The Movie Doll", 10000,"Margot Robbie as Barbie, Collectible Doll Wearing Pink and White Gingham Dress with Daisy Chain Necklace for 6 years and up","Toys"},
	{3,"Wilson Racket", 3000,"Wilson US Open Junior/Youth Recreational Tennis Rackets","Sports"},
}

// or main() in a single folder application
func MyProducts(){
	// new product data
	newProduct := Product{4,"Rail Chief Train", 15000,"Bachmann Trains - Ready To Run 130 Piece Electric Train Set","Toys"} 
	newProduct.Save()	

	// print All Products
	fmt.Println("\nProducts List:")
	GetAll()
	fmt.Println("---------------------------------------------------")

	// find product by ID
	//productID := 6 // WITH ERROR
	productId := 4 // CORRECT

	p, err := getById(productId)
	if err != nil {
		fmt.Printf("Error: %v\n", err)
	} else {
		fmt.Println("\n¡Product found!:")
		ProductDetails(p)
	}
}

//Method that saves a new product and adds it to the slice
func (newProduct Product)Save(){	
	products = append(products, newProduct)
	return
}

//! Method to get All Products  - no se ve como un método.... 
func GetAll()[]Product{
	for _, p := range products {
		ProductDetails(p)
	}
	return products
}

// function to get Product by Id or error
func getById(id int) (Product, error) {
	for _, p := range products {
		if p.ID == id {
			return p, nil
		}
	}
	return Product{}, fmt.Errorf("\nProduct with ID %d not found", id)
}


// Function to print Products with details - Not requested
func ProductDetails(p Product ) {
	fmt.Printf("\nID: %d\n", p.ID)
	fmt.Printf("\tName: %s\n", p.Name)
	fmt.Printf("\tPrice: %.2f\n", p.Price)
	fmt.Printf("\tDescription: %s\n", p.Description)
	fmt.Printf("\tCategory: %s\n", p.Category)
}

/* Ejercicio
Crear un programa que cumpla los siguientes puntos:
Tener una estructura llamada Product con los campos ID, Name, Price, Description y Category.
Tener un slice global de Product llamado Products, instanciado con valores.
Dos métodos asociados a la estructura Product: Save(), GetAll().
	-El método Save() deberá tomar el slice de Products y añadir el producto desde el cual se llama al método.
	-El método GetAll() deberá imprimir todos los productos guardados en el slice Products.
Una función getById() a la cual se le deberá pasar un INT como parámetro y retorna el producto correspondiente al parámetro pasado.
Ejecutar al menos una vez cada método y función definidos desde main().
*/
