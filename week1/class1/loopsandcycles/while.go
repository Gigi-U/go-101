package loopsandcycles

import (
	"fmt"
)
//* Es en realidad un *FOR CONTINUO*
func While() {
	sum := 1
	for sum < 10 {
	sum += sum
	}
	fmt.Println(sum)	
}

