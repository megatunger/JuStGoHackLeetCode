package main

import (
	"fmt"
)

// Complete the breakingRecords function below.
func breakingRecords(scores []int32) []int32 {
	var min = scores[0]
	var max = scores[0]
	var countMin = int32(0)
	var countMax = int32(0)
	for i := 1; i < len(scores); i++ {
		fmt.Println(scores[i])
		if scores[i] < min {
			min = scores[i]
			countMin += 1
		}
		if scores[i] > max {
			max = scores[i]
			countMax += 1
		}
	}
	return []int32{countMax, countMin}
}

func main() {
	fmt.Println(breakingRecords([]int32{3, 4, 21, 36, 10, 28, 35, 5, 24, 42}))
}