package main

import "fmt"

func countApplesAndOranges(s int32, t int32, a int32, b int32, apples []int32, oranges []int32) {
	var countApples = 0
	var countOranges = 0
	for _, element := range apples {
		value := a + element
		if s <= value && value <= t {
			countApples += 1
		}
	}

	for _, element := range oranges {
		value := b + element
		if s <= value && value <= t {
			countOranges += 1
		}
	}
	fmt.Println(countApples)
	fmt.Println(countOranges)
}

func main() {
	countApplesAndOranges(7, 10, 4, 12, []int32{2, 3, -4}, []int32{3, -2, -4})
}