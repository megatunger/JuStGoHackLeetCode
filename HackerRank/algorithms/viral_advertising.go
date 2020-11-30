package main

import (
	"fmt"
	"math"
)

func viralAdvertising(n int32) int32 {
	var shared = int32(5)
	var liked = int32(2)
	var cumulative = int32(2)
	for i:=2; i <= int(n); i++ {
		shared = int32(math.Floor(float64(shared)/2))*3
		liked = shared / 2
		cumulative = cumulative + liked
	}
	return cumulative
}

func main() {
	fmt.Println(viralAdvertising(3))
}