package main

import (
	"reflect"
	"testing"
)

func logic(nums []int, target int) int {
	result := -1

	lo := 0
	hi := len(nums) - 1

	for lo <= hi {
		mid := lo + (hi-lo)/2
		val := nums[mid]

		left := nums[lo]
		right := nums[hi]

		if val == target {
			result = mid
			break
		} else if left <= val {
			if left <= target && target < val {
				hi = mid - 1
			} else {
				lo = mid + 1
			}
		} else {
			if target <= right && val < target {
				lo = mid + 1
			} else {
				hi = mid - 1
			}
		}
	}

	return result
}

type singleTest struct {
	input    []int
	target   int
	expected int
}

var tests = []singleTest{
	{[]int{4, 5, 6, 7, 0, 1, 2}, 0, 4},
	{[]int{4, 5, 6, 7, 0, 1, 2}, 3, -1},
	{[]int{1}, 0, -1},
	{[]int{1}, 1, 0},
	{[]int{1, 3, 5}, 1, 0},
	{[]int{1, 3, 5}, 5, 2},
	{[]int{3, 5, 1}, 3, 0},
}

func TestLogic(t *testing.T) {
	for _, test := range tests {
		if output := logic(test.input, test.target); !reflect.DeepEqual(output, test.expected) {
			t.Error(output, test.expected)
		}
	}
}
