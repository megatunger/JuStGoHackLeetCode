package main

import (
	"fmt"
	"math"
)

func catAndMouse(x int32, y int32, z int32) string {
	var distanceCatA = int32(math.Abs(float64(z) - float64(x)))
	var distanceCatB = int32(math.Abs(float64(z) - float64(y)))

	if distanceCatA < distanceCatB {
		return "Cat A"
	} else if distanceCatB < distanceCatA {
		return "Cat B"
	} else {
		return "Mouse C"
	}
}

func main() {
	fmt.Println(catAndMouse(1, 3, 2))
}