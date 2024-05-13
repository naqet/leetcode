package main

import (
	"reflect"
	"testing"
)

func logic(nums []int) [][]int {
    result := [][]int{}

    subset := []int{}
    dfs(0, &result, &nums, &subset)

	return result
}

func dfs(i int, res *[][]int, nums, subset *[]int) {
	if i >= len(*nums) {
        sub := make([]int, len(*subset))
        copy(sub, *subset)
		*res = append(*res, sub)
		return
	}

	*subset = append(*subset, (*nums)[i])
	dfs(i + 1, res, nums, subset)

	*subset = (*subset)[:len(*subset)-1]
	dfs(i + 1, res, nums, subset)
}

type singleTest struct {
	input    []int
	expected [][]int
}

var tests = []singleTest{
	{[]int{1, 2, 3}, [][]int{{}, {1}, {2}, {1, 2}, {3}, {1, 3}, {2, 3}, {1, 2, 3}}},
	{[]int{0}, [][]int{{}, {0}}},
}

func TestLogic(t *testing.T) {
	for _, test := range tests {
		if output := logic(test.input); !reflect.DeepEqual(output, test.expected) {
			t.Error(output, test.expected)
		}
	}
}
