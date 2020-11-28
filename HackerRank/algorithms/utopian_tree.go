package main

import (
	"fmt"
)

func utopianTree(n int32) int32 {
	var height = int32(1)
	for i:=int32(1); i<=n; i++ {
		if i%2 != 0 {
			height = height * 2
		} else {
			height += 1
		}
	}
	return height
}

func main() {
	fmt.Println(utopianTree(4))
}
