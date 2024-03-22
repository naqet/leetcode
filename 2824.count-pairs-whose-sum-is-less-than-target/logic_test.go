package main

import (
	"reflect"
	"sort"
	"testing"
)

func logic(nums []int, target int) int {
	if len(nums) < 2 {
		return 0
	}

	sort.Ints(nums)

	var result int

	i, j := 0, len(nums)-1

	for i < j && j < len(nums) {
		val := nums[i] + nums[j]

		if val < target {
            result += j - i;
			i++
		} else {
			j--
		}
	}
	return result
}

type singleTest struct {
	nums     []int
	target   int
	expected int
}

var tests = []singleTest{
	{[]int{-1, 1, 2, 3, 1}, 2, 3},
	{[]int{-6, 2, 5, -2, -7, -1, 3}, -2, 10},
}

func TestLogic(t *testing.T) {
	for _, test := range tests {
		if output := logic(test.nums, test.target); !reflect.DeepEqual(output, test.expected) {
			t.Error(output, test.expected)
		}
	}
}
