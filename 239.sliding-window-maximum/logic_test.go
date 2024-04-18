package main

import (
	"reflect"
	"testing"
)

type Node struct {
	value int
	left  *Node
	right *Node
}

type Deque struct {
	head   *Node
	tail   *Node
	length int
}

func (d *Deque) leftVal() (int, bool) {
	if d.head == nil {
		return 0, false
	}

	return d.head.value, true
}

func (d *Deque) rightVal() (int, bool) {
	if d.tail == nil {
		return 0, false
	}

	return d.tail.value, true
}

func (d *Deque) pushRight(val int) {
	newNode := &Node{value: val}
	d.length++
	if d.tail == nil {
		d.head = newNode
		d.tail = newNode
		return
	}

	d.tail.right = newNode
	newNode.left = d.tail
	d.tail = newNode
}

func (d *Deque) popLeft() (int, bool) {
	if d.length == 0 {
		return 0, false
	}

	d.length--

	head := d.head
	if d.length == 0 {
		head := d.head
		d.head = nil
		d.tail = nil
		return head.value, true
	}

	d.head.right.left = nil
	d.head = d.head.right

	return head.value, true
}

func (d *Deque) popRight() (int, bool) {
	if d.length == 0 {
		return 0, false
	}

	d.length--

	tail := d.tail
	if d.length == 0 {
		tail := d.tail
		d.head = nil
		d.tail = nil
		return tail.value, true
	}

	d.tail.left.right = nil
	d.tail = d.tail.left

	return tail.value, true
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
