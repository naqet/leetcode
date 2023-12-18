package main

import (
	"reflect"
	"testing"
)

func logic(input []int) int {
	profit := 0

	l := 0
	for r := 0; r < len(input); r++ {
		left := input[l]
		right := input[r]

		if left > right {
			l = r
		} else if right-left > profit {
			profit = right - left
		}
	}

	return profit
}

type singleTest struct {
	arg1     []int
	expected int
}

var tests = []singleTest{
	{[]int{7, 1, 5, 3, 6, 4}, 5},
	{[]int{7, 6, 4, 3, 1}, 0},
}

func TestLogic(t *testing.T) {
	for _, test := range tests {
		if output := logic(test.arg1); !reflect.DeepEqual(output, test.expected) {
			t.Error(output, test.expected)
		}
	}
}
