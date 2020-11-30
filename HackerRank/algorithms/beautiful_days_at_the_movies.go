package main

import (
	"fmt"
	"math"
	"strconv"
)


func ReverseNumber(number int) int {
	strNumber := strconv.Itoa(number)
	reverseStrNumber := ""
	for length := len(strNumber); length > 0; length-- {
		reverseStrNumber += string(strNumber[length-1])
	}
	reverseNum, error := strconv.Atoi(reverseStrNumber)
	if error != nil {
		fmt.Println("Failure to cast String to int")
	}
	return reverseNum
}

func beautifulDays(i int32, j int32, k int32) int32 {
	var count = int32(0)
	for d := i; d <= j; d++ {
		var result = math.Abs(float64(d-int32(ReverseNumber(int(d)))))/float64(k)
		if result == float64(int32(result)) {
			count += 1
		}
	}
	return count
}


func main() {
	fmt.Println(beautifulDays(20, 23, 6))
}
