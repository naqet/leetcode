package main

import (
	"testing"
)

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func isSubtree(root *TreeNode, subRoot *TreeNode) bool {
    if subRoot == nil {
        return true
    }

    if root == nil {
        return false
    }

    if isSameTree(root, subRoot) {
        return true
    }
    
    return isSubtree(root.Left, subRoot) || isSubtree(root.Right, subRoot)
}

func isSameTree(one *TreeNode, two *TreeNode) bool {
	if one == nil && two == nil {
		return true
	}

	if one == nil || two == nil {
		return false
	}

	if one.Val != two.Val {
		return false
	}

	return isSameTree(one.Left, two.Left) && isSameTree(one.Right, two.Right)
}
 
type singleTest struct {
	input    string
	expected bool
}

var tests = []singleTest{}

func TestLogic(t *testing.T) {
	root := &TreeNode{3, &TreeNode{4, &TreeNode{1, nil, nil}, &TreeNode{2, nil, nil}}, &TreeNode{5, nil, nil}}
	subRoot := &TreeNode{4, &TreeNode{1, nil, nil}, &TreeNode{2, nil, nil}}

	result := isSubtree(root, subRoot)

    if !result {
		t.Error(result)
		return
	}
}
