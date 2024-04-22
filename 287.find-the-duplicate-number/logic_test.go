package main

import (
	"reflect"
	"testing"
)

func logic(nums []int) int {
	var result int

    slow, fast := 0, 0;

    for {
         slow = nums[slow]
         fast = nums[nums[fast]]

         if slow == fast {
             result = nums[slow]
             break;
         }
    }


    second := 0;

    for {
        if slow == second {
            result = slow
            break
        }

        slow = nums[slow]
        second = nums[second]
    }

	return result
}

type singleTest struct {
	input    []int
	expected int
}

var tests = []singleTest{
	{[]int{1, 3, 4, 2, 2}, 2},
	{[]int{3, 1, 3, 4, 2}, 3},
	{[]int{3, 3, 3, 3, 3}, 2},
}

func TestLogic(t *testing.T) {
	for _, test := range tests {
		if output := logic(test.input); !reflect.DeepEqual(output, test.expected) {
			t.Error(output, test.expected)
		}
	}
}
