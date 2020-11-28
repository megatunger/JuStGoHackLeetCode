package main

import (
	"fmt"
)

func angryProfessor(k int32, a []int32) string {
	var count = int32(0)
	for _, element := range a {
		if element <= 0 {
			count += 1
		}
	}
	if count < k {
		return "YES"
	} else {
		return "NO"
	}
}



func main() {
	fmt.Println(angryProfessor(3, []int32{-1, -3, 4, 2}))
}