package main

import "testing"

type Node struct {
	Val    int
	Next   *Node
	Random *Node
}

func logic(head *Node) *Node {
	if head == nil {
		return nil
	}

	curr := head
	hash := map[*Node]*Node{}

	for curr != nil {
		hash[curr] = &Node{Val: curr.Val}
		curr = curr.Next
	}

	curr = head

	for curr != nil {
		hash[curr].Next = hash[curr.Next]
		hash[curr].Random = hash[curr.Random]
		curr = curr.Next
	}

	return hash[head]
}

func TestLogic(t *testing.T) {
	head := &Node{1, &Node{2, &Node{3, nil, nil}, nil}, nil}

	result := logic(head)

	if result == head {
		t.Error(result, head)
		return
	}
}
