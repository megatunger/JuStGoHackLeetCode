package main

import (
	"fmt"
)

func enqueue(queue[] int32, element int32) []int32 {
	queue = append(queue, element)
	fmt.Println("Enqueued:", element)
	return queue
}

func dequeue(queue[] int32) ([]int32) {
	element := queue[len(queue) - 1]
	fmt.Println("Dequeued:", element)
	return queue[:len(queue)-1]
}

func circularArrayRotation(a []int32, k int32, queries []int32) []int32 {
	var queue[] int32
	for _, element := range a {
		queue = enqueue(queue, element)
	}
	for i:=0; i<int(k); i++ {
		_lastElement := queue[len(queue)-1]
		queue = dequeue(queue)
		queue = append([]int32{_lastElement}, queue...)
	}
	var result[] int32
	for _, element := range queries {
		result = enqueue(result, queue[element])
	}
	return result
}

func main() {
	fmt.Println(circularArrayRotation([]int32{3, 4, 5}, 2, []int32{1, 2}))
}