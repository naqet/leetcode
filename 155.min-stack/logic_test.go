package main

import (
	"testing"
)

type MinStack struct {
    stack []int
    min[]int
}


func Constructor() MinStack {
    return MinStack{}
}


func (this *MinStack) Push(val int)  {
    if len(this.min) == 0 || val <= this.GetMin() {
        this.min = append(this.min, val);
    }
    this.stack = append(this.stack, val)
}


func (this *MinStack) Pop()  {
    if this.Top() == this.GetMin() {
        this.min = this.min[:len(this.min)-1]
    }
    this.stack = this.stack[:len(this.stack)-1]
}


func (this *MinStack) Top() int {
    return this.stack[len(this.stack)-1];
}


func (this *MinStack) GetMin() int {
    return this.min[len(this.min)-1];
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

