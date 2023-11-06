package pgc9

import (
	"encoding/json"
	"fmt"
	"log"
)

type product struct {
	Name      string
	Price     int
	Published bool
}

func Marshal() {
	p := product{
		Name:      "MacBook Pro",
		Price:     1500,
		Published: true,
	}

	jsonData, err := json.Marshal(p)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(string(jsonData))

}
func UnMarshal(){
	jsonData := `{"Name": "MacBook Air", "Price": 900, "Published":	true}`

	//jsonData := `{"Name": "MacBook Air", "Price": "Un Peso" , "Published":true}` con error

	var p product

	if err := json.Unmarshal([]byte(jsonData), &p); err != nil {
		log.Fatal(err)
		log.Printf("ERROR: %v",err)
	}
}

/* 
* MARSHAL
La función func Marshal(v interface{}) ([]byte, error) toma como parámetro un valor de cualquier tipo, y retorna una slice de bytes que contiene su representación en formato JSON. 
También retorna un error en caso de encontrar uno.

	En caso de querer usar json.Marshal( ) y pasar un struct como parámetro, los campos a transformar a JSON tienen que estar exportados, es decir, en mayúsculas.

* UNMARSHAL
La función func Unmarshal(data []byte, v interface{}) error  recibe como primer parámetro un array de bytes y, como segundo parámetro, un puntero a un struct. Si el array de bytes es data en JSON, entonces la función unmarshal() va a tratar de decodificarlo y llenar el struct con esos datos. La función devuelve un error, en caso de encontrar uno.	
 */