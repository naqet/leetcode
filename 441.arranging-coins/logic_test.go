package main

import (
	"reflect"
	"testing"
)

func logic(n int) int {
    level := 0;

    for n > level {
        level++;
        n -= level;
    }
    return level;
}

type singleTest struct {
	arg1     int
	expected int
}

var tests = []singleTest{
    {5, 2},
}

func TestLogic(t *testing.T) {
	for _, test := range tests {
		if output := logic(test.arg1); !reflect.DeepEqual(output, test.expected) {
			t.Error(output, test.expected)
		}
	}
}

