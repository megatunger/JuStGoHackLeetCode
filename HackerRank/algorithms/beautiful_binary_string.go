package main

import (
	"fmt"
	"strings"
)

func beautifulBinaryString(b string) int32 {
	return int32(strings.Count(b, "010"))
}

func main() {
	fmt.Println(beautifulBinaryString("0100101010"))
}