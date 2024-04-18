package main

import (
	"reflect"
	"testing"
)

type Deque []int

func (d *Deque) leftVal() (int, bool) {
	if len(*d) <= 0 {
		return 0, false
	}

	return (*d)[0], true
}

func (d *Deque) rightVal() (int, bool) {
	if len(*d) <= 0 {
		return 0, false
	}

	return (*d)[len(*d)-1], true
}

func (d *Deque) pushRight(val int) {
	*d = append(*d, val)
}

func (d *Deque) pushLeft(val int) {
	*d = append([]int{val}, *d...)
}

func (d *Deque) popLeft() (int, bool) {
	if len(*d) <= 0 {
		return 0, false
	}

	left := (*d)[0]
	*d = (*d)[1:]

	return left, true
}

func (d *Deque) popRight() (int, bool) {
	if len(*d) <= 0 {
		return 0, false
	}

	right := (*d)[len(*d)-1]
	*d = (*d)[:len(*d)-1]

	return right, true
}

func (d *Deque) push(val int) {
	right, ok := d.rightVal()

	if !ok || val <= right {
		d.pushRight(val)
		return
	}

	for ok && right < val {
		d.popRight()

		right, ok = d.rightVal()
	}

	d.pushRight(val)
}

func logic(nums []int, k int) []int {
	var result []int
	deque := Deque{}
	l, r := 0, 0

	for r < len(nums) {
		if r-l+1 != k {
			deque.push(nums[r])
			r++
			continue
		}

		deque.push(nums[r])

		left, _ := deque.leftVal()

		result = append(result, left)

		if nums[l] == left {
			deque.popLeft()
		}

		r++
		l++
	}

	return result
}

type singleTest struct {
	nums     []int
	k        int
	expected []int
}

var tests = []singleTest{
	{[]int{1, 3, -1, -3, 5, 3, 6, 7}, 3, []int{3, 3, 5, 5, 6, 7}},
	{[]int{1}, 1, []int{1}},
	{[]int{1, -1}, 1, []int{1, -1}},
	{[]int{1, 3, 1, 2, 0, 5}, 3, []int{3, 3, 2, 5}},
	{[]int{-7, -8, 7, 5, 7, 1, 6, 0}, 4, []int{7, 7, 7, 7, 7}},
}

func TestLogic(t *testing.T) {
	for _, test := range tests {
		if output := logic(test.nums, test.k); !reflect.DeepEqual(output, test.expected) {
			t.Error(output, test.expected)
		}
	}
}
