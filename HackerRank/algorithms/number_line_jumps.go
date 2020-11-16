package main

import (
	"fmt"
)

func kangaroo(x1 int32, v1 int32, x2 int32, v2 int32) string {
	if v2-v1==0 {
		return "NO"
	}
	if v1 > v2 {
		var remainder = (x1-x2) % (v2-v1)
		if remainder == 0 {
			return "YES"
		}
	}
	return "NO"
}


func main() {
	fmt.Println(kangaroo(100000, 10000, 200, 1))
}