package main

import (
	"fmt"
)

func bonAppetit(bill []int32, k int32, b int32) {
	var sum = int32(0)
	for i, element := range bill {
		if int32(i) != k {
			sum += element
		}
	}
	var fairDivide = sum/2
	if b == fairDivide {
		fmt.Println("Bon Appetit")
	} else {
		fmt.Println(b - fairDivide)
	}
}

func main() {
	bonAppetit([]int32{3, 10, 2, 9}, 0, 12)
}