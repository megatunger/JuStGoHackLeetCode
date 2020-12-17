package main

import (
	"fmt"
	"math"
)

func serviceLane(n int32, cases [][]int32, width[]int32) []int32 {
	var result []int32
	for _, row := range cases {
		var minNum = int32(math.MaxInt32)
		for i:= row[0]; i<=row[1]; i++ {
			if width[i] < minNum {
				minNum = width[i]
			}
		}
		result = append(result, minNum)
	}
	return result
}


func main() {
	fmt.Println(serviceLane(8, [][]int32{{2, 3}, {1, 4}, {2, 4}, {2, 4}, {2, 3}}, []int32{1, 2, 2, 2, 1}))
}