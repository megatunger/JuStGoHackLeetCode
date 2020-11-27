package main

import (
	"fmt"
	"math"
)

func getMoneySpent(keyboards []int32, drives []int32, b int32) int32 {
	var i = 0
	var max = int32(math.MinInt32)
	for _, keyboard := range keyboards {
		for _, drive := range drives {
			if keyboard + drive <= b {
				i++
				if keyboard + drive > max {
					max = keyboard + drive
				}
			}
		}
	}
	if i == 0 {
		return -1
	}
	return max
}

func main() {
	fmt.Println(getMoneySpent([]int32{3, 1}, []int32{5, 2, 8}, 10))
}