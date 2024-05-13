package main

import (
	"reflect"
	"testing"
)

func permute(nums []int) [][]int {
    res := [][]int{}
    dfs(&res, 0, nums)
    return res
}

func dfs(result *[][]int, idx int, nums []int) {
    if idx == len(nums) {
        c := make([]int, len(nums))
        copy(c, nums)
        *result = append(*result, c)
        return 
    }

    for i := idx; i < len(nums);i++ {
        nums[idx], nums[i] = nums[i], nums[idx]
        dfs(result, idx + 1, nums)
        nums[i], nums[idx] = nums[idx], nums[i]
    }
}

type singleTest struct {
	input    []int
	expected [][]int
}

var tests = []singleTest{
	{[]int{1, 2, 3}, [][]int{{1, 2, 3}, {1, 3, 2}, {2, 1, 3}, {2, 3, 1}, {3, 1, 2}, {3, 2, 1}}},
}

func TestLogic(t *testing.T) {
	for _, test := range tests {
		if output := permute(test.input); !reflect.DeepEqual(output, test.expected) {
			t.Error(output, test.expected)
		}
	}
}
