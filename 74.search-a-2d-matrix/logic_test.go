package main

import (
	"reflect"
	"testing"
)

func logic(matrix [][]int, target int) bool {
    var result bool;

    rows := len(matrix);
    cols := len(matrix[0])

    lo := 0;
    hi := rows * cols - 1;

    for lo <= hi {
        mid := lo + (hi - lo) / 2;
        val := matrix[mid/cols][mid%cols];

        if val == target {
            result = true
            break;
        } else if target < val {
            hi = mid - 1;
        } else {
            lo = mid + 1;
        }
    }

    return result
}

type singleTest struct {
	input    [][]int
	target   int
	expected bool
}

var tests = []singleTest{
	{[][]int{{1, 3, 5, 7}, {10, 11, 16, 20}, {23, 30, 34, 60}}, 3, true},
	{[][]int{{1, 3, 5, 7}, {10, 11, 16, 20}, {23, 30, 34, 60}}, 13, false},
    {[][]int{{1}}, 1, true},
}

func TestLogic(t *testing.T) {
	for _, test := range tests {
		if output := logic(test.input, test.target); !reflect.DeepEqual(output, test.expected) {
			t.Error(output, test.expected)
		}
	}
}
