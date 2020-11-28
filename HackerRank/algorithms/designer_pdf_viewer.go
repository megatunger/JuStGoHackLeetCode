package main

import (
	"fmt"
	"math"
)

func designerPdfViewer(h []int32, word string) int32 {
	var max = int32(math.MinInt32)
	for _, element := range word {
		var pos = element-int32('a')
		if h[pos] > max {
			max = h[pos]
		}
	}
	return max * int32(len(word))
}

func main() {
	fmt.Println(designerPdfViewer([]int32 {1, 3, 1, 3, 1, 4, 1, 3, 2, 5, 5, 5, 5, 5, 5, 5, 5, 5, 5, 5, 5, 5, 5, 5, 5, 7}, "zaba"))
}

