package main

import "fmt"

func Find(a []int32, x int32) int32 {
	for i, n := range a {
		if x == n {
			return int32(i+1)
		}
	}
	return -1
}

func permutationEquation(p []int32) []int32 {
	var result []int32
	for i := 1; i <= len(p); i++ {
		y := Find(p, int32(i))
		x := Find(p, y)
		result = append(result, x)
	}
	return result
}

func main() {
	//permutationEquation([]int32{4, 3, 5, 1, 2})
	fmt.Println(permutationEquation([]int32{2, 3, 1}))
}