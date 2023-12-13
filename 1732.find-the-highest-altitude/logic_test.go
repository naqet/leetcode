package main

import (
	"reflect"
	"testing"
)

func logic(input []int) int {
	alt := 0
	highest := 0
	for _, v := range input {
		alt += v
		if alt > highest {
			highest = alt
		}
	}
	return highest
}

type singleTest struct {
	arg1     []int
	expected int
}

var tests = []singleTest{
	{[]int{-5, 1, 5, 0, -7}, 1},
	{[]int{-4, -3, -2, -1, 4, 3, 2}, 0},
}

func TestLogic(t *testing.T) {
	for _, test := range tests {
		if output := logic(test.arg1); !reflect.DeepEqual(output, test.expected) {
			t.Error(output, test.expected)
		}
	}
}
