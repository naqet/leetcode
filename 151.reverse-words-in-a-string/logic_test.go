package main

import (
	"strings"
	"testing"
)

func logic(s string) string {
    ss := strings.Fields(s);

    for i, j := 0, len(ss)-1; i < j; i, j = i+1, j-1 {
        ss[i], ss[j] = ss[j], ss[i]
    }
    return strings.Join(ss, " ");
}

type singleTest struct {
	arg1     string
	expected string
}

var tests = []singleTest{
    {"the sky is blue", "blue is sky the"},
    {"  hello world  ", "world hello"},
    {"a good   example", "example good a"},
}

func TestLogic(t *testing.T) {
	for _, test := range tests {
		if output := logic(test.arg1); output != test.expected { 
			t.Errorf("%q is not equal to %q", output, test.expected)
		}
	}
}
