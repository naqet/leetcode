package main

import (
	"reflect"
	"testing"
)

type ListNode struct {
	Val  int
	Next *ListNode
}

func logic(head *ListNode, n int) *ListNode {
	dummy := &ListNode{-1, head}
	left, right := dummy, head

	for n > 0 && right != nil {
		right = right.Next
		n--
	}

	for right != nil {
		left = left.Next
		right = right.Next
	}

    if left.Next != nil {
        left.Next = left.Next.Next
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
	input    *ListNode
	n        int
	expected []int
}

var tests = []singleTest{
	{&ListNode{1, &ListNode{2, &ListNode{3, &ListNode{4, nil}}}}, 2, []int{1, 2, 4}},
	{&ListNode{1, nil}, 1, []int{}},
	{&ListNode{1, &ListNode{2, nil}}, 1, []int{1}},
}

func TestLogic(t *testing.T) {

	for _, test := range tests {
		result := walk(logic(test.input, test.n))

		if !reflect.DeepEqual(result, test.expected) {
			t.Error(result)
			break
		}
	}
}
