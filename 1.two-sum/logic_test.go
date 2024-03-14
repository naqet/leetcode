package main

import (
	"reflect"
	"testing"
)

func logic(nums []int, target int) []int {
	hash := make(map[int]int, len(nums))

	for i, v := range nums {
		hash[v] = i
	}

	result := make([]int, 2)

	for idx, num := range nums {
		if idx2, ok := hash[target-num]; ok && idx != idx2 {
			result[0] = idx
			result[1] = idx2
			break
		}
	}

    return result;
}

type singleTest struct {
	nums     []int
	target   int
	expected []int
}

var tests = []singleTest{
	//{[]int{1, 2, 3, 4}, 3, []int{0, 1}},
	{[]int{3, 3}, 6, []int{0, 1}},
}

func TestLogic(t *testing.T) {
	for _, test := range tests {
		if output := logic(test.nums, test.target); !reflect.DeepEqual(output, test.expected) {
			t.Error(output, test.expected)
		}
	}
}
