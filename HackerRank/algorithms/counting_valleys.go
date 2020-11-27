package main

import (
	"fmt"
)

func countingValleys(steps int32, path string) int32 {
	var level = 0
	var valley = int32(0)
	for i := int32(0); i < steps; i++ {
		if path[i] == 'U' {
			level += 1
		} else {
			level -= 1
		}
		if level == 0 && path[i] == 'U' {
			valley += 1
		}
	}
	return valley
}

func main() {
	fmt.Println(countingValleys(8, "DDUUDDUDUUUD"))
}