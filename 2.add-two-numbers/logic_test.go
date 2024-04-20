package main

import (
	"reflect"
	"testing"
)

type ListNode struct {
	Val  int
	Next *ListNode
}

func logic(l1 *ListNode, l2 *ListNode) *ListNode {
	dummy := &ListNode{}
	curr := dummy

	carry := 0
	for l1 != nil || l2 != nil {
		one, two := 0, 0

		if l1 != nil {
			one = l1.Val
		}

		if l2 != nil {
			two = l2.Val
		}

		sum := one + two + carry

		carry = sum / 10

		sum = sum % 10
		curr.Next = &ListNode{sum, nil}

		curr = curr.Next

		if l1 != nil {
			l1 = l1.Next
		}

		if l2 != nil {
			l2 = l2.Next
		}

	}

	if carry != 0 {
		curr.Next = &ListNode{carry, nil}
	}

	return dummy.Next
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

type singleTest struct {
	l1       *ListNode
	l2       *ListNode
	expected []int
}

var tests = []singleTest{
	{
		&ListNode{2, &ListNode{4, &ListNode{3, nil}}},
		&ListNode{5, &ListNode{6, &ListNode{4, nil}}},
		[]int{7, 0, 8},
	},
}

func TestLogic(t *testing.T) {

	for _, test := range tests {
		result := walk(logic(test.l1, test.l2))

		if !reflect.DeepEqual(result, test.expected) {
			t.Error(result)
			break
		}
	}
}
