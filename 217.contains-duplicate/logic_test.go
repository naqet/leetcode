package main

import (
	"reflect"
	"testing"
)

func logic(nums []int) bool {
    hash:= make(map[int]struct{}, len(nums));

    var result bool;

    for _, v := range nums {
        if _, ok := hash[v]; !ok {
            hash[v] = struct{}{};
        } else {
            result = true
            break;
        }
    }

    return result;
}

type singleTest struct {
	input     []int
	expected bool
}

var tests = []singleTest{
    {[]int{1,2,3,1}, true},
    {[]int{1,2,3,4}, false},
    {[]int{1,1,1,3,3,4,3,2,4,2}, true},
}

func TestLogic(t *testing.T) {
	for _, test := range tests {
		if output := logic(test.input); !reflect.DeepEqual(output, test.expected) {
			t.Error(output, test.expected)
		}
	}
}

