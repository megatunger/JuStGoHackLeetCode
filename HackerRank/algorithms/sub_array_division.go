package main

import "fmt"

func birthday(s []int32, d int32, m int32) int32 {
	var sum [102]int32
	sum[0] = 0
	var count = 0
	if len(s) == 1 {
		if s[0] == d {
			return 1
		} else {
			return 0
		}
	}
	for i, number := range s {
		sum[i+1] = sum[i] + number
	}
	for i := 0; i <= len(s)-int(m); i++ {
		if sum[i+int(m)] - sum[i] == d {
			count += 1
		}
	}
	return int32(count)
}

func main() {
	fmt.Println(birthday([]int32{4}, 4, 1))
}