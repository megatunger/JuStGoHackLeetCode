package main

import (
	"fmt"
	"sort"
)

func miniMaxSum(arr []int32) {
	var sum = int64(0)
	sort.Slice(arr, func(i, j int) bool { return arr[i] < arr[j]})
	for _, num := range arr {
		sum += int64(num)
	}
	fmt.Print(sum-int64(arr[len(arr)-1]))
	fmt.Print(" ")
	fmt.Println(sum-int64(arr[0]))
}

func main() {
	miniMaxSum([]int32{1293912931,1230123213,5244,340034,34343})
}
