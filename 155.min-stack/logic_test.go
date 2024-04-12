package main

import (
	"testing"
)

type Node struct {
    value int
    min int
    prev *Node
}

type MinStack struct {
    head *Node
}


func Constructor() MinStack {
    return MinStack{}
}


func (this *MinStack) Push(val int) {
    newNode := &Node{value: val, min: val}
    if this.head == nil {
        this.head = newNode;
        return
    }

    if this.head.min < newNode.min {
        newNode.min = this.head.min;
    }

    newNode.prev = this.head;
    this.head = newNode;
}


func (this *MinStack) Pop() {
    if this.head == nil {
        return
    }

    if this.head.prev != nil {
        this.head = this.head.prev
    } else {
        this.head = nil;
    }
}


func (this *MinStack) Top() int {
    return this.head.value;
}


func (this *MinStack) GetMin() int {
    return this.head.min;
}

func TestLogic(t *testing.T) {
    s := MinStack{}

    s.Push(-2);

    if s.Top() != -2 {
        t.Error("Error on 1")
        return
    }

    s.Push(0);

    if s.Top() != 0 {
        t.Error("Error on 2")
        return
    }

    s.Push(-3);

    if s.Top() != -3 {
        t.Error("Error on 3")
        return
    }

    min := s.GetMin()

    if min != -3 {
        t.Error("Error on 4")
        return
    }

    s.Pop()

    if s.Top() != 0 {
        t.Error("Error on 5")
        return
    }

    min = s.GetMin()

    if min != -2 {
        t.Error("Error on 6")
        return
    }
}

