package functions

import (
	"fmt"
	"strings"
)

func Convert(text string)(string, string){
	minuscula:=strings.ToLower(text)
	mayuscula:=strings.ToUpper(text)
	fmt.Println(minuscula, mayuscula)
	
	return minuscula, mayuscula
}




 

 