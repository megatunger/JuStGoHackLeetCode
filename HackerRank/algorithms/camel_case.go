package main

import (
	"fmt"
)

func camelcase(s string) int32 {
	count := 1
	for pos, char := range s {
		if int32(char) >= 65 && int32(char) <= 90 && s[pos+1] > 90 {
			count++
		}
	}
	return int32(count)
}

func main() {
	fmt.Println(camelcase("saveChangesInTheEditor"))
}