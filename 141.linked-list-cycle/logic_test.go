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

	if head == nil {
		return result
	}

	slow, fast := head, head.Next

	for fast != nil && fast.Next != nil {
		if slow == fast {
			result = true
			break
		}

		slow = slow.Next
		fast = fast.Next.Next
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
