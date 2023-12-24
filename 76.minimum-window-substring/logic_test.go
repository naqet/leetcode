package main

import (
	"reflect"
	"testing"
)

func logic(s string, t string) string {
	need := map[byte]int{}
	rem := 0

	for i := 0; i < len(t); i++ {
		rem++
		need[t[i]]++
	}

	l, r := 0, 0

	result := ""

	for r < len(s) {
		right := s[r]

		c, ok := need[right]

		if ok {
			if c > 0 {
				rem--
			}
			need[right]--
		}

		for rem <= 0 {
			res := s[l : r+1]

			if len(result) > len(res) || len(result) == 0 {
				result = res
			}
			left := s[l]
			_, ok := need[left]

			if ok {
				need[left]++

                c := need[left]

                if c > 0 {
                    rem++;
                }

			}

			l++
		}

		r++
	}
	return result
}

type singleTest struct {
	arg1     string
	arg2     string
	expected string
}

var tests = []singleTest{
	{"ADOBECODEBANC", "ABC", "BANC"},
	{"cabwefgewcwaefgcf", "cae", "cwae"},
	{"a", "a", "a"},
	{"a", "aa", ""},
	{"aa", "aa", "aa"},
}

func TestLogic(t *testing.T) {
	for _, test := range tests {
		if output := logic(test.arg1, test.arg2); !reflect.DeepEqual(output, test.expected) {
			t.Error(output, test.expected)
		}
	}
}
