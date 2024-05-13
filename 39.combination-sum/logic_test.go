package main

import (
	"reflect"
	"testing"
)

func logic(candidates []int, target int) [][]int {
	var result [][]int
	dfs(0, target, 0, &result, &candidates, &[]int{})
	return result
}

func dfs(i, target, total int, result *[][]int, nums, subset *[]int) {
	if total == target {
		sub := make([]int, len(*subset))
		copy(sub, *subset)
		*result = append(*result, sub)
		return
	}

	if i >= len(*nums) || total > target {
		return
	}

	*subset = append(*subset, (*nums)[i])

	dfs(i, target, total+(*nums)[i], result, nums, subset)

	*subset = (*subset)[:len(*subset)-1]

	dfs(i+1, target, total, result, nums, subset)
}

type singleTest struct {
	input    []int
	target   int
	expected [][]int
}

var tests = []singleTest{
	{[]int{2, 3, 6, 7}, 7, [][]int{{2, 2, 3}, {7}}},
	{[]int{2, 3, 5}, 8, [][]int{{2, 2, 2, 2}, {2, 3, 3}, {3, 5}}},
}

func TestLogic(t *testing.T) {
	for _, test := range tests {
		if output := logic(test.input, test.target); !reflect.DeepEqual(output, test.expected) {
			t.Error(output, test.expected)
		}
	}
}
