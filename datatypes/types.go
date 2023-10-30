package datatypes

import (
	"fmt"
)

func DataTypes() {
	numero1 := -100		// o var numero 1 = -100 | valores positivos negativos si no tengo nunguna variable antes me lo va a tomar comi int. sino ojo porque va a tomar la propiedad de la anterior

	var numero2 uint	// solo positivos

	var numero3 int8	// -127 a 127 (255 nros enteros) = 8 bits de memoria. si no definimos: go define entre int16 int32 y int64

	var numero4 float32	// numeros reales de 32 bits (hta 24 digitos)

	var numero5 float64	// numeros reales de 64 bits (hta 53 digitos)

	var nombre string // incluye caracteres alfanumericos 

	var apellido string

	var booleano1 bool

	var booleano2 bool

	var miByte byte	// 1 byte = 8 bits == 0 a 255

	var miRune rune 

	var suma int = 0

	numero2 = 2540
	numero3 = 125
	numero4 = 111111111111111111111111111111111111111.11
	numero5 = 11111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111.11
	nombre, apellido = "Cosme", "Fulanito"
	booleano1 = true
	booleano2 = false
	miByte = 255 // es un alias para un uint8 - internamente se va a mostrar como un uint8

	var integer int = 7
	var float = float32(integer) // Go permite realizar Typecasting o encasillamiento de tipos, que es que podemos convertir el valor de una variable de un tipo a otra que sea compatible. float = decimales ej 23.45

	miRune  = 2147483647 // es un alias para un int32 - internamente se va a mostrar como un int32. ADEMAS rune representa un código en Unicode. si le pasamos la comilla simple y le indicamos un caracter ej. var a rune = 'a' (si veo en la tabla de unicode, me va a mistrar que la letra a es el codigo 97) ver: https://en.wikipedia.org/wiki/List_of_Unicode_characters

	var _ = 234 // si quiero crear una variable y tener listo el valor,  pero no quiero declararla (darle nombre) utilizo el IDENTIFICADOR BLANK	
	var _ string = "234"

	fmt.Println( numero1, numero2, numero3,numero4,numero5, nombre, apellido, booleano1, booleano2, miByte, integer, float, miRune, suma)

}


func ConPrintf(){
	booleano1 := true
	fmt.Printf("Tipo: %T ; Valor: %v ",booleano1, booleano1)
}

/* fmt.Printf("") me permite imprimir cosas en pantalla, con formato.
 Dentro de las comillas quiero imprimir tipo de dato y su valor. En la cadena de caracteres van los VERBOS, para que Go sepa lo q tiene que hacer. Los verbos se identifican precedidos de % y seguidos de una letra. La T significa: tipo. v es para poder imprimir el valor. %q se usa para textos, que me hace poner comillas. Ej: fmt:Printf("%q", nombre). en booleanos el valor 0 es FALSE */

func DifTiposDatos(){
	var numero2 uint= 2540
	var numero3 int8= 125
	suma:=0
	suma =  int(numero2) + int(numero3) // antes numero2 era uint y numero3 era int8 
	fmt.Printf("Tipo: %T ; Valor: %v\n",suma, suma)
}// si quiero sumar diferentes tipos de datos Go no me lo va a permitir. pero puede solucionarse con el casting
