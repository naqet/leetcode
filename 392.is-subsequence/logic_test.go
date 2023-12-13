package main

import (
	"reflect"
	"testing"
)

func logic(s string, t string) bool {
    j := 0
    for i := 0; i < len(t); i++ {
        if j < len(s) && t[i] == s[j] {
            j++
        }
    }
    return j == len(s);
}

type singleTest struct {
	arg1     string
	arg2     string
	expected bool
}

var tests = []singleTest{
    {"abc", "ahbgdc", true},
    {"axc", "ahbgdc", false},
}

func TestLogic(t *testing.T) {
	for _, test := range tests {
		if output := logic(test.arg1, test.arg2); !reflect.DeepEqual(output, test.expected) {
			t.Error(output, test.expected)
		}
	}
}

