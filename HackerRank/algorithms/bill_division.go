package main

import (
	"fmt"
)

func bonAppetit(bill []int32, k int32, b int32) {
	var sum = int64(0)
	for _, element := range bill {
		if bill[k] != element {
			sum += int64(element)
		}
	}
	var fairDivide = int32(sum/2)
	if b == fairDivide {
		fmt.Println("Bon Appetit")
	} else {
		fmt.Println(b - fairDivide)
	}
}

func main() {
	bonAppetit([]int32{3, 10, 2, 9}, 0, 12)
}