package structures

import (
	"fmt"
)
func Switch() {
  
	// FORMA 1 
		emoji:= "🐶"

		switch emoji{
		case  "🐶":
			fmt.Println("Es un perrito!")		
		case "🍎":
			fmt.Println("Es una manzana!")
		default:
			fmt.Println("Es un/a: ", emoji)
		}

		// FORMA 2  - evitar la repeticion

		switch emoji{
		case  "🐶", "🐱":
			fmt.Println("Es un animal!")		
		case "🍎", "🍇":
			fmt.Println("Es una fruta!")
		default:
			fmt.Println("No es animal ni fruta!")
		}


		// FORMA 3  - CON COMPARADORES LÓGICOSn

		switch {
		case emoji ==  "🐶" || emoji == "🐱":
			fmt.Println("Es un animal!")		
		case emoji ==  "🍎"|| emoji ==  "🍇":
			fmt.Println("Es una fruta!")
		default:
			fmt.Println("No es animal ni fruta!")
		}


		//!FORMA 3 - con declaración corta  (no me deja meter el var adelante)

		switch day:= 4; day {
	    case 1:
		    fmt.Println("Monday")
	    case 2:
		    fmt.Println("Tuesday")
	    case 3:
		    fmt.Println("Wednesday")
	    case 4:
		    fmt.Println("Thursday") 
	    case 5:
		    fmt.Println("Friday")
	    case 6:
		    fmt.Println("Saturday")
	    default:
		    fmt.Println("Sunday")
    }  

	 	//FORMA 4 - con fallthrough
		 num := 0
		 switch num { 
		   case 0:
			 fmt.Println("Case: 0")
			 fallthrough   // SE EJECUTA EL CASE 1 TAMBIÉN ( si el numero es 1 se van a ejecutar el caso 1 y 2)
		   case 1:
			 fmt.Println("Case: 1")
		   case 2:
			 fmt.Println("Case: 2")
		   default:
			 fmt.Println("default")
		 }

		fmt.Println()

}

