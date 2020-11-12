package main

import "fmt"

func staircase(n int32) {
	for i:= int32(1); i <= n; i++ {
		for j:= int32(1); j <= n-i; j++ {
			fmt.Print(" ")
		}
		for j:= int32(n-i+1); j <= n; j++ {
			fmt.Print("#")
		}
		fmt.Print("\n")
	}
	fmt.Println()
}

func main() {
	staircase(6)
}
