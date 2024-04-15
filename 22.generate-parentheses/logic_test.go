package main

import (
	"reflect"
	"testing"
)

func logic(n int) []string {
	var result []string
    generate("", 0, 0, n, &result);
	return result
}

func generate(input string, open int, closed int, n int, result *[]string) {
    if n == open && open == closed {
        *result = append(*result, input);
        return
    }

    if open < n {
        generate(input + "(", open + 1, closed, n, result);
    }

    if closed < open {
        generate(input + ")", open, closed + 1, n, result);
    }
}

type singleTest struct {
	input    int
	expected []string
}

var tests = []singleTest{
	{3, []string{"((()))", "(()())", "(())()", "()(())", "()()()"}},
	{1, []string{"()"}},
}

func TestLogic(t *testing.T) {
	for _, test := range tests {
		if output := logic(test.input); !reflect.DeepEqual(output, test.expected) {
			t.Error(output, test.expected)
		}
	}
}
