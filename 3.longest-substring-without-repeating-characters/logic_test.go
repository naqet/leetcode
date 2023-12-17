package main

import (
	"reflect"
	"testing"
)

func logic(input string) int {
	count := make(map[byte]bool, 0)
    maxLength := 0;

	l, r := 0, 0;

	for r < len(input) {
		val := input[r]
		_, ok := count[val]

		if ok {
            delete(count, input[l]);
            l++;
        } else {
            count[val] = true;

            if r - l + 1 > maxLength {
                maxLength = r - l + 1
            }
            r++;
        }
	}

	return maxLength
}

type singleTest struct {
	arg1     string
	expected int
}

var tests = []singleTest{
    //{"abcabcbb", 3},
    //{"bbbb", 1},
    {"pwwkew", 3},
}

func TestLogic(t *testing.T) {
	for _, test := range tests {
		if output := logic(test.arg1); !reflect.DeepEqual(output, test.expected) {
			t.Error(output, test.expected)
		}
	}
}
