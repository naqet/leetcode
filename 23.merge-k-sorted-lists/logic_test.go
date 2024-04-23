package main

import (
	"reflect"
	"testing"
)

type ListNode struct {
	Val  int
	Next *ListNode
}

func logic(lists []*ListNode) *ListNode {
	var result *ListNode

    for _, list := range lists {
        result = mergeLists(result, list);
    }
	return result
}

func mergeLists(one *ListNode, two *ListNode) *ListNode {
    if one == nil && two == nil {
        return nil
    }
    
    if one == nil {
        return two
    }

    if two == nil {
        return one
    }

    dummy := &ListNode{}
    curr := dummy

    for one != nil || two != nil {
        if one == nil {
            curr.Next = &ListNode{Val: two.Val}
            two = two.Next
        } else if two == nil {
            curr.Next = &ListNode{Val: one.Val}
            one = one.Next
        } else if one.Val <= two.Val {
            curr.Next = &ListNode{Val: one.Val}
            one = one.Next
        } else if one.Val > two.Val {
            curr.Next = &ListNode{Val: two.Val}
            two = two.Next
        }

        curr = curr.Next
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

func TestLogic(t *testing.T) {
	one := &ListNode{1, &ListNode{4, &ListNode{5, nil}}}
	two := &ListNode{1, &ListNode{3, &ListNode{4, nil}}}
	three := &ListNode{2, &ListNode{6, nil}}

	result := walk(logic([]*ListNode{one, two, three}))

	if !reflect.DeepEqual(result, []int{1, 1, 2, 3, 4, 4, 5, 6}) {
		t.Error(result)
		return
	}
}
