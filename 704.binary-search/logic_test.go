package main

import (
	"reflect"
	"testing"
)

func logic(input []int, target int) int {
    result := -1;

    lo := 0;
    hi := len(input);

    for lo < hi {
        mid := lo + (hi - lo) / 2;
        val := input[mid];

        if val == target {
            result = mid;
            break;
        } else if val < target {
            lo = mid + 1;
        } else {
            hi = mid;
        }
    }

	return result
}

type singleTest struct {
	input    []int
	target   int
	expected int
}

var tests = []singleTest{
    {[]int{-1,0,3,5,9,12}, 9, 4},
    {[]int{-1,0,3,5,9,12}, 2, -1},
}

func TestLogic(t *testing.T) {
	for _, test := range tests {
		if output := logic(test.input, test.target); !reflect.DeepEqual(output, test.expected) {
			t.Error(output, test.expected)
		}
	}
}
