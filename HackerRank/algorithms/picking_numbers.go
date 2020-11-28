package main

import (
	"fmt"
	"math"
)
func pickingNumbers(a []int32) int32 {
	var count = [1000]int32{0}
	for _, element := range a {
		count[element] += 1
	}
	var max = int32(math.MinInt32)
	for i := 0; i < len(count)-1; i++ {
		if count[i] + count[i+1] > max {
			max = count[i] + count[i+1]
		}
	}
	return max
}

func main() {
	fmt.Println(pickingNumbers([]int32{4, 6, 5, 3, 3, 1}))
}