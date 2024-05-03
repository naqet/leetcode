package main

import (
	"testing"
)

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func maxPathSum(root *TreeNode) int {
    result := root.Val

    dfs(root, &result)

    return result
}

func dfs(node *TreeNode, res *int) int {
    if node == nil {
        return 0
    }

    maxLeft := max(dfs(node.Left, res), 0)
    maxRight := max(dfs(node.Right, res), 0)

    maxWithSplit := node.Val + maxLeft + maxRight;

    *res = max(*res, maxWithSplit)
    
    return node.Val + max(maxLeft, maxRight)
}

func TestLogic(t *testing.T) {
	root := &TreeNode{1, &TreeNode{2, nil, nil}, &TreeNode{3, nil, nil}}
    result := maxPathSum(root)

    if result != 6 {
        t.Error(result, 6)
        return
    }
	root = &TreeNode{-10, &TreeNode{9, nil, nil}, &TreeNode{20, &TreeNode{15, nil, nil}, &TreeNode{7, nil, nil}}}
    result = maxPathSum(root)

    if result != 42 {
        t.Error(result, 42)
        return
    }
}

