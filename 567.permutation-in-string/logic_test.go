package main

import (
	"reflect"
	"testing"
)

func logic(s1 string, s2 string) bool {
	var result bool
	if len(s2) < len(s1) {
		return result
	}

	one := [26]int{}

	for _, char := range s1 {
		one[char-'a']++
	}


	i, j := 0, 0

	two := [26]int{}

	for i <= j && j < len(s2) {
		left := s2[i]
		right := s2[j]

        two[right-'a']++

		if j-i+1 != len(s1) {
			j++
			continue
		}

		if one == two {
			result = true
			break
		}

		i++
		two[left-'a']--

		j++
	}

	return result
}

type singleTest struct {
	s1       string
	s2       string
	expected bool
}

var tests = []singleTest{
	{"ab", "eidbaooo", true},
	{"ab", "eidboaoo", false},
}

func TestLogic(t *testing.T) {
	for _, test := range tests {
		if output := logic(test.s1, test.s2); !reflect.DeepEqual(output, test.expected) {
			t.Error(output, test.expected)
		}
	}
}
