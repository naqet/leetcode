package main

import (
	"fmt"
	"reflect"
	"testing"
)

func logic(board [][]byte) bool {
	result := true
	rows := [9][9]bool{}
	cols := [9][9]bool{}
	squares := map[[2]int]map[int]struct{}{}

outer:
	for row := 0; row < 9; row++ {
		for col := 0; col < 9; col++ {
			val := board[row][col]

			if val == '.' {
				continue
			}
			// 48 is ASCII for 0
			// -1 because SUDOKU doesn't have 0
			intVal := int(val) - 49

			// /3 = to have squares indentified
			key := [2]int{row / 3, col / 3}

			if squares[key] == nil {
				squares[key] = map[int]struct{}{}
			}
			_, ok := squares[key][intVal]

			if (row == 1 || row == 2) && (col == 3 || col == 5) {
				fmt.Println(key, row, col, squares[key], intVal)
			}

			if rows[row][intVal] || cols[col][intVal] || ok {
				result = false
				break outer
			} else {
				rows[row][intVal], cols[col][intVal] = true, true
				squares[key][intVal] = struct{}{}
			}
		}
	}

	return result
}

type singleTest struct {
	input    [][]byte
	expected bool
}

var tests = []singleTest{
	{
		[][]byte{
			{'5', '3', '.', '.', '7', '.', '.', '.', '.'},
			{'6', '.', '.', '1', '9', '5', '.', '.', '.'},
			{'.', '9', '8', '.', '.', '.', '.', '6', '.'},
			{'8', '.', '.', '.', '6', '.', '.', '.', '3'},
			{'4', '.', '.', '8', '.', '3', '.', '.', '1'},
			{'7', '.', '.', '.', '2', '.', '.', '.', '6'},
			{'.', '6', '.', '.', '.', '.', '2', '8', '.'},
			{'.', '.', '.', '4', '1', '9', '.', '.', '5'},
			{'.', '.', '.', '.', '8', '.', '.', '7', '9'},
		},
		true,
	},
	{
		[][]byte{
			{'8', '3', '.', '.', '7', '.', '.', '.', '.'},
			{'6', '.', '.', '1', '9', '5', '.', '.', '.'},
			{'.', '9', '8', '.', '.', '.', '.', '6', '.'},
			{'8', '.', '.', '.', '6', '.', '.', '.', '3'},
			{'4', '.', '.', '8', '.', '3', '.', '.', '1'},
			{'7', '.', '.', '.', '2', '.', '.', '.', '6'},
			{'.', '6', '.', '.', '.', '.', '2', '8', '.'},
			{'.', '.', '.', '4', '1', '9', '.', '.', '5'},
			{'.', '.', '.', '.', '8', '.', '.', '7', '9'},
		}, false,
	},
	{
		[][]byte{
			{'.', '.', '.', '.', '5', '.', '.', '1', '.'},
			{'.', '4', '.', '3', '.', '.', '.', '.', '.'},
			{'.', '.', '.', '.', '.', '3', '.', '.', '1'},
			{'8', '.', '.', '.', '.', '.', '.', '2', '.'},
			{'.', '.', '2', '.', '7', '.', '.', '.', '.'},
			{'.', '1', '5', '.', '.', '.', '.', '.', '.'},
			{'.', '.', '.', '.', '.', '2', '.', '.', '.'},
			{'.', '2', '.', '9', '.', '.', '.', '.', '.'},
			{'.', '.', '4', '.', '.', '.', '.', '.', '.'},
		}, false},
}

func TestLogic(t *testing.T) {
	for i, test := range tests {
		if output := logic(test.input); !reflect.DeepEqual(output, test.expected) {
			t.Error(i, output, test.expected)
		}
	}
}
