package main

import (
	"reflect"
	"strconv"
	"testing"
)

var (
	add   = "+"
	sub   = "-"
	multi = "*"
	div   = "/"
)

func logic(input []string) int {
	stack := []int{}

	for _, ch := range input {
		operator := ch == add || ch == sub || ch == multi || ch == div

		if !operator {
            num, _ := strconv.ParseInt(ch, 10, 32);
			stack = append(stack, int(num))
			continue
		}

		first, second := stack[len(stack)-2], stack[len(stack)-1]
		stack = stack[:len(stack)-2]

		var result int
		switch ch {
		case add:
			result = first + second
		case sub:
			result = first - second
		case multi:
			result = first * second
		case div:
			result = first / second
		}

		stack = append(stack, result);
	}

	return stack[0]
}

type singleTest struct {
	input    []string
	expected int
}

var tests = []singleTest{
	{[]string{"2", "1", "+", "3", "*"}, 9},
	{[]string{"4", "13", "5", "/", "+"}, 6},
	{[]string{"10", "6", "9", "3", "+", "-11", "*", "/", "*", "17", "+", "5", "+"}, 22},
}

func TestLogic(t *testing.T) {
	for _, test := range tests {
		if output := logic(test.input); !reflect.DeepEqual(output, test.expected) {
			t.Error(output, test.expected)
		}
	}
}
