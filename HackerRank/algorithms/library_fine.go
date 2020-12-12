package main

import "fmt"

func libraryFine(d1 int32, m1 int32, y1 int32, d2 int32, m2 int32, y2 int32) int32 {
	if y1-y2>0 {
		return 10000
	} else if m1-m2>0 && y1==y2 {
		return 500 * (m1-m2)
	} else if d1-d2>0 && m1==m2 && y1==y2 {
		return 15 * (d1-d2)
	} else {
		return 0
	}
}

func main() {
	fmt.Println(libraryFine(9, 6, 2015, 6, 6, 2015))
}