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

func rightSideView(root *TreeNode) []int {
    if root == nil {
        return []int{}
    }

	queue := []*TreeNode{}
	result := []int{}

	queue = append(queue, root)

	for len(queue) > 0 {
		level := len(queue)

		for i := 0; i < level; i++ {
            popped := queue[0]
			if popped.Left != nil {
				queue = append(queue, popped.Left)
			}

			if popped.Right != nil {
				queue = append(queue, popped.Right)
			}

            if level - i == 1 {
                result = append(result, popped.Val)
            }
            queue = queue[1:]
		}
	}

	return result
}

type singleTest struct {
	input    string
	expected bool
}

var tests = []singleTest{}

func TestLogic(t *testing.T) {
	head := &TreeNode{1, &TreeNode{2, nil, nil}, &TreeNode{3, nil, &TreeNode{4, nil, nil}}}

	result := rightSideView(head)

	if !reflect.DeepEqual(result, []int{1, 3, 4}) {
		t.Error(result, []int{1,3,4})
		return
	}

	head = &TreeNode{1, &TreeNode{2, nil, nil}, &TreeNode{3, nil, &TreeNode{4, &TreeNode{5, nil, nil}, nil}}}

	result = rightSideView(head)

	if !reflect.DeepEqual(result, []int{1, 3, 4, 5}) {
		t.Error(result)
		return
	}
}
