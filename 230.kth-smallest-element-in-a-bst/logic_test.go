package main

import (
	"math"
	"testing"
)

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func kthSmallest(root *TreeNode, k int) int {
	return walk(root, &k)
}

func walk(node *TreeNode, k *int) int {
	res := math.MinInt
	if node == nil {
		return 0
	}

	res = max(walk(node.Left, k), res)

	*k--

	if *k == 0 {
		return node.Val
	}

	res = max(walk(node.Right, k), res)

	return res
}

type singleTest struct {
	input    string
	expected bool
}

var tests = []singleTest{}

func TestLogic(t *testing.T) {
	head := &TreeNode{4, &TreeNode{1, nil, nil}, &TreeNode{5, nil, &TreeNode{9, nil, nil}}}

	result := kthSmallest(head, 2)

	if result != 4 {
		t.Error(result)
		return
	}

	head = &TreeNode{4, nil, &TreeNode{7, &TreeNode{4, nil, nil}, &TreeNode{10, nil, nil}}}

	result = kthSmallest(head, 1)

	if result != 4 {
		t.Error(result)
		return
	}
}
