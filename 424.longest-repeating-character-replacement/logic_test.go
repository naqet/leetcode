package main

import (
	"reflect"
	"testing"
)

func logic(s string, k int) int {
    if len(s) < 2 {
        return len(s);
    }

    result := 1;
    arr := [26]int{};
    maxF := 0;

    i, j := 0, 0;

    for i <= j && j < len(s) {
        left := s[i] - 65
        right := s[j] - 65

        arr[right]++
        maxF = max(maxF, arr[right])
        length := j - i + 1;

        if length - maxF <= k {
            result = max(result, length);
            j++
        } else {
            i++
            arr[right]--
            arr[left]--
        }
    }


	return result
}

type singleTest struct {
	input    string
	k        int
	expected int
}

var tests = []singleTest{
	{"ABAB", 2, 4},
	{"AABABBA", 1, 4},
	{"ABAA", 0, 2},
}

func TestLogic(t *testing.T) {
	for _, test := range tests {
		if output := logic(test.input, test.k); !reflect.DeepEqual(output, test.expected) {
			t.Error(output, test.expected)
		}
	}
}
