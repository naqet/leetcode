package main

import (
	"reflect"
	"testing"
)

func logic(input []int, target int) []int {
    hash := make(map[int]int, 0);

    for i, v := range input {
        hash[v] = i;
    }

    result := make([]int, 2);

    for idx, num := range input {
        if idx2, ok := hash[target - num]; ok && idx != idx2 {
            result[0] = idx
            result[1] = idx2
            break;
        }
    }
    return result;
}

type singleTest struct {
	arg1     []int
	arg2     int
	expected []int
}

var tests = []singleTest{
    {[]int{2,7,11,15}, 9, []int{0,1}},
    {[]int{3,2,4}, 6, []int{1,2}},
    {[]int{3,3}, 6, []int{0,1}},
}

func TestLogic(t *testing.T) {
	for _, test := range tests {
		if output := logic(test.arg1, test.arg2); !reflect.DeepEqual(output, test.expected) {
			t.Error(output, test.expected)
		}
	}
}
