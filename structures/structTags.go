package structures

import (
	"encoding/json"
	"fmt"
)
  
type Person struct {
	PrimerNombre string `json:"primer_nombre"`
	Apellido string `json:"apellido"`
	Telefono string `json:"telefono"`
	Direccion string `json:"direccion,omitempty"`
	User Usuario

	//Eliminar campos JSON vacíos. si queremos omitir cualquier resultado cuyo valor sea “cero”, debemos incluir en la etiqueta el atributo omitempty.
}
type Usuario struct {
	Username  string `json:"username"`
	Email  string `json:"email"`
	Password  string `json:"-"`

	//Ignorar campos privados. Cuando un campo es de naturaleza sensible, como en el caso de una contraseña, queremos que el codificador de JSON ignore el campo por completo, incluso cuando esté establecido. Esto se hace usando el valor especial “-” como argumento del valor para una etiqueta de estructura JSON.

 }
 
  
func Etiquetas() {

	p := Person{"Juan", "Perez", "43434343", "Calle falsa 123", Usuario{"GiiDev","giidev@gmail.com",""}}
	miJSON, err := json.Marshal(p)


	fmt.Println("------------Imprimiendo JSON")
	fmt.Println(string(miJSON))

	fmt.Println("------------Error")
	fmt.Println(err)

	fmt.Println("------------Imprimiendo Usuario")
	fmt.Println(p.User)
}
  