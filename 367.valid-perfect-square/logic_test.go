package main

import (
	"reflect"
	"testing"
)

func logic(num int) bool {
    lo, hi := 0, num + 1;

    square := false;

    for lo < hi {
        mid := lo + (hi - lo) / 2
        val := mid * mid;

        if val  == num {
            square = true
            break;
        } else if num < val {
            hi = mid
        } else if num > val {
            lo = mid + 1
        }
    }

    return square;
}


type singleTest struct {
	arg1     int
	expected bool
}

var tests = []singleTest{
    {16, true},
    {14, false},

}

func TestLogic(t *testing.T) {
	for _, test := range tests {
		if output := logic(test.arg1); !reflect.DeepEqual(output, test.expected) {
			t.Error(output, test.expected)
		}
	}
}

