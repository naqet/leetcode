package main

import (
	"testing"
)

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func logic(p *TreeNode, q *TreeNode) bool {
	if p == nil && q == nil {
		return true
	}

	if p == nil || q == nil {
		return false
	}

	if p.Val != q.Val {
		return false
	}

	return logic(p.Left, q.Left) && logic(p.Right, q.Right)
}

type singleTest struct {
	input    string
	expected bool
}

var tests = []singleTest{}

func TestLogic(t *testing.T) {
	first := &TreeNode{1, &TreeNode{2, nil, nil}, &TreeNode{3, nil, nil}}
	second := &TreeNode{1, &TreeNode{2, nil, nil}, &TreeNode{3, nil, nil}}

	result := logic(first, second)

	if !result {
		t.Error()
		return
	}

	second = &TreeNode{1, &TreeNode{2, &TreeNode{4, nil, nil}, nil}, &TreeNode{3, nil, nil}}

	result = logic(first, second)

	if result {
		t.Error()
		return
	}
}
