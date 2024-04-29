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

func isValidBST(root *TreeNode) bool {
	if root == nil {
		return false
	}

	return dfs(root, math.MinInt, math.MaxInt)
}

func dfs(root *TreeNode, minVal int, maxVal int) bool {
	if root == nil {
		return true
	}

	if root.Val >= maxVal || root.Val <= minVal {
		return false
	}

	return dfs(root.Left, minVal, root.Val) && dfs(root.Right, root.Val, maxVal)
}

type Test struct {
	head   *TreeNode
	result bool
}

var tests = []Test{
	{&TreeNode{5, &TreeNode{4, nil, nil}, &TreeNode{6, &TreeNode{3, nil, nil}, &TreeNode{10, nil, nil}}}, false},
	{&TreeNode{2, &TreeNode{2, nil, nil}, &TreeNode{2, nil, nil}}, false},
	{&TreeNode{5, &TreeNode{2, nil, nil}, &TreeNode{10, nil, nil}}, true},
}

func TestLogic(t *testing.T) {
	for _, test := range tests {
		result := isValidBST(test.head)

		if result != test.result {
			t.Error(result, test.result)
			return
		}

	}

}
