package main

import (
	"reflect"
	"testing"
)

func logic(input []int) {
	for i, j := 0, 0; i < len(input); i++ {
		num1 := input[i]
		num2 := input[j]

		if num1 != 0 {
			input[i], input[j] = num2, num1
            j++
		}
	}
}

type singleTest struct {
	arg1     []int
	expected []int
}

var tests = []singleTest{
	{[]int{0, 1, 0, 4, 12}, []int{1, 4, 12, 0, 0}},
	{[]int{0}, []int{0}},
	{[]int{1, 0, 1}, []int{1, 1, 0}},
}

func TestLogic(t *testing.T) {
	for _, test := range tests {
		logic(test.arg1)
		if !reflect.DeepEqual(test.arg1, test.expected) {
			t.Errorf("%d != %d", test.arg1, test.expected)
		}
	}
}
