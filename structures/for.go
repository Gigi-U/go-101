package structures

import (
	"fmt"

)
func For() {

	fmt.Println("--------FOR SIMPLE-------------")	
	for i := 0; i < 10; i++{
		fmt.Println(i)
		//IMPRIME HASTA 10
	}

	fmt.Println("----FOR CONTINUO(tmb se lo llama While)----")	
	i:= 1
	for i <= 10{
	   fmt.Println(i)
	   i++
   	}

	fmt.Println("----FOR RANGE SLICE----")	
	nums:=[]uint8{2,4,6,8}
	// aquí solo imprimo 4 veces el string
	//for range nums{
	//	fmt.Println("imprimo")
	//}
	// aquí imprimo índice y valor del array
	for i,v:= range nums{
	fmt.Println("Indice: ",i,"Valor: ",v)
	}

	// si quiero que los valores se multipliquen x 2
	// _,v -> como no uso la variable i puedo usar el blank
	for i := range nums{
		nums[i] *=2
	fmt.Println(nums)
	}
	fmt.Println("----FOR MAP----")	
	sports:=map[string]string{"futbol":"⚽️","basket":"🏀"}
	for key, v := range sports{
		fmt.Println("Clave: ",key,"Valor: ",v)
	}

	// si quiero que se imprima cada uno de los caracteres de un string
	fmt.Println("----FOR RANGE STRING----")	
	hello:="hello"
	for _, v := range hello{
		fmt.Println(string(v))
	}


	fmt.Println("----------FRUTAS-----------")
	frutas:=[]string{"manzana",	"banana", "pera"}

	for i, fruta := range
	frutas {
	fmt.Println(i, fruta)
	}

	fmt.Println("------ES IMPAR- Saltar a la siguiente iteración--------")	
	for i := 0; i < 10; i++{
	if i % 2 == 0 {
	continue
	}
	fmt.Println(i, "es impar")
	}


	fmt.Println("--------ROMPER BUCLE-------------")	

	sum := 0
	for {
	sum++
	if sum >= 1000 {
	break
	}
	}
	fmt.Println(sum)

/* 	fmt.Println("--------BUCLES INFINITOS-------------")	

	sum := 0
	for true {
	sum++
	}

   fmt.Println("--------FOR FOREVER (infinito)--------")	
   i:= 1
   for {
	  fmt.Println(i)
	  i++
  	}
	//* los Bucles infinitos me sirven para procesos que deben ser escuchados permanentementes, como sockets. También la programación en automatización y robótica puede basarse en un bucle infinito, como por ejemplo la función void loop() en Arduino, en este caso el bucle infinito deja de ser un error para pasar a ser una virtud, ya que puede reconocerse el estado de un sensor generando una respuesta miles de veces por segundo.
	*/


}

