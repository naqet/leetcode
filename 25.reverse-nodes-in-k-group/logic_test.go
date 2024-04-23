package main

import (
	"reflect"
	"testing"
)

type ListNode struct {
	Val  int
	Next *ListNode
}

func logic(head *ListNode, k int) *ListNode {
	if head == nil || k == 1 {
		return head
	}

	dummy := &ListNode{-1, head}

	group := dummy

	for {
		kth := getKth(group, k)

		if kth == nil {
			break
		}

		next := kth.Next

		prev := next
		curr := group.Next

		for curr != next {
			tmp := curr.Next
			curr.Next = prev
			prev = curr
			curr = tmp
		}

		tmp := group.Next
		group.Next = kth
		group = tmp
	}

	return dummy.Next
}

func getKth(curr *ListNode, k int) *ListNode {
	for curr != nil && k > 0 {
		curr = curr.Next
		k--
	}

	return curr
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

	result := walk(logic(head, 2))

	if !reflect.DeepEqual(result, []int{2, 1, 4, 3, 5}) {
		t.Error(result)
		return
	}
}
