package main

import (
	"reflect"
	"testing"
)

func logic(nums []int, k int) []int {
    result := make([]int, 0);
    hash := make(map[int]int, len(nums));

    for _, num := range nums {
        hash[num]++;
    }

    arr := make([][]int, len(nums) + 1);

    for k, v := range hash {
        arr[v] = append(arr[v], k);
    }

    outer: for i := len(arr)-1; i >= 0; i-- {
        if len(arr[i]) == 0 {
            continue;
        }

        for _, num := range arr[i] {
            if len(result) == k {
                break outer;
            }
            result = append(result, num);
        }
    }

	return result
}

type singleTest struct {
	nums     []int
	k        int
	expected []int
}

var tests = []singleTest{
	{[]int{1, 1, 1, 2, 2, 3}, 2, []int{1, 2}},
	{[]int{1}, 1, []int{1}},
}

func TestLogic(t *testing.T) {
	for _, test := range tests {
		if output := logic(test.nums, test.k); !reflect.DeepEqual(output, test.expected) {
			t.Error(output, test.expected)
		}
	}
}
