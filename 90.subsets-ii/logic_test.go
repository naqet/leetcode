package main

import (
	"reflect"
	"sort"
	"testing"
)

func logic(nums []int) [][]int {
	var result [][]int
	sort.Ints(nums)
	dfs(0, &nums, &[]int{}, &result)
	return result
}

func dfs(i int, nums, subs *[]int, result *[][]int) {
	if i >= len(*nums) {
		s := make([]int, len(*subs))
		copy(s, *subs)
		*result = append(*result, s)
		return
	}

	*subs = append(*subs, (*nums)[i])
	dfs(i+1, nums, subs, result)

	for i < len(*nums)-1 && (*nums)[i] == (*nums)[i+1] {
		i++
	}

	*subs = (*subs)[:len(*subs)-1]
	dfs(i+1, nums, subs, result)
}

type singleTest struct {
	input    []int
	expected [][]int
}

var tests = []singleTest{
	{[]int{1, 2, 2}, [][]int{{1, 2, 2}, {1, 2}, {1}, {2, 2}, {2}, {}}},
	{[]int{0}, [][]int{{0}, {}}},
}

func TestLogic(t *testing.T) {
	for _, test := range tests {
		if output := logic(test.input); !reflect.DeepEqual(output, test.expected) {
			t.Error(output, test.expected)
		}
	}
}
