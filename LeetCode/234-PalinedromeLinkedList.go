package main

import (
	"fmt"

)

type ListNode struct {
	Val int
	Next *ListNode
}

func isPalindrome(head *ListNode) bool {
	res := []int{}
	for head != nil {
		res = append(res, head.Val)
		head = head.Next
	}
	for i:=0; i< len(res)/2; i++ {
		if res[i] != res[len(res) - 1 - i] {
			return false
		}
	}
	return true
}

func main() {
	f1 := ListNode{1, &ListNode{Val: 2, Next: &ListNode{Val: 2, Next: &ListNode{Val: 2, Next: &ListNode{Val: 1, Next: nil}}}}}
	fmt.Println(isPalindrome(&f1))
}