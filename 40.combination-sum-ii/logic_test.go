package main

import (
	"reflect"
	"sort"
	"testing"
)

func combinationSum2(nums []int, target int) [][]int {
    result := [][]int{}
    sort.Ints(nums)
    dfs(0, 0, target, &nums, &[]int{}, &result)

    return result
}

func dfs(i, sum, target int, nums, sub *[]int, result *[][]int) {
    if sum == target {
        s := make([]int, len(*sub))
        copy(s, *sub)
        *result = append(*result, s)
        return
    }

    if i == len(*nums) || sum > target {
        return
    }

    *sub = append(*sub, (*nums)[i])
    dfs(i+1, sum + (*nums)[i], target, nums, sub, result)

    for i < len(*nums)-1 && (*nums)[i] == (*nums)[i+1] {
        i++
    }

    *sub = (*sub)[:len(*sub)-1]
    dfs(i+1, sum, target, nums, sub, result)
}

type singleTest struct {
	input    []int
    target int
	expected [][]int
}

var tests = []singleTest{
	{[]int{10,1,2,7,6,1,5}, 8, [][]int{{1,1,6}, {1,2,5}, {1,7}, {2,6}}},
	{[]int{5}, 5, [][]int{{5}}},
	{[]int{2,5,2,1,2}, 5, [][]int{{1,2,2}, {5}}},
}

func TestLogic(t *testing.T) {
	for _, test := range tests {
		if output := combinationSum2(test.input, test.target); !reflect.DeepEqual(output, test.expected) {
			t.Error(output, test.expected)
		}
	}
}

