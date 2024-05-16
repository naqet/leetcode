package main

import (
	"reflect"
	"testing"
)

func partition(s string) [][]string {
	var result [][]string
	dfs(0, s, &[]string{}, &result)
	return result
}

func dfs(idx int, s string, subs *[]string, result *[][]string) {
	if idx == len(s) {
		sub := make([]string, len(*subs))
		copy(sub, *subs)
		*result = append(*result, sub)
		return
	}

    for j := idx; j < len(s); j++ {
        if isPalindrome(s[idx:j+1]) {
            *subs = append(*subs, s[idx:j+1])
            dfs(j+1, s, subs, result)
            *subs = (*subs)[:len(*subs)-1]
        }
    }
}

func isPalindrome(s string) bool {
	if len(s) == 0 {
		return false
	}

	i, j := 0, len(s)-1

	result := true
	for i < j {
		if s[i] != s[j] {
			result = false
			break
		}

		i++
		j--
	}

	return result
}

type singleTest struct {
	input    string
	expected [][]string
}

var tests = []singleTest{
	{"aab", [][]string{{"a", "a", "b"}, {"aa", "b"}}},
	{"a", [][]string{{"a"}}},
}

func TestLogic(t *testing.T) {
	for _, test := range tests {
		if output := partition(test.input); !reflect.DeepEqual(output, test.expected) {
			t.Error(output, test.expected)
		}
	}
}
