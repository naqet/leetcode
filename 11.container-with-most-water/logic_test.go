package main

import (
	"reflect"
	"testing"
)

func logic(height []int) int {
	var result int

    left, right := 0, len(height) - 1;

    for left < right {
        var area int;
        width := right - left;
        l := height[left]
        r := height[right]

        if l < r {
            area = l * width;
            left++;
        } else {
            area = r * width;
            right--
        }

        if area > result {
            result = area;
        }
    }

	return result
}

type singleTest struct {
	input    []int
	expected int
}

var tests = []singleTest{
	{[]int{1, 8, 6, 2, 5, 4, 8, 3, 7}, 49},
	{[]int{1, 1}, 1},
}

func TestLogic(t *testing.T) {
	for _, test := range tests {
		if output := logic(test.input); !reflect.DeepEqual(output, test.expected) {
			t.Error(output, test.expected)
		}
	}
}
