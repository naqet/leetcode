package main

import (
	"reflect"
	"testing"
)

func logic(numbers []int, target int) []int {
	var result []int

	i, j := 0, len(numbers) - 1;

	for i < j && j < len(numbers) {
		sum := numbers[i] + numbers[j]

		if sum == target {
			result = append(result, i+1, j+1)
			break
		} else if sum > target {
            j--
		} else if sum < target {
			i++
		}
	}

	return result
}

type singleTest struct {
	numbers  []int
	target   int
	expected []int
}

var tests = []singleTest{
	{
		[]int{2, 7, 11, 15},
		9,
		[]int{1, 2},
	},
	{
		[]int{2, 3, 4},
		6,
		[]int{1, 3},
	},
	{
		[]int{-1, 0},
		-1,
		[]int{1, 2},
	},
	{
		[]int{5, 25, 75},
		100,
		[]int{2, 3},
	},
	{
		[]int{3, 24, 50, 79, 88, 150, 345},
		200,
		[]int{3, 6},
	},
}

func TestLogic(t *testing.T) {
	for _, test := range tests {
		if output := logic(test.numbers, test.target); !reflect.DeepEqual(output, test.expected) {
			t.Error(output, test.expected)
		}
	}
}
