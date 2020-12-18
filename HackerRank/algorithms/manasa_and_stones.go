package main

import (
	"fmt"
	"sort"
)

func stones(n int32, a int32, b int32) []int32 {
	var arrStones []int32
	for i:=int32(0); i<n; i++ {
		var temp = a * ((n-1)-i) + b*i
		arrStones = append(arrStones, temp)
	}
	sort.Slice(arrStones, func(i, j int) bool { return arrStones[i] < arrStones[j] })
	return arrStones
}

func main() {
	fmt.Println(stones(3, 1, 2))
}


//func main() {
//	reader := bufio.NewReaderSize(os.Stdin, 1024 * 1024)
//
//	stdout, err := os.Create(os.Getenv("OUTPUT_PATH"))
//	checkError(err)
//
//	defer stdout.Close()
//
//	writer := bufio.NewWriterSize(stdout, 1024 * 1024)
//
//	TTemp, err := strconv.ParseInt(readLine(reader), 10, 64)
//	checkError(err)
//	T := int32(TTemp)
//
//	for TItr := 0; TItr < int(T); TItr++ {
//		nTemp, err := strconv.ParseInt(readLine(reader), 10, 64)
//		checkError(err)
//		n := int32(nTemp)
//
//		aTemp, err := strconv.ParseInt(readLine(reader), 10, 64)
//		checkError(err)
//		a := int32(aTemp)
//
//		bTemp, err := strconv.ParseInt(readLine(reader), 10, 64)
//		checkError(err)
//		b := int32(bTemp)
//
//		result := stones(n, a, b)
//		if result[0] == result[len(result) - 1] {
//			fmt.Fprintf(writer, "%v", result[0])
//		} else {
//			for i, resultItem := range result {
//				fmt.Fprintf(writer, "%d", resultItem)
//
//				if i != len(result) - 1 {
//					fmt.Fprintf(writer, " ")
//				}
//			}
//		}
//
//		fmt.Fprintf(writer, "\n")
//	}
//
//	writer.Flush()
//}
//
//func readLine(reader *bufio.Reader) string {
//	str, _, err := reader.ReadLine()
//	if err == io.EOF {
//		return ""
//	}
//
//	return strings.TrimRight(string(str), "\r\n")
//}
//
//func checkError(err error) {
//	if err != nil {
//		panic(err)
//	}
//}

