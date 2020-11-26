package main

import (
	"fmt"
	"math"
)

func pageCount(n int32, p int32) int32 {
	var fromLeft = p / 2
	var fromRight = (n-p)/2
	return int32(math.Min(float64(fromLeft), float64(fromRight)))
}

func main() {
	fmt.Println(pageCount(6, 2))
}
