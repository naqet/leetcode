package main

import (
	"testing"
)

type WordDictionary struct {
	children map[byte]*WordDictionary
	isEnd    bool
}

func Constructor() WordDictionary {
	return WordDictionary{map[byte]*WordDictionary{}, false}
}

func (this *WordDictionary) AddWord(word string) {
	curr := this

	for i := 0; i < len(word); i++ {
		ch := word[i]
		if _, ok := curr.children[ch]; !ok {
			curr.children[ch] = &WordDictionary{map[byte]*WordDictionary{}, false}
		}

		curr = curr.children[ch]
	}

	curr.isEnd = true
}

func (this *WordDictionary) Search(word string) bool {
	return dfs(0, this, word)
}

func dfs(j int, root *WordDictionary, word string) bool {
	curr := root

	for i := j; i < len(word); i++ {
		ch := word[i]
		if ch == '.' {
			for _, child := range curr.children {
                if ok := dfs(i+1, child, word); ok {
                    return true
                }
			}

            return false
		} else {
			if _, ok := curr.children[ch]; !ok {
				return false
			}
			curr = curr.children[ch]
		}
	}

    return curr.isEnd
}

func TestLogic(t *testing.T) {
	wordDictionary := Constructor()
	wordDictionary.AddWord("bad")
	wordDictionary.AddWord("dad")
	wordDictionary.AddWord("mad")
	res := wordDictionary.Search("pad") // return False
	if res {
		t.Error()
		return
	}
	res = wordDictionary.Search("bad") // return True

	if !res {
		t.Error()
		return
	}
	res = wordDictionary.Search(".ad") // return True

	if !res {
		t.Error()
		return
	}
	res = wordDictionary.Search("b..") // return True

	if !res {
		t.Error()
		return
	}
}
