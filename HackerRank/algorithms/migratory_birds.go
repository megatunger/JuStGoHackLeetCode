package main

import (
	"fmt"
)

func migratoryBirds(arr []int32) int32 {
	var count = make(map[int32]int)
	for _, element := range arr {
		count[element] = count[element] + 1
	}
	var max_value = -999999999
	var max_key = -999999999
	for key, element := range count {
		if element > max_value {
			max_value = element
			max_key = int(key)
		}
	}
	return int32(max_key)
}

func main() {
	fmt.Println(migratoryBirds([]int32{1, 4, 4, 4, 5, 3}))
}