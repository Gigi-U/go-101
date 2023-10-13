package functions

import (
	"errors"
	"fmt"
	"os"
)

// en Go los errores se manejan en el momento en que se presentan a diferencia de otros lenguajes que se usan excepciones.
//* necesitamos el paquete os. ReadFile me pide como argumento el nombre de la ruta donde tengo el archivo (en el caso de ejemplo pongo cualquier cosa para forzar el error).


func ManejoErrores(){
	//([]byte, error) --> content, err	
	//content, err := os.ReadFile("./unexistingFile.txt") /*con error*/
		content, err := os.ReadFile("./functions/hello.txt") /*archivo existente y path correcto*/

	// lo que hago acá es decir q si hay un error
	if err != nil{
		fmt.Printf("Ocurrió un error %v",err)
		return
	} // como la funcion me retorna un slice de bytes. entonces le hago un casteo
		fmt.Println(string(content))
	
}

func ErrorDivision(){
	result, err :=division(10,10)
	if err != nil{
		fmt.Printf("Ocurrió un error %v",err)
		return
	} // como la funcion me retorna un slice de bytes. entonces le hago un casteo
		fmt.Println(result)
}
func division(dividendo, divisor int)(int, error){
	if divisor == 0{
		return 0, errors.New("No puedes dividir por cero")
	}
	return dividendo / divisor, nil
}


/* 
Starting from Go 1.16, ioutil.ReadAll, ioutil.ReadFile and ioutil.ReadDir are deprecated, as the package io/ioutil is deprecated.
SO:
ioutil.ReadAll -> io.ReadAll
ioutil.ReadFile -> os.ReadFile
ioutil.ReadDir -> os.ReadDir

Others:
ioutil.NopCloser -> io.NopCloser
ioutil.TempDir -> os.MkdirTemp
ioutil.TempFile -> os.CreateTemp
ioutil.WriteFile -> os.WriteFile */