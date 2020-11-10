package main

import "fmt"

func twoSum(nums []int, target int) []int {
	numsInMap := make(map[int]int)
	for i := 0; i < len(nums); i++ {
		complement := target - nums[i]
		if val, ok := numsInMap[complement]; ok {
			return []int {val, i}
		}
		numsInMap[nums[i]] = i
	}

	return []int {-1, -1}
}

func main() {
	var numbers = []int {2,7,11,15}
	fmt.Println(twoSum(numbers, 9))
}
