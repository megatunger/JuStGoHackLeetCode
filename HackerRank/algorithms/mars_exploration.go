package main

import (
	"fmt"
)

func marsExploration(s string) int32 {
	var compare = "SOS"
	var result = 0
	for position, char := range s {
		if int32(compare[position%3]) != char {
			result += 1
		}
	}
	return int32(result)
}


func main() {
	fmt.Println(marsExploration("SOSSPSSQSSOR"))
}