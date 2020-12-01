package main

import (
	"fmt"
)

func findDigits(n int32) int32 {
	var count = int32(0)
	var origin = n
	for n!=0 {
		var remain = n % 10
		if remain!=0 && origin%remain==0 {
			count += 1
		}
		n = n / 10
	}
	return count
}

func main() {
	fmt.Println(findDigits(1012))
}