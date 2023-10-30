package functions

import "fmt"

//& FUNCION QUE RECIBE FUNCIONES
func FiltradorDeValores(){
	nums:=[]int{1,4,60,33,12,5,8} //(slice de enteros)

	result:= filtro(nums, func(num int) bool {
		if num <=10 {
			return true
		}
		return false
	})
	fmt.Println(result)
}

func filtro(nums []int, callback func(int)bool)[]int{
//callback (o lo llamo como quiera): funcion que me permitirá el filtrado. recibe un entero y retorna un boolean y retorna un slice de enteros	
	result := []int{}
	for _,v := range nums{
		if callback(v){
			result = append(result,v)
		}
	}
	return result	
}

//& FUNCION QUE RETORNA FUNCION
func Hi(){
	x:=hello("Gise")("Urriza")
	fmt.Println(x)
}

func hello(name string)func(string)string{
	return func(text string) string {
		return "hello "+ name +" "+ text
	}
}

