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

// TODO: make it faster

func getMinimumDifference(root *TreeNode) int {
	minDiff := math.MaxInt64
	prev := -1

	var move func(node *TreeNode)
	move = func(node *TreeNode) {
		if node == nil {
			return
		}
		move(node.Left)
		if prev != -1 && node.Val-prev < minDiff {
			minDiff = node.Val - prev
		}
		prev = node.Val
		move(node.Right)
	}

	move(root)
	return minDiff
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

	var second_test *TreeNode = &TreeNode{
		Val:  1,
		Left: nil,
		Right: &TreeNode{
			Val:   2,
			Left:  nil,
			Right: nil,
		},
	}

	fmt.Println(getMinimumDifference(first_test))
	fmt.Println(getMinimumDifference(second_test))
}
