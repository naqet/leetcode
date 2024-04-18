package main

import (
	"reflect"
	"testing"
)

type ListNode struct {
	Val  int
	Next *ListNode
}

func logic(head *ListNode) *ListNode {
    var newHead *ListNode

	for head != nil {
		next := head.Next
		head.Next = newHead
		newHead = head
		head = next
	}

	return newHead
}

func walk(head *ListNode) []int {
	result := []int{}

	curr := head
	for curr != nil {
		result = append(result, curr.Val)
		curr = curr.Next
	}

	return result
}

func TestLogic(t *testing.T) {
	head := &ListNode{1, &ListNode{2, &ListNode{3, &ListNode{4, &ListNode{5, nil}}}}}

	target := logic(head)
	result := walk(target)
	if !reflect.DeepEqual(result, []int{5, 4, 3, 2, 1}) {
		t.Error("Nope", result)
		return
	}
}
