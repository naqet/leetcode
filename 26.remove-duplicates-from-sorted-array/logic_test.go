package main

import (
	"reflect"
	"testing"
)

func logic(nums []int) int {
	if len(nums) <= 1 {
		return len(nums)
	}

	j := 0
	for i := 1; i < len(nums); i++ {
		if nums[j] != nums[i] {
            j++
            nums[j] = nums[i]
		}
	}

	return j + 1
}

type singleTest struct {
	input    []int
	expected int
}

var tests = []singleTest{
	{[]int{1, 1, 2}, 2},
	{[]int{0, 0, 1, 1, 1, 2, 2, 3, 3, 4}, 5},
}

func TestLogic(t *testing.T) {
	for _, test := range tests {
		if output := logic(test.input); !reflect.DeepEqual(output, test.expected) {
			t.Error(output, test.expected)
		}
	}
}
