package main

import (
	"fmt"
	"math"
)

func hurdleRace(k int32, height []int32) int32 {
	var max = int32(math.MinInt32)
	for _, element := range height {
		if element > max {
			max = element
		}
	}
	if max >= k {
		return max - k
	} else {
		return 0
	}
}

func main() {
	fmt.Println(hurdleRace(7, []int32{2, 5, 4, 5, 2}))
}
