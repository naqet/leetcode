package main

import (
	"reflect"
	"testing"
)

func logic(input []string) [][]string {
    hash := make(map[[26]int][]string, 0);

    for _, w := range input {
        key := [26]int{}

        for _, char := range w {
            key[int(char - 'a')]++;
        }

        hash[key] = append(hash[key], w);
    }

    result := [][]string{};

    for _, group := range hash {
        result = append(result, group);
    }

    return result;
}

type singleTest struct {
	arg1     []string
	expected []string
}

var tests = []singleTest{ }

func TestLogic(t *testing.T) {
	for _, test := range tests {
		if output := logic(test.arg1); !reflect.DeepEqual(output, test.expected) {
			t.Error(output, test.expected)
		}
	}
}

