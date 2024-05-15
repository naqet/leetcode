package main

import (
	"reflect"
	"testing"
)

func exists(board [][]byte, word string) bool {
	rows, cols := len(board), len(board[0])
	path := map[[2]int]struct{}{}

	var dfs func(r, c, idx int) bool

	dfs = func(r, c, idx int) bool {
		if idx == len(word) {
			return true
		}

		if r < 0 || c < 0 {
			return false
		}

		if r >= rows || c >= cols {
			return false
		}

		pathKey := [2]int{r, c}
		if _, ok := path[pathKey]; ok {
			return false
		}

		if word[idx] != board[r][c] {
			return false
		}

		path[pathKey] = struct{}{}

		pathRes := false

		if dfs(r, c+1, idx+1) || dfs(r+1, c, idx+1) || dfs(r, c-1, idx+1) || dfs(r-1, c, idx+1) {
			pathRes = true
		}

		delete(path, pathKey)

		return pathRes
	}

	result := false

	for r := range rows {
		for c := range cols {
			if dfs(r, c, 0) {
				result = true
				break
			}
		}
	}

	return result
}

type singleTest struct {
	input    [][]byte
	word     string
	expected bool
}

var tests = []singleTest{
	{[][]byte{
		{'A', 'B', 'C', 'E'},
		{'S', 'F', 'C', 'S'},
		{'A', 'D', 'E', 'E'},
	}, "ABCCED", true},
}

func TestLogic(t *testing.T) {
	for _, test := range tests {
		if output := exists(test.input, test.word); !reflect.DeepEqual(output, test.expected) {
			t.Error(output, test.expected)
		}
	}
}
