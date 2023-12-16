package main

import (
	"reflect"
	"testing"
)

func logic(input string, amount int) int {
    maxLength := 0;

    count := make(map[byte]int, 26);
    maxAmount := 0;

    l := 0;

    for r := 0; r < len(input); r++ {
        length := r - l + 1;
        count[input[r] - 'A']++;
        if count[input[r] - 'A'] > maxAmount {
            maxAmount = count[input[r] - 'A']
        }

        if length - maxAmount > amount {
            count[input[l] - 'A']--;
            l++;
        } else if length > maxLength {
            maxLength = length;
        }
    }
    return maxLength;
}

type singleTest struct {
	arg1     string
	arg2     int
	expected int
}

var tests = []singleTest{
    //{"ABAB", 2, 4},
    {"AABABBA", 1, 4},
}

func TestLogic(t *testing.T) {
	for _, test := range tests {
		if output := logic(test.arg1, test.arg2); !reflect.DeepEqual(output, test.expected) {
			t.Error(output, test.expected)
		}
	}
}

