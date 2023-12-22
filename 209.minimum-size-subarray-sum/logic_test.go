package main

import (
	"reflect"
	"testing"
)

func logic(target int, nums []int) int {
	minLength := 0
	count := 0

	l, r := 0, 0

	for r < len(nums) {
		right := nums[r]
		count += right
		length := r - l + 1

		if count >= target {
			if length < minLength || minLength == 0 {
				minLength = length
			}

			count -= nums[l]
			count -= right
			l++
		} else {
			r++
		}
	}

	return minLength
}

type singleTest struct {
	arg1     int
	arg2     []int
	expected int
}

var tests = []singleTest{
	{7, []int{2, 3, 1, 2, 4, 3}, 2},
	{4, []int{1,4,4}, 1},
	{11, []int{1,1,1,1,1,1,1}, 0},
}

func TestLogic(t *testing.T) {
	for _, test := range tests {
		if output := logic(test.arg1, test.arg2); !reflect.DeepEqual(output, test.expected) {
			t.Error(output, test.expected)
		}
	}
}
