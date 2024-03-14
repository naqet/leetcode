package main

import (
	"reflect"
	"testing"
)

func logic(s string, t string) bool {
    if len(s) != len(t) {
        return false;
    }

	arr := make([]int, 26)

    for i := 0; i < len(s); i++ {
		arr[s[i] - 97]++
		arr[t[i] - 97]--
	}

    result := true;

    for _, v := range arr {
        if v != 0 {
            result = false
            break;
        }
    }

	return result
}

type singleTest struct {
	s        string
	t        string
	expected bool
}

var tests = []singleTest{
    {"a", "b", false},
    {"a", "ab", false},
    {"cat", "tca", true},
}

func TestLogic(t *testing.T) {
	for _, test := range tests {
		if output := logic(test.s, test.t); !reflect.DeepEqual(output, test.expected) {
			t.Error(output, test.expected)
		}
	}
}
