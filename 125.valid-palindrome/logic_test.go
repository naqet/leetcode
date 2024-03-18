package main

import (
	"reflect"
	"strings"
	"testing"
	"unicode"
)

func logic(s string) bool {
    s = strings.ToLower(s)
    var cleanRunes []rune;

    for _, ch := range s {
        if unicode.IsLetter(ch) || unicode.IsDigit(ch) {
            cleanRunes = append(cleanRunes, ch);
        }
    }

    clean := string(cleanRunes);

    result := true;
    i, j := 0, len(clean) - 1;

    for i < len(clean) && j >= 0 {
        if clean[i] != clean[j] {
            result = false;
            break;
        }
        i++;
        j--;
    }

    return result;
}

type singleTest struct {
	input     string
	expected bool
}

var tests = []singleTest{
    {"A man, a plan, a canal: Panama", true},
    {"race a car", false},
    {" ", true},
}

func TestLogic(t *testing.T) {
	for _, test := range tests {
		if output := logic(test.input); !reflect.DeepEqual(output, test.expected) {
			t.Error(output, test.expected)
		}
	}
}

