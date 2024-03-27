package main

import (
	"reflect"
	"slices"
	"testing"
)

func logic(nums1 []int, nums2 []int) float64 {
	var result float64

    arr := slices.Concat(nums1, nums2);

    slices.Sort(arr);

    if len(arr) % 2 == 0 {
        first := float64(arr[len(arr) / 2]);
        second := float64(arr[(len(arr) / 2) - 1]);
        result = (first + second) / 2
    } else {
        result = float64(arr[len(arr) / 2]);
    }
	return result
}

type singleTest struct {
	nums1    []int
	nums2    []int
	expected float64
}

var tests = []singleTest{
	{[]int{1, 3}, []int{2}, 2},
	{[]int{1, 2}, []int{3, 4}, 2.5},
}

func TestLogic(t *testing.T) {
	for _, test := range tests {
		if output := logic(test.nums1, test.nums2); !reflect.DeepEqual(output, test.expected) {
			t.Error(output, test.expected)
		}
	}
}
