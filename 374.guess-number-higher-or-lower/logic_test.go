package main

import (
	"reflect"
	"testing"
)

// This function has been provided by leetcode
func guess(input int) int {
    if input < 1 {
        return 1
    } else if input > 1 {
        return -1
    }

    return 0;
};

func logic(input int) int {
    var result int;

    lo := 0
    hi := input;

    for lo <= hi {
        mid := lo + (hi - lo) / 2

        val := guess(mid);

        if val == -1 {
            hi = mid;
        } else if val == 1 {
            lo = mid + 1;
        } else if val == 0 {
            result = mid;
            break;
        }
    }

    return result;
}

type singleTest struct {
	arg1     int
	expected int
}

var tests = []singleTest{
    {1, 1},
}

func TestLogic(t *testing.T) {
	for _, test := range tests {
		if output := logic(test.arg1); !reflect.DeepEqual(output, test.expected) {
			t.Error(output, test.expected)
		}
	}
}

