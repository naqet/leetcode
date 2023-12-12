package main;

import (
	"strings"
	"testing"
)

func logic(word1 string, word2 string) string {
    var result []string;

    for i, char := range word1 {
        result = append(result, string(char));

        if i < len(word2) {
            result = append(result, string(word2[i]));
        }
    }

    if len(word1) < len(word2) {
        result = append(result, word2[len(result):]);
    }

    return strings.Join(result, "");
} 

type singleTest struct {
    arg1, arg2, expected string;
}

var tests = []singleTest{
    {"abc", "pqr", "apbqcr"},
    {"abc", "pqr", "apbqcr"},
}

func TestLogic(t *testing.T) {
    for _, test := range tests {
        if output := logic(test.arg1, test.arg2); output != test.expected {
            t.Errorf("%q != %q", output, test.expected)
        }
    }
}
