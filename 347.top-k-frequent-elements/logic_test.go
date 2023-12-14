package main

import (
	"fmt"
	"reflect"
	"testing"
)

func logic(nums []int, amount int) []int {
    hash := make(map[int]int, 0);

    for _, v := range nums {
        hash[v]++;
    }

    countArr := make([][]int, len(nums) + 1);

    for num, count := range hash {
        countArr[count] = append(countArr[count], num);
    }

    result := []int{};

    outer: for i := len(countArr) - 1; i >= 0; i-- {
        values := countArr[i];

        if len(values) > 0 {
            for _, v := range values {
                if len(result) < amount {
                    result = append(result, v);
                } else {
                    break outer;
                }
            }
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
    {[]int{1,1,1,2,2,3}, 2, []int{1,2}},
    {[]int{1}, 1, []int{1}},
    {[]int{-1, -1}, 1, []int{-1}},
}

func TestLogic(t *testing.T) {
	for _, test := range tests {
		if output := logic(test.arg1, test.arg2); !reflect.DeepEqual(output, test.expected) {
			t.Error(output, test.expected)
		}
	}
}

