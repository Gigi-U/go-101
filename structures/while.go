package structures

import (
	"fmt"
)

func While() {
	sum := 1
	for sum < 10 {
	sum += sum
	}
	fmt.Println(sum)	
}

