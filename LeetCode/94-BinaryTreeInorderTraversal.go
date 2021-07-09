package main

import "fmt"

type TreeNode struct {
	Val int
	Left *TreeNode
	Right *TreeNode
}

var results []int

func traversalProcessing(root *TreeNode) {
	if root == nil {
		return
	} else {
		traversalProcessing(root.Left)
		results = append(results, root.Val)
		traversalProcessing(root.Right)
	}
}

func inorderTraversal(root *TreeNode) []int {
	results = []int{}
	traversalProcessing(root)
	return results
}

func main() {
	f1 := TreeNode{Val: 1, Left: nil, Right: &TreeNode{Val: 2, Left: &TreeNode{Val: 3, Right: nil}}};
	fmt.Println(inorderTraversal(&f1));
}