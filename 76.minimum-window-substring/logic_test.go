package main

import (
	"reflect"
	"testing"
)

func logic(s string, t string) string {
	var result string

	if len(t) > len(s) {
		return result
	}

	need := map[byte]int{}

    for i := range t {
        need[t[i]]++
    }

    rem := len(need);

    l, r := 0, 0;

    for l <= r && r < len(s) {
        right := s[r];

        count, ok := need[right]
        
        if ok {
            if count == 1 {
                rem--
            }

            need[right]--
        }

        for rem <= 0 {
            curr := s[l:r+1];

            if len(result) == 0 || len(curr) < len(result) {
                result = curr
            }

            left := s[l];

            count, ok := need[left]

            if ok {
                if count == 0 {
                    rem++;
                }
                need[left]++
            }
            l++
        }

        r++
    }

	return result
}

type singleTest struct {
	s        string
	t        string
	expected string
}

var tests = []singleTest{
	{"ADOBECODEBANC", "ABC", "BANC"},
	{"cabwefgewcwaefgcf", "cae", "cwae"},
	{"a", "aa", ""},
	{"a", "a", "a"},
	{"aa", "aa", "aa"},
    {"acbbaca", "aba", "baca"},
}

func TestLogic(t *testing.T) {
	for _, test := range tests {
		if output := logic(test.s, test.t); !reflect.DeepEqual(output, test.expected) {
			t.Error(output, test.expected)
		}
	}
}
