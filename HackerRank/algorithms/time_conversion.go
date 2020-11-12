package main

import (
	"fmt"
	"time"
)

func timeConversion(s string) string {
	formattedTime, err := time.Parse("03:04:05PM", s)
	if err != nil {
		fmt.Println(err)
	}
	return formattedTime.Format("15:04:05")
}

func main() {
	fmt.Println(timeConversion("07:05:45PM"))
}