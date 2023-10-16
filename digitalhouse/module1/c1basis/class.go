package c1basis

import "fmt"

func Class(){

	palabra:="abcdefghi"

	fmt.Println("la longitud de la palabra es de : ", len(palabra))

	for i:=0;i<len(palabra);i++{
		//fmt.Printf("%s\n", string (palabra[i]))
		//fmt.Printf("%d:%s\n", i, string (palabra[i]))
		fmt.Printf("%d - %s\n", i, string (palabra[i]))
/* 		ME MUESTRA:
			la longitud de la palabra es de :  9
			0 - a
			1 - b
			2 - c
			3 - d
			4 - e
			5 - f
			6 - g
			7 - h
			8 - i */
	}

	//& Decir mes x número
	// FORMA 1
	mes := 3
	switch mes {
	case 1:
		fmt.Println("Enero")
	case 2:
		fmt.Println("Febrero")
	case 3:
		fmt.Println("Marzo")	
	default:
		fmt.Println("Ese mes no existe")	
	}
	// FORMA 2
	mapMes := map[int]string{
		1:"Enero",
		2:"Febrero",
		3:"Marzo",
	}

	resultado, ok := mapMes[4]
	if !ok{ 
		fmt.Println("el mes no existe")
	}else {
		fmt.Println(resultado)
	}
}