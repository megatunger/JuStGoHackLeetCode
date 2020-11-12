package main

import "fmt"

// Complete the aVeryBigSum function below.
func aVeryBigSum(ar []int64) int64 {
	var sum int64 = 0
	for _, element := range ar {
		sum += element
	}
	return sum;
}

func main() {
	var ar = []int64{1000000001, 1000000002, 1000000003, 1000000004, 1000000005}
	result := aVeryBigSum(ar)
	fmt.Println(result)
}

