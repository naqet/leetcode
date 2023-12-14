package main

import (
	"reflect"
	"testing"
)

func logic(s string, t string) bool {
    if len(s) != len(t) {
        return false;
    }
    ascii := make([]int, 26);

    for i := 0; i < len(s); i++ {
        ascii[int(s[i]) - 97]++
        ascii[int(t[i]) - 97]--
    }

    result := true;

    for _, v := range ascii {
        if v != 0 {
            result = false
        }
    }
    return result;
}

type singleTest struct {
	arg1     string
	arg2     string
	expected bool
}

var tests = []singleTest{
    {"anagram", "nagaram", true},
    {"rat", "car", false},
}

func TestLogic(t *testing.T) {
	for _, test := range tests {
		if output := logic(test.arg1, test.arg2); !reflect.DeepEqual(output, test.expected) {
			t.Error(output, test.expected)
		}
	}
}

