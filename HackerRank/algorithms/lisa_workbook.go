package main

import (
	"fmt"
)

func workbook(n int32, k int32, arr []int32) int32 {
	var page = 1
	var result = 0
	for _, num := range arr {
		for counter := 1; counter <= int(num); counter++ {
			if counter == page {
				result += 1
			}
			if int32(counter) == num || int32(counter)%k == 0 {
				page += 1
			}
		}
	}
	return int32(result)
}

func main() {
	//fmt.Println(workbook(5, 3, []int32{4, 2, 6, 1, 10}))
	fmt.Println(workbook(10, 5, []int32{3, 8, 15, 11, 14, 1, 9, 2, 24, 31}))
}