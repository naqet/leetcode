package main

import (
	"reflect"
	"testing"
)

func logic(piles []int, h int) int {

    maxVal := 0;

    for _, val := range piles {
        if maxVal < val {
            maxVal = val;
        }
    }

    result := maxVal;
    lo := 1;

    if maxVal / h != 0 {
        lo = maxVal / h;
    }

    hi := maxVal;

    for lo <= hi {
        mid := lo + (hi - lo) / 2

        hours := 0;
        for _, pile := range piles {
            if pile % mid == 0 {
                hours += pile / mid
            } else {
                hours += (pile / mid) + 1
            }
        }

        if hours <= h {
            result = mid;
            hi = mid - 1;
        } else {
            lo = mid + 1;
        }
    }

	return result
}

type singleTest struct {
	input    []int
	h        int
	expected int
}

var tests = []singleTest{
	{[]int{3, 6, 7, 11}, 8, 4},
	{[]int{30, 11, 23, 4, 20}, 5, 30},
	{[]int{30, 11, 23, 4, 20}, 6, 23},
}

func TestLogic(t *testing.T) {
	for _, test := range tests {
		if output := logic(test.input, test.h); !reflect.DeepEqual(output, test.expected) {
			t.Error(output, test.expected)
		}
	}
}
