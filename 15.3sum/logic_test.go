package main

import (
	"reflect"
	"slices"
	"testing"
)

func logic(nums []int) [][]int {
	result := [][]int{}
	slices.Sort(nums)

	for i, num := range nums {
		if i > 0 && num == nums[i-1] {
			continue
		}

		j, k := i+1, len(nums)-1
		for j < k {
			left := nums[j]
			right := nums[k]
			value := num + left + right

			if value == 0 {
				entry := []int{num, left, right}
				result = append(result, entry)
				j++
				k--

				for j < k && nums[j] == nums[j-1] {
					j++
				}

				for j < k && nums[k] == nums[k+1] {
					k--
				}
			} else if value < 0 {
				j++
			} else if value > 0 {
				k--
			}
		}
	}

	return result
}

type singleTest struct {
	input    []int
	expected [][]int
}

var tests = []singleTest{
	{[]int{-1, 0, 1, 2, -1, -4}, [][]int{{-1, -1, 2}, {-1, 0, 1}}},
	{[]int{-2, 0, 1, 1, 2}, [][]int{{-2, 0, 2}, {-2, 1, 1}}},
	{[]int{0, 1, 1}, [][]int{}},
	{[]int{0, 0, 0, 0}, [][]int{{0, 0, 0}}},
	{[]int{0, 0, 0}, [][]int{{0, 0, 0}}},
}

func TestLogic(t *testing.T) {
	for _, test := range tests {
		if output := logic(test.input); !reflect.DeepEqual(output, test.expected) {
			t.Error(output, test.expected)
		}
	}
}
