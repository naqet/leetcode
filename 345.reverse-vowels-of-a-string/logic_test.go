package main

import (
	"fmt"
	"strings"
	"testing"
)

type swap struct {
    idx int
    letter string
}

func logic(s string) string {
    vowels := []rune{'a', 'e', 'i', 'o', 'u', 'A', 'E', 'I', 'O', 'U'}

    idxToLetter := make([]swap, 0)

    for i, char := range s {
        if isInSlice(vowels, char) {
            idxToLetter = append(idxToLetter, swap{i, string(char)})
        }
    }

    fmt.Println(idxToLetter)

    result := strings.Split(s, "");

    for i, s := range idxToLetter {
        toUpdate := idxToLetter[len(idxToLetter) - 1 - i];
        result[s.idx] = toUpdate.letter;
        result[toUpdate.idx] = s.letter;
    }

    return strings.Join(result, "")
}

func isInSlice(arr []rune, target rune) bool {
    found := false;
    for _, v := range arr {
        if v == target {
            found = true;
            break;
        }
    }

    return found;
}

type singleTest struct {
	arg1     string
	expected string
}

var tests = []singleTest{
    {"hello", "holle"},
    {"leetcode", "leotcede"},
    {"aA", "Aa"},
}

func TestLogic(t *testing.T) {
	for _, test := range tests {
		if output := logic(test.arg1); output != test.expected { 
			t.Errorf("%q is not equal to %q", output, test.expected)
		}
	}
}
