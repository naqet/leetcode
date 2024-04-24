package main

import (
	"reflect"
	"testing"
)

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func invertTree(root *TreeNode) *TreeNode {
    if root == nil {
        return nil
    }

    tmp := root.Left
    root.Left = root.Right
    root.Right = tmp

    invertTree(root.Left)
    invertTree(root.Right)
	return root
}

func walk(node *TreeNode, path *[]int) {
    if node == nil {
        return
    }

    walk(node.Left, path)

    *path = append(*path, node.Val)

    walk(node.Right, path)
}

func TestLogic(t *testing.T) {

	root := &TreeNode{4, &TreeNode{2, &TreeNode{1, nil, nil}, &TreeNode{3, nil, nil}}, &TreeNode{7, &TreeNode{6, nil, nil}, &TreeNode{9, nil, nil}}}

    path := []int{}
    walk(invertTree(root), &path)

    if !reflect.DeepEqual(path, []int{9, 7, 6, 4, 3, 2, 1}) {
        t.Error(path)
        return
    }
}
