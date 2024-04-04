package main

import (
	"reflect"
	"testing"
)

func logic(s string) bool {
    result := true;

	stack := []rune{};

    pairs := map[rune]rune{
        '(' : ')',
        '{' : '}',
        '[' : ']',
    }

	for _, ch := range s {
        if _, ok := pairs[ch]; ok {
            stack = append(stack, ch);
        } else if len(stack) == 0 || pairs[stack[len(stack) - 1]] != ch {
            result = false 
            break;
        } else {
            stack = stack[:len(stack)-1]
        }
	}

    if result && len(stack) != 0 {
        result = false;
    }

	return result
}

type singleTest struct {
	input    string
	expected bool
}

var tests = []singleTest{
	{"()", true},
	{"()[]{}", true},
	{"([", false},
}

func TestLogic(t *testing.T) {
	for _, test := range tests {
		if output := logic(test.input); !reflect.DeepEqual(output, test.expected) {
			t.Error(output, test.expected)
		}
	}
}
