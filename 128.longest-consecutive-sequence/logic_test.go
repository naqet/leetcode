package main

import (
	"reflect"
	"testing"
)

func logic(nums []int) int {
	var result int
	set := map[int]struct{}{}

	for _, num := range nums {
		set[num] = struct{}{}
	}

	for num := range set {
		_, prev := set[num-1]

        if prev {
            continue;
        }

        num := num;
        consecutive := 1;
		for {
			_, ok := set[num+1]
			if !ok {
				break
			}
            consecutive++;
            num += 1;
		}

        if consecutive > result {
            result = consecutive;
        }

	}

	return result
}

type singleTest struct {
	input    []int
	expected int
}

var tests = []singleTest{
	{[]int{100, 4, 200, 1, 3, 2}, 4},
	{[]int{}, 0},
	{[]int{1}, 1},
	{[]int{0, 3, 7, 2, 5, 8, 4, 6, 0, 1}, 9},
}

func TestLogic(t *testing.T) {
	for _, test := range tests {
		if output := logic(test.input); !reflect.DeepEqual(output, test.expected) {
			t.Error(output, test.expected)
		}
	}
}
