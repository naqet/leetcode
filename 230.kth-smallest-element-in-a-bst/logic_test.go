package main

import (
	"testing"
)

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func kthSmallest(root *TreeNode, k int) int {
	stack := []*TreeNode{}

    curr := root

	for len(stack) > 0 || curr != nil {
        for curr != nil {
            stack = append(stack, curr)
            curr = curr.Left
        }

        curr = stack[len(stack)-1]
        stack = stack[:len(stack)-1]

        k--
        if k == 0 {
            break
        }


        curr = curr.Right
	}

	return curr.Val
}

type singleTest struct {
	input    string
	expected bool
}

var tests = []singleTest{}

func TestLogic(t *testing.T) {
	head := &TreeNode{4, &TreeNode{1, nil, nil}, &TreeNode{5, nil, &TreeNode{9, nil, nil}}}

	result := kthSmallest(head, 1)

	if result != 1 {
		t.Error(result)
		return
	}

	//head = &TreeNode{4, nil, &TreeNode{7, &TreeNode{4, nil, nil}, &TreeNode{10, nil, nil}}}

	//result = kthSmallest(head, 1)

	//if result != 4 {
	//	t.Error(result)
	//	return
	//}
}
