package main

import (
	"fmt"
	"math"
)

func equalizeArray(arr []int32) int32 {
	dict:= make(map[int32]int)
	var max = int32(math.MinInt32)
	for _, num :=  range arr {
		dict[num] = dict[num] + 1
		if int32(dict[num]) > max {
			max = int32(dict[num])
		}
	}
	return int32(int32(len(arr)) - max)
}

func main() {
	fmt.Println(equalizeArray([]int32{3, 3, 2, 1, 3}))
}