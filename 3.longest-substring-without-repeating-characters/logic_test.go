package main

import (
	"reflect"
	"testing"
)

func logic(s string) int {
	if len(s) <= 1 {
		return len(s)
	}

	var result int
	hash := map[byte]int{}

	i, j := 0, 0

	for i <= j && j < len(s) {
		left := s[i]
		right := s[j]
		count, ok := hash[right]

		if !ok || count == 0 {
			hash[right]++
			result = max(result, j-i+1)
			j++
		} else {
			hash[left]--
			i++
		}
	}

	return result
}

type singleTest struct {
	input    string
	expected int
}

var tests = []singleTest{
	{"abcabcbb", 3},
	{"bbbb", 1},
	{"pwwkew", 3},
}

func TestLogic(t *testing.T) {
	for _, test := range tests {
		if output := logic(test.input); !reflect.DeepEqual(output, test.expected) {
			t.Error(output, test.expected)
		}
	}
}
