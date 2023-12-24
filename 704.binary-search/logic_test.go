package main

import (
	"reflect"
	"testing"
)

func logic(nums []int, target int) int {
    idx := -1;

    lo := 0;
    hi := len(nums);

    for lo < hi {
        mid := lo + (hi - lo) / 2;
        val := nums[mid];

        if target < val {
            hi = mid;
        } else if target > val {
            lo = mid + 1;
        } else {
            idx = mid;
            break;
        }
    }

    return idx;
}

type singleTest struct {
	arg1     []int
	arg2     int
	expected int
}

var tests = []singleTest{}

func TestLogic(t *testing.T) {
	for _, test := range tests {
		if output := logic(test.arg1, test.arg2); !reflect.DeepEqual(output, test.expected) {
			t.Error(output, test.expected)
		}
	}
}

