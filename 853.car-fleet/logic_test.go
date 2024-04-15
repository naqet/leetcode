package main

import (
	"reflect"
	"slices"
	"testing"
)

type Car struct {
	position float64
	speed    float64
}

func logic(target int, position []int, speed []int) int {
	var result int
	cars := []Car{}

	for i := 0; i < len(position); i++ {
		cars = append(cars, Car{float64(position[i]), float64(speed[i])})
	}

	slices.SortFunc(cars, func(a Car, b Car) int {
		if a.position < b.position {
			return 1
		} else if a.position == b.position {
			return 0
		}
		return -1
	})

    var last float64;

	for i, car := range cars {
		result++
		time := (float64(target) - car.position) / car.speed

        if i == 0 {
            last = time
            continue;
        }

        if time <= last {
            result--;
        } else {
            last = time;
        }
	}

	return result
}

type singleTest struct {
	input    int
	position []int
	speed    []int
	expected int
}

var tests = []singleTest{
	{12, []int{10, 8, 0, 5, 3}, []int{2, 4, 1, 1, 3}, 3},
	{10, []int{3}, []int{3}, 1},
	{100, []int{0, 2, 4}, []int{4, 2, 1}, 1},
    {10, []int{6,8}, []int{3,2}, 2},
}

func TestLogic(t *testing.T) {
	for _, test := range tests {
		if output := logic(test.input, test.position, test.speed); !reflect.DeepEqual(output, test.expected) {
			t.Error(output, test.expected)
		}
	}
}
