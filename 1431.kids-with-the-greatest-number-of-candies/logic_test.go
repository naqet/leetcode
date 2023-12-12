package main

import (
	"reflect"
	"testing"
)

func logic(candies []int, extraCandies int) []bool {
    maxCandie := maxInArray(candies);

    result := make([]bool, len(candies));

    for i, kid := range candies {
        if kid + extraCandies >= maxCandie {
            result[i] = true;
        }
    }

    return result
}

func maxInArray(arr []int) int {
    var maxInt int;

    for _, num := range arr {
        if maxInt < num {
            maxInt = num;
        }
    }
    return maxInt;
}
type singleTest struct {
	arg1     []int
	arg2     int
	expected []bool
}

var tests = []singleTest{
	{[]int{2, 3, 5, 1, 3}, 3, []bool{true, true, true, false, true}},
	{[]int{4, 2, 1, 1, 2}, 1, []bool{true, false, false, false, false}},
}

func TestLogic(t *testing.T) {
	for _, test := range tests {
		if output := logic(test.arg1, test.arg2); !reflect.DeepEqual(output, test.expected) {
			t.Error(output, test.expected)
		}
	}
}

