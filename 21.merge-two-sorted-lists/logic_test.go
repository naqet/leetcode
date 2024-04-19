package main

import (
	"reflect"
	"testing"
)

type ListNode struct {
	Val  int
	Next *ListNode
}

func logic(list1 *ListNode, list2 *ListNode) *ListNode {
	first := &ListNode{}
	var tail *ListNode = first

	for list1 != nil && list2 != nil {
		if list1.Val <= list2.Val {
			tail.Next = list1
			list1 = list1.Next
		} else {
			tail.Next = list2
			list2 = list2.Next
		}

		tail = tail.Next
	}

    if list1 != nil {
        tail.Next = list1
    } else if list2 != nil {
        tail.Next = list2
    }

	return first.Next
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
	one := &ListNode{1, &ListNode{2, &ListNode{4, nil}}}
	two := &ListNode{1, &ListNode{3, &ListNode{4, nil}}}

	target := logic(one, two)
	result := walk(target)
	if !reflect.DeepEqual(result, []int{1, 1, 2, 3, 4, 4}) {
		t.Error("Nope", result)
		return
	}
}
