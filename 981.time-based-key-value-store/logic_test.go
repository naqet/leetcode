package main

import (
	"testing"
)

type TimeMap struct {
	db    map[string][]Pair
}

type Pair struct {
    val string
    time int
}

func Constructor() TimeMap {
	return TimeMap{map[string][]Pair{}}
}

func (this *TimeMap) Set(key string, value string, timestamp int) {
    _, ok := this.db[key];

    if !ok {
        this.db[key] = []Pair{};
    }

	this.db[key] = append(this.db[key], Pair{value, timestamp})
}

func (this *TimeMap) Get(key string, timestamp int) string {
	result := ""

    pairs, ok := this.db[key];

    if !ok {
        return result;
    }

    lo, hi:= 0, len(pairs) - 1;

    for lo <= hi {
        mid := lo + (hi - lo) / 2
        time := pairs[mid].time

        if time <= timestamp {
            result = pairs[mid].val;
            lo = mid + 1;
        } else {
            hi = mid - 1
        }
    }
	return result
}

type singleTest struct {
	input    string
	expected bool
}

var tests = []singleTest{}

func TestLogic(t *testing.T) {
	timeMap := Constructor()

	timeMap.Set("foo", "bar", 1)
	timeMap.Set("hello", "world", 2)
	res := timeMap.Get("foo", 1)
	if res != "bar" {
		t.Error("Not valid value at 1")
		return
	}

	res = timeMap.Get("hello", 2)
	if res != "world" {
		t.Error("Not valid value at 1A")
		return
	}
	res = timeMap.Get("foo", 3)
	if res != "bar" {
		t.Error("Not valid value at 2")
		return
	}
	timeMap.Set("foo", "bar2", 4)
	res = timeMap.Get("foo", 4)

	if res != "bar2" {
		t.Error("Not valid value at 3")
		return
	}
	res = timeMap.Get("foo", 5)

	if res != "bar2" {
		t.Error("Not valid value at 4")
		return
	}
}
