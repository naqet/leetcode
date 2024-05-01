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

func buildTree(preorder []int, inorder []int) *TreeNode {
    if len(preorder) == 0 || len(inorder) == 0 {
        return nil
    }
    root := &TreeNode{Val: preorder[0]}

    mid := 0;

    for i, val := range inorder {
        if val == root.Val {
            mid = i
            break;
        }
    }

    root.Left = buildTree(preorder[1:mid+1], inorder[:mid])
    root.Right = buildTree(preorder[mid+1:], inorder[mid+1:])
	return root
}

func inOrder(node *TreeNode, path *[]int) {
	if node == nil {
		return
	}

	inOrder(node.Left, path)

	*path = append(*path, node.Val)

	inOrder(node.Right, path)
}

func TestLogic(t *testing.T) {
	preorder := []int{3, 9, 20, 15, 7}
	inorder := []int{9, 3, 15, 20, 7}

	head := buildTree(preorder, inorder)

	path := []int{}
	inOrder(head, &path)

	if !reflect.DeepEqual(path, inorder) {
		t.Error()
		return
	}
}
