package main

import (
	"reflect"
	"testing"
)

type ListNode struct {
    Val int
    Next *ListNode
}

func logic(head *ListNode) {
    // find middle
    slow, fast := head, head.Next

    for fast != nil && fast.Next != nil {
        slow = head.Next
        fast = fast.Next.Next
    }

    // reverse second half
    second := slow.Next
    slow.Next = nil
    var prev *ListNode


    for second != nil {
        tmp := second.Next
        second.Next = prev
        prev = second
        second = tmp;
    }


    // merge
    first := head;
    second = prev;

    for first != nil && second != nil {
        one, two := first.Next, second.Next
        first.Next = second;
        second.Next = one
        first = one
        second = two
    }
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

    logic(head);

    res := walk(head)

    if !reflect.DeepEqual(res, []int{1,5,2,4,3}) {
        t.Error(res)
        return
    }

    head = &ListNode{1, &ListNode{2, &ListNode{3, nil}}}

    logic(head);

    res = walk(head)

    if !reflect.DeepEqual(res, []int{1,3,2}) {
        t.Error(res)
        return
    }
}


