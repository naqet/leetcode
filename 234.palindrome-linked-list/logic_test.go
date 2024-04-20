package main

import (
	"testing"
)

type ListNode struct {
	Val  int
	Next *ListNode
}

func logic(head *ListNode) bool {
	if head == nil {
		return false
	}

	if head.Next == nil {
		return true
	}

	// find middle
	slow, fast := head, head.Next

	for fast != nil && fast.Next != nil {
		slow = slow.Next
		fast = fast.Next.Next
	}

	// reverse
	second := slow.Next
	slow.Next = nil
	var prev *ListNode

	for second != nil {
		tmp := second.Next
		second.Next = prev
		prev = second
		second = tmp
	}

	// compare
	result := true
	for head != nil && prev != nil {
		if head.Val != prev.Val {
			result = false
			break
		}
		head = head.Next
		prev = prev.Next
	}

	return result
}

type singleTest struct {
	input    *ListNode
	expected bool
}

var tests = []singleTest{
	{&ListNode{1, &ListNode{2, &ListNode{2, &ListNode{1, nil}}}}, true},
	{&ListNode{1, &ListNode{2, &ListNode{2, &ListNode{3, nil}}}}, false},
	{&ListNode{1, nil}, true},
	{nil, false},
}

func TestLogic(t *testing.T) {

	for _, test := range tests {
		result := logic(test.input)

		if result != test.expected {
			t.Error(result)
			break
		}
	}
}
