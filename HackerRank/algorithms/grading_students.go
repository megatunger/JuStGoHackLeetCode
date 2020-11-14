package main

import "fmt"

func gradingStudents(grades []int32) []int32 {
	result := make([]int32, len(grades))
	for i, num := range grades {
		if num < 38 {
			result[i] = num
		} else {
			compare := (num/5 +1)*5
			if compare-num < 3 {
				result[i] = compare
			} else {
				result[i] = num
			}
		}
	}
	return result
}

func main() {
	fmt.Println(gradingStudents([]int32{73, 67, 38, 33}))
}