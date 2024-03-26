package main

import (
	"testing"
)

type TimeMap struct {
	db    map[int][2]string
	times []int
}

func Constructor() TimeMap {
	return TimeMap{map[int][2]string{}, []int{}}
}

func (this *TimeMap) Set(key string, value string, timestamp int) {
	this.db[timestamp] = [2]string{key, value}
	this.times = append(this.times, timestamp)
}

func (this *TimeMap) Get(key string, timestamp int) string {
	result := ""

	for i := len(this.times) - 1; i >= 0; i-- {
		val := this.times[i]
		if val <= timestamp {
			storedKey := this.db[val][0]
			if key == storedKey {
				result = this.db[val][1]
				break
			}
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
