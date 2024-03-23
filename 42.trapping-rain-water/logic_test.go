package main

import (
	"reflect"
	"testing"
)

func logic(height []int) int {
	if len(height) < 3 {
		return 0
	}

	var result int
	var rocks int

	i, j := 0, 1

	for j < len(height) {
		left := height[i]
		right := height[j]

		if i == 0 && left == 0 {
			i++
			j++
			continue
		}

		if left > right {
			if j == len(height)-1 {
				rocks = 0
				i2, j2 := j-1, j

				for i <= i2 {
					l2 := height[i2]
					r2 := height[j2]

					if j2 == len(height)-1 && r2 == 0 {
						i2--
						j2--
						continue
					}

					if r2 > l2 {
						rocks += l2
                        i2--
					} else {
						result += r2*(j2-i2-1) - rocks
						rocks = 0
						j2 = i2
						i2 = i2 - 1
					}
				}
			}

			rocks += right
			j++
		} else if left <= right {
			result += left*(j-i-1) - rocks
			rocks = 0
			i = j
			j = j + 1
		}
	}

	return result
}

type singleTest struct {
	input    []int
	expected int
}

var tests = []singleTest{
	{[]int{0, 1, 0, 2, 1, 0, 1, 3, 2, 1, 2, 1}, 6},
	{[]int{4,2,0,3,2,5}, 9},
}

func TestLogic(t *testing.T) {
	for _, test := range tests {
		if output := logic(test.input); !reflect.DeepEqual(output, test.expected) {
			t.Error(output, test.expected)
		}
	}
}
