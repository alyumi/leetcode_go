package main

import (
	"fmt"
	"math"
)

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func notNil(root *TreeNode) bool {
	return root != nil
}

func difference(root *TreeNode, node *TreeNode) int {
	if root.Val > node.Val {
		return root.Val - node.Val
	} else {
		return node.Val - root.Val
	}
}

func getMin(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func getMinimumDifference(root *TreeNode) int {
	minDifference := math.MaxInt

	if notNil(root.Left) {
		diff := difference(root, root.Left)

		minDifference = getMin(minDifference, getMinimumDifference(root.Left))
		minDifference = getMin(minDifference, diff)
	}

	if notNil(root.Right) {
		diff := difference(root, root.Right)

		minDifference = getMin(minDifference, getMinimumDifference(root.Right))
		minDifference = getMin(minDifference, diff)
	}

	return minDifference
}

func main() {

	var first_test *TreeNode = &TreeNode{
		Val: 4,
		Left: &TreeNode{
			Val: 2,
			Left: &TreeNode{
				Val:   1,
				Left:  nil,
				Right: nil,
			},
			Right: &TreeNode{
				Val:   3,
				Left:  nil,
				Right: nil,
			},
		},
		Right: &TreeNode{
			Val:   6,
			Left:  nil,
			Right: nil,
		},
	}

	fmt.Println(getMinimumDifference(first_test))
}
