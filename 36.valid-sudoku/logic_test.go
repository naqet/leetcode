package main

import (
	"reflect"
	"testing"
)

func logic(input [][]byte) bool {
    result := true;

    rows := map[int]map[byte]bool{}
    cols := map[int]map[byte]bool{}
    squares := map[[2]int]map[byte]bool{}

    outer: for rIdx := range input {
       for cIdx := range input {
           val := input[rIdx][cIdx]
           if val == '.' {
               continue;
           }

           _, inRow := rows[rIdx][val];

           if inRow {
               result = false;
               break outer;
           } else {
               _, rowExists := rows[rIdx];

               if rowExists {
                   rows[rIdx][val] = true;
               } else {
                   rows[rIdx] = map[byte]bool{val: true};
               }
           }

           _, inCol := cols[cIdx][val];

           if inCol {
               result = false;
               break outer;
           } else {
               _, colExists := cols[cIdx];

               if colExists {
                   cols[cIdx][val] = true;
               } else {
                   cols[cIdx] = map[byte]bool{val: true};
               }
           }

           key := [2]int{rIdx / 3, cIdx / 3};
           _, inSquare := squares[key][val];

           if inSquare {
               result = false;
               break outer;
           } else {
               _, squareExists := squares[key];

               if squareExists {
                   squares[key][val] = true;
               } else {
                   squares[key] = map[byte]bool{val: true};
               }
           }
       } 
    }

    return result;
}

type singleTest struct {
	arg1     [][]byte
	expected bool
}

var tests = []singleTest{ }

func TestLogic(t *testing.T) {
	for _, test := range tests {
		if output := logic(test.arg1); !reflect.DeepEqual(output, test.expected) {
			t.Error(output, test.expected)
		}
	}
}
