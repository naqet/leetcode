package main

import (
	"reflect"
	"sort"
	"testing"
)

func logic(nums []int, k int) int {
	sort.Ints(nums)

	l, r := 0, 0
	total := 0
	maxLength := 0

	for r < len(nums) {
		length := r - l + 1
		right := nums[r]
		left := nums[l]

		if right*length <= total+right+k {
            total += right
			r++
            if maxLength < length {
                maxLength = length
            }
		} else {
			total -= left
			l++
		}
	}

	return maxLength
}

type singleTest struct {
	arg1     []int
	arg2     int
	expected int
}

var tests = []singleTest{
	{[]int{1, 2, 4}, 5, 3},
	{[]int{1, 4, 8, 13}, 5, 2},
	{[]int{3,9,6}, 2, 1},
}

func TestLogic(t *testing.T) {
	for _, test := range tests {
		if output := logic(test.arg1, test.arg2); !reflect.DeepEqual(output, test.expected) {
			t.Error(output, test.expected)
		}
	}
}
