package main

import (
	"reflect"
	"testing"
)

var chars = map[byte]string{
	'2': "abc",
	'3': "def",
	'4': "ghi",
	'5': "jkl",
	'6': "mno",
	'7': "pqrs",
	'8': "tuv",
	'9': "wxyz",
}

func letterCombinations(nums string) []string {
    result := []string{}

    var backtrack func(idx int, s, sub string,  result *[]string)
    backtrack = func(idx int, s string, sub string, result *[]string) {
        if idx == len(s) {
            if len(sub) > 0 {
                *result = append(*result, sub)
            }
            return
        }

        numChars := chars[s[idx]]

        for _, curr := range numChars {
            sub += string(curr)
            backtrack(idx+1, s, sub, result)
            sub = sub[:len(sub)-1]
        }
    }

    backtrack(0, nums, "", &result)

	return result
}


type singleTest struct {
	input    string
	expected []string
}

var tests = []singleTest{
	{"23", []string{"ad", "ae", "af", "bd", "be", "bf", "cd", "ce", "cf"}},
	{"", []string{}},
	{"2", []string{"a", "b", "c"}},
}

func TestLogic(t *testing.T) {
	for i, test := range tests {
		if output := letterCombinations(test.input); !reflect.DeepEqual(output, test.expected) {
            t.Error("Nr:", i , "Output:", output, "Expected:", test.expected)
		}
	}
}
