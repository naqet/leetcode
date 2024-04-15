package main

import (
	"errors"
	"reflect"
	"testing"
)

var EmptyError = errors.New("Empty")

type Point struct {
	idx    int
	height int
}

type Stack struct {
	data []Point
}

func (s *Stack) push(idx, height int) {
	s.data = append(s.data, Point{idx, height})
}

func (s *Stack) pop() Point {
	result := s.data[len(s.data)-1]

	s.data = s.data[:len(s.data)-1]

	return result
}

func (s *Stack) top() (Point, error) {
	if len(s.data) == 0 {
		return Point{}, EmptyError
	}
	return s.data[len(s.data)-1], nil
}

func logic(heights []int) int {
	var result int
	stack := Stack{}

	for i, height := range heights {
		start := i
		point, err := stack.top()

		for err == nil && point.height > height {
			stack.pop()
			area := (i - point.idx) * point.height
			if area > result {
				result = area
			}

            start = point.idx;
            point, err = stack.top()
		}

		stack.push(start, height)
	}

	for _, point := range stack.data {
		area := (len(heights) - point.idx) * point.height
		if result < area {
			result = area
		}
	}

	return result
}

type singleTest struct {
	input    []int
	expected int
}

var tests = []singleTest{
	{[]int{2, 1, 5, 6, 2, 3}, 10},
	{[]int{2, 4}, 4},
	{[]int{1, 1}, 2},
	{[]int{2, 3}, 4},
	{[]int{2, 1, 2}, 3},
}

func TestLogic(t *testing.T) {
	for _, test := range tests {
		if output := logic(test.input); !reflect.DeepEqual(output, test.expected) {
			t.Error(output, test.expected)
		}
	}
}
