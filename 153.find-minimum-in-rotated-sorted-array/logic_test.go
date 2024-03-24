package main

import (
	"reflect"
	"testing"
)

func logic(nums []int) int {
    lo := 0;
    hi := len(nums) - 1;


    for lo < hi {
        mid := lo + (hi - lo) - 1
        val := nums[mid];
        right := nums[hi];

        if val > right {
            lo = mid + 1;
        } else {
            hi = mid;
        }
    }

    return nums[lo];
}

type singleTest struct {
	input     []int
	expected int
}

var tests = []singleTest{
    {[]int{3,4,5,1,2}, 1},
    {[]int{4,5,6,7,0,1,2}, 0},
    {[]int{11,13,15,17}, 11},
}

func TestLogic(t *testing.T) {
	for _, test := range tests {
		if output := logic(test.input); !reflect.DeepEqual(output, test.expected) {
			t.Error(output, test.expected)
		}
	}
}

