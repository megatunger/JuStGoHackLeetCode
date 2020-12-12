package main

import (
	"fmt"
	"sort"
)

func cutTheSticks(arr []int32) []int32 {
	result := make([]int32, 0)
	sort.Slice(arr, func(i, j int) bool { return arr[i] < arr[j] })
	fmt.Println(arr)
	result = append(result, int32(len(arr)))
	for i:=1; i < len(arr); i++ {
		if arr[i]!=arr[i-1] {
			result = append(result, int32(len(arr)-i))
		}
	}
	return result
}

func main() {
	fmt.Println(cutTheSticks([]int32{5,4,4,2,2,8}))
}