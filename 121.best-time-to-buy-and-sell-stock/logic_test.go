package main

import (
	"reflect"
	"testing"
)

func logic(prices []int) int {
	var result int
	if len(prices) <= 1 {
		return result
	}

	i, j := 0, 1
	for j < len(prices) && i < j {
		left := prices[i]
		right := prices[j]

		if left < right {
			val := right - left
			result = max(val, result)
		} else {
			i = j
		}
		j++
	}

	return result
}

type singleTest struct {
	input    []int
	expected int
}

var tests = []singleTest{
	{[]int{7, 1, 5, 3, 6, 4}, 5},
	{[]int{7, 6, 4, 3, 1}, 0},
	{[]int{3, 3}, 0},
}

func TestLogic(t *testing.T) {
	for _, test := range tests {
		if output := logic(test.input); !reflect.DeepEqual(output, test.expected) {
			t.Error(output, test.expected)
		}
	}
}
