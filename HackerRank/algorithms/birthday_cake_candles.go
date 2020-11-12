package main

import (
	"fmt"
)

func birthdayCakeCandles(candles []int32) int32 {
	var count[100000000] int32
	var max = int32(-999999999)
	for _, element := range candles {
		count[element] += 1
		if element > max {
			max = element
		}
	}
	return count[max]
}

func main() {
	fmt.Println(birthdayCakeCandles([]int32{3, 2, 1, 3}))

}