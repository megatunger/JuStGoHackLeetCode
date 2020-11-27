package main

import(
	"fmt"
)

// Complete the compareTriplets function below.
func compareTriplets(a []int32, b []int32) []int32 {
	var alice = int32(0)
	var bob = int32(0)
	for i, element := range a {
		if element > b[i] {
			alice += 1
		} else if element < b[i] {
			bob += 1
		}
	}
	return []int32{alice, bob}
}

func main() {
	fmt.Println(compareTriplets([]int32{17, 28, 30}, []int32{99, 16, 8}))
}