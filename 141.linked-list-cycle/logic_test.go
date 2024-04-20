package main

import (
	"testing"
)

type ListNode struct {
	Val  int
	Next *ListNode
}

func logic(head *ListNode) bool {
	result := false
	hash := map[*ListNode]struct{}{}

	for head != nil {
		_, ok := hash[head]

		if ok {
			result = true
			break
		}

		hash[head] = struct{}{}
		head = head.Next
	}
	return result
}

func TestLogic(t *testing.T) {
	head := &ListNode{1, nil}
	head.Next = &ListNode{2, &ListNode{3, &ListNode{4, head}}}

	result := logic(head)

	if !result {
		t.Error(result)
		return
	}
}
