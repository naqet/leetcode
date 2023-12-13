package main

import (
	"reflect"
	"testing"
)

func logic(nums []int) []int {
	length := len(nums)
	newNums := make([]int, length)

	prefix := 1
	for i := 0; i < length; i++ {
		newNums[i] = prefix
		prefix *= nums[i]
	}

	right := 1
	for i := length - 2; i >= 0; i-- {
		right *= nums[i+1]
		newNums[i] *= right
	}

	return newNums
}

type singleTest struct {
	arg1     []int
	expected []int
}

var tests = []singleTest{
	{[]int{1, 2, 3, 4}, []int{24, 12, 8, 6}},
	{[]int{-1, -1, 0, -3, 3}, []int{0,0,9,0,0}},
}

func TestLogic(t *testing.T) {
	for _, test := range tests {
		if output := logic(test.arg1); !reflect.DeepEqual(output, test.expected) {
			t.Error(output, test.expected)
		}
	}
}
