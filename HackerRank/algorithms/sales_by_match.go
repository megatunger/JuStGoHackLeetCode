package main

import (
	"fmt"
)

func sockMerchant(n int32, ar []int32) int32 {
	var m = make(map[int32]int32)
	var count = 0
	for _, element := range ar {
		if m[element] != 0 {
			count += 1
			m[element] = 0
		} else {
			m[element] = 1
		}
	}
	return int32(count)
}

func main() {
	fmt.Println(sockMerchant(9, []int32{10, 20, 20, 10, 10, 30, 50, 10, 20}))
}