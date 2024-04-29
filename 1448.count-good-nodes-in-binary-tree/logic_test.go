package main

import (
	"testing"
)

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func goodNodes(root *TreeNode) int {
	return dfs(root, root.Val)
}

func dfs(node *TreeNode, maxVal int) int {
	if node == nil {
		return 0
	}

	if node.Val >= maxVal {
		maxVal = node.Val
        return 1 + dfs(node.Left, maxVal) + dfs(node.Right, maxVal)
	}

    return dfs(node.Left, maxVal) + dfs(node.Right, maxVal);

}

type singleTest struct {
	input    string
	expected bool
}

var tests = []singleTest{}

func TestLogic(t *testing.T) {
	head := &TreeNode{7, &TreeNode{3, nil, nil}, &TreeNode{5, nil, &TreeNode{11, nil, nil}}}

	result := goodNodes(head)

	if result != 2 {
		t.Error(result, 2)
		return
	}
}
