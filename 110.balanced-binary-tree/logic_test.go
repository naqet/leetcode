package main

import (
	"testing"
)

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func isBalanced(root *TreeNode) bool {
    if root == nil {
        return true
    }

    _, ok := dfs(root)

    return ok
}

func dfs(node *TreeNode) (int, bool) {
    if node == nil {
        return 0, true
    }

    left, okL := dfs(node.Left)
    right, okR := dfs(node.Right)

    if !okL || !okR {
        return 0, false
    }

    if max(left, right) - min(left, right) > 1 {
        return 0, false
    }

    return 1 + max(left, right), true
}

type singleTest struct {
	input    string
	expected bool
}

var tests = []singleTest{}

func TestLogic(t *testing.T) {
    root := &TreeNode{1, nil, &TreeNode{2, nil, &TreeNode{3, nil, nil}}}

    result := isBalanced(root)

	if result {
		t.Error(result)
		return
	}
}
