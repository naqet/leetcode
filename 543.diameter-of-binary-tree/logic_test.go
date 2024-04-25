package main

import (
	"testing"
)

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

var maxD int

func logic(root *TreeNode) int {
	maxD = 0
	find(root)
	return maxD
}

func find(node *TreeNode) int {
	if node == nil {
		return 0
	}
	left := find(node.Left)
	right := find(node.Right)
	localMax := left + right
	maxD = max(maxD, localMax)
	return 1 + max(left, right)
}

type singleTest struct {
	input    string
	expected bool
}

var tests = []singleTest{}

func TestLogic(t *testing.T) {
	root := &TreeNode{4, &TreeNode{2, &TreeNode{1, nil, nil}, &TreeNode{3, nil, nil}}, &TreeNode{7, &TreeNode{6, nil, nil}, &TreeNode{9, nil, nil}}}

	result := logic(root)

	if result != 4 {
		t.Error(result, 4)
		return
	}
}
