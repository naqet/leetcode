package main

import (
	"reflect"
	"testing"
)

func logic(input []int, k int) []int {
    result := []int{};

    l, r := 0, 0;
    dq := Dq{[]int{}};

    for r < len(input) {
        right := input[r];

        maxRight, ok := dq.rightVal();
        for len(dq.arr) > 0 && ok && maxRight < right {
            dq.pop();
            mR, ok := dq.rightVal();

            if ok {
                maxRight = mR;
            }
        }

        dq.push(right);

        maxLeft, ok := dq.leftVal();

        if ok && r + 1 >= k {
            result = append(result, maxLeft)
            left := input[l];
            if left == maxLeft {
                dq.popLeft();
            }

            l++;
        }

        r++;
    }


	return result
}

type Dq struct {
	arr []int
}

func (dq *Dq) pushLeft(val int) {
	dq.arr = append([]int{val}, dq.arr...)
}

func (dq *Dq) push(val int) {
	dq.arr = append(dq.arr, val)
}

func (dq *Dq) popLeft() {
	dq.arr = dq.arr[1:]
}

func (dq *Dq) pop() {
	dq.arr = dq.arr[:len(dq.arr)-1]
}

func (dq *Dq) leftVal() (int, bool) {
	if len(dq.arr) > 0 {
		return dq.arr[0], true
	}
	return 0, false
}

func (dq *Dq) rightVal() (int, bool) {
	if len(dq.arr) > 0 {
		return dq.arr[len(dq.arr)-1], true
	}
	return 0, false
}

type singleTest struct {
	arg1     []int
	arg2     int
	expected []int
}

var tests = []singleTest{
	{[]int{1, 3, -1, -3, 5, 3, 6, 7}, 3, []int{3, 3, 5, 5, 6, 7}},
	{[]int{1}, 1, []int{1}},
	{[]int{1, -1}, 1, []int{1, -1}},
}

func TestLogic(t *testing.T) {
	for _, test := range tests {
		if output := logic(test.arg1, test.arg2); !reflect.DeepEqual(output, test.expected) {
			t.Error(output, test.expected)
		}
	}
}
