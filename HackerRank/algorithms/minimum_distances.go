package main

import (
	"fmt"
	"math"
)

func minimumDistances(a []int32) int32 {
	var minResult = int32(math.MaxInt32)
	dict:= make(map[int32]int)
	for pos, num :=  range a {
		_, exist := dict[num]
		if !exist {
			dict[num] = pos
		} else {
			var distance = float64(pos-dict[num])
			if int32(math.Abs(distance)) < minResult {
				minResult = int32(distance)
			}
		}
		fmt.Println(num, dict)
	}
	if minResult == int32(math.MaxInt32) {
		return -1
	}
	return int32(math.Abs(float64(minResult)))
}

func main() {
	fmt.Println(minimumDistances([]int32{7, 1, 3, 4, 1, 7}))
	fmt.Println(minimumDistances([]int32{3, 2, 1, 2, 3}))
	fmt.Println(minimumDistances([]int32{1, 1}))
}
