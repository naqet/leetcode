package main

import (
	"reflect"
	"testing"
)

func logic(input []int) int {
    left := 0;
    right := sum(input);

    result := -1;

    for i, v := range input {
        right -= v;

        if left == right {
            result = i;
            break;
        }

        left += v;
    }

    return result;
}

func sum(input []int) int {
    var sum int;
    for _, v := range input {
        sum += v;
    }

    return sum;
}

type singleTest struct {
	arg1     []int
	expected int
}

var tests = []singleTest{
    {[]int{1,7,3,6,5,6}, 3},
}

func TestLogic(t *testing.T) {
	for _, test := range tests {
		if output := logic(test.arg1); !reflect.DeepEqual(output, test.expected) {
			t.Error(output, test.expected)
		}
	}
}
