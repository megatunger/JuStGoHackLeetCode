package main

import (
	"fmt"
)

func divisibleSumPairs(n int32, k int32, ar []int32) int32 {
	var remain [150]int
	var count = 0
	for _, element := range ar {
		remain[element%k]+=1
	}
	count += (remain[0]*(remain[0]-1))/2
	if (k%2==0) {
		count += (remain[k/2]*(remain[k/2]-1))/2
	}
	for i := 1; i <= int(k)/2 && i!=int(k)-i; i++ {
		count += remain[i]*remain[int(k)-i]
	}
	return int32(count)
}

func main() {
	fmt.Println(divisibleSumPairs(6, 3, []int32{1, 3, 2, 6, 1, 2}))
}
