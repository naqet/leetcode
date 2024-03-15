package main

import (
	"reflect"
	"testing"
)

func logic(strs []string) [][]string {
	result := make([][]string, 0)
    hash := make(map[[26]int][]string, 0)

    for _, s := range strs {
        arr := [26]int{};
        for _, ch := range s {
            arr[ch-97]++
        }
        _, ok := hash[arr];

        if ok {
            hash[arr] = append(hash[arr], s)
        } else {
            hash[arr] = []string{s}
        }
    }

    for _, anagrams := range hash {
        result = append(result, anagrams)
    }

	return result
}

type singleTest struct {
	input    []string
	expected [][]string
}

var tests = []singleTest{
	{
		[]string{"eat", "tea", "tan", "ate", "nat", "bat"},
		[][]string{ {"eat", "tea", "ate"}, {"tan", "nat"}, {"bat"}},
	},
	{
		[]string{""},
		[][]string{{""}},
	},
	{
		[]string{"a"},
		[][]string{{"a"}},
	},
}

func TestLogic(t *testing.T) {
	for _, test := range tests {
		if output := logic(test.input); !reflect.DeepEqual(output, test.expected) {
			t.Error(output, test.expected)
		}
	}
}
