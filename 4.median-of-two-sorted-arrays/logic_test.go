package main

import (
	"math"
	"reflect"
	"testing"
)

func logic(nums1 []int, nums2 []int) float64 {
	var result float64
	if len(nums2) < len(nums1) {
		nums1, nums2 = nums2, nums1
	}

	length := len(nums1) + len(nums2)

	lo := 0
	hi := len(nums1)

	for lo <= hi {
		pleft := lo + (hi-lo)/2
		pright := ((length + 1) / 2) - pleft

        maxLeft := math.MinInt;
        if pleft > 0 {
            maxLeft = nums1[pleft - 1]
        }

        minLeft := math.MaxInt;
        if pleft < len(nums1) {
            minLeft = nums1[pleft];
        }

        maxRight := math.MinInt;
        if pright > 0 {
            maxRight = nums2[pright - 1]
        }

        minRight := math.MaxInt;
        if pright < len(nums2) {
            minRight = nums2[pright];
        }

		if maxLeft <= minRight && maxRight <= minLeft {
			realMax := float64(max(maxLeft, maxRight))

			if length%2 == 0 {
				realMin := float64(min(minLeft, minRight))
				result = (realMax + realMin) / 2
			} else {
				result = realMax
			}
			break
		} else if maxLeft > minRight {
			hi = pleft - 1
		} else {
			lo = pleft + 1
		}
	}

	return result
}

func max(one int, two int) int {
	if one >= two {
		return one
	}
	return two
}

func min(one int, two int) int {
	if one <= two {
		return one
	}
	return two
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
