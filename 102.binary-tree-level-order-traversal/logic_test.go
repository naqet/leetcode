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

func levelOrder(root *TreeNode) [][]int {
    if root == nil {
        return [][]int{}
    }

	queue := []*TreeNode{root}
	result := [][]int{}

    for len(queue) > 0 {
        level := len(queue)
        curr := []int{}

        for i := 0; i < level; i++ {
            head := queue[0]
            curr = append(curr, head.Val)

            if head.Left != nil {
                queue = append(queue, head.Left)
            }

            if head.Right != nil {
                queue = append(queue, head.Right)
            }
            queue = queue[1:]
        }
        result = append(result, curr)
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

	result := levelOrder(head)

	if !reflect.DeepEqual(result, [][]int{{1}, {2,3}, {4}}) {
		t.Error(result)
		return
	}

	head = &TreeNode{1, nil, &TreeNode{3, nil, &TreeNode{4, &TreeNode{5, nil, nil}, nil}}}

	result = levelOrder(head)

	if !reflect.DeepEqual(result, [][]int{{1}, {3}, {4}, {5}}) {
		t.Error(result)
		return
	}
}

