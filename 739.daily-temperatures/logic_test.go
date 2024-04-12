package main

import (
	"reflect"
	"testing"
)

type Node struct {
    value int
    idx int
    prev *Node
}

type Stack struct {
    length int
    head *Node
}

func (s *Stack) pop() (int, int) {
    if s.head == nil {
        return -1, -1;
    }

    s.length--;
    head := s.head;
    s.head = s.head.prev;
    return head.value, head.idx;
}

func (s *Stack) push(val int, idx int) {
    newNode := &Node{value: val, idx: idx}

    s.length++;
    if s.head == nil {
        s.head = newNode;
        return
    }

    newNode.prev = s.head;
    s.head = newNode;
}

func (s *Stack) top() (int, int) {
    if s.head == nil {
        return -1, -1
    }

    return s.head.value, s.head.idx;
}

func (s *Stack) view() []int {
    result := make([]int, s.length);
    head := s.head;

    for i:= 0; i < s.length && head != nil;i++ {
        result[i] = head.value;
        head = head.prev;
    }

    return result;
}

func dailyTemperatures(temperatures []int) []int {
    result := make([]int, len(temperatures));
    stack := Stack{}

    for idx, num := range temperatures {
        if stack.length == 0 {
            stack.push(num, idx);
            continue;
        }

        top, topIdx := stack.top();

        for top < num && stack.length != 0 {
            stack.pop()
            result[topIdx] = idx - topIdx;
            top, topIdx = stack.top()
        }

        stack.push(num, idx);
    }

	return result
}

type singleTest struct {
	input    []int
	expected []int
}

var tests = []singleTest{
	{[]int{73, 74, 75, 71, 69, 72, 76, 73}, []int{1, 1, 4, 2, 1, 1, 0, 0}},
	{[]int{30, 40, 50, 60}, []int{1, 1, 1, 0}},
	{[]int{30, 60, 90}, []int{1, 1, 0}},
}

func TestLogic(t *testing.T) {
	for _, test := range tests {
		if output := dailyTemperatures(test.input); !reflect.DeepEqual(output, test.expected) {
			t.Error(output, test.expected)
		}
	}
}
