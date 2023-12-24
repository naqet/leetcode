package main

import (
	"reflect"
	"testing"
)

func logic(s string, p string) []int {
    var result []int;

    need := [26]int{};

    for i := 0; i < len(p); i++ {
        need[p[i] - 'a']++;
    }

    l, r := 0, 0;

    have := [26]int{};

    for r < len(s) {
        right := s[r];
        left := s[l];

        have[right - 'a']++;

        if need == have {
            result = append(result, l);
            have[left - 'a']--;
            l++;
        } else if (r - l + 1) % len(p) == 0 {
            have[left - 'a']--;
            l++;
        }

        r++;
    }
    return result;
}

type singleTest struct {
	arg1     string
	arg2     string
	expected []int
}

var tests = []singleTest{
    {"cbaebabacd", "abc", []int{0, 6}},
}

func TestLogic(t *testing.T) {
	for _, test := range tests {
		if output := logic(test.arg1, test.arg2); !reflect.DeepEqual(output, test.expected) {
			t.Error(output, test.expected)
		}
	}
}

