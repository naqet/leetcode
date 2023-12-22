package main

import (
	"reflect"
	"testing"
)

func logic(s1 string, s2 string) bool {
	if len(s1) > len(s2) {
		return false
	}

	result := false

	s1c := [26]int{}

	for _, v := range s1 {
		s1c[v-'a']++
	}

	count := [26]int{}

	l, r := 0, 0

	for r < len(s2) {
		right := s2[r]
		count[right-'a']++

		if r-l+1 != len(s1) {
			r++
			continue
		}

		if s1c == count {
			result = true
			break
		}

		left := s2[l]
		count[left-'a']--
		r++
		l++
	}

	return result
}

type singleTest struct {
	arg1     string
	arg2     string
	expected bool
}

var tests = []singleTest{
	{"ab", "eidbaooo", true},
	{"ba", "eidboaoo", false},
	{"rfc", "edcrfv", true},
	{"rfv", "rfvsadfsdafsdarfv", true},
}

func TestLogic(t *testing.T) {
	for _, test := range tests {
		if output := logic(test.arg1, test.arg2); !reflect.DeepEqual(output, test.expected) {
			t.Error(output, test.expected)
		}
	}
}
