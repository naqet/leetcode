package main

import (
	"testing"
)

type WordDictionary struct {
	children [26]*WordDictionary
	isEnd    bool
}

func Constructor() WordDictionary {
	return WordDictionary{}
}

func (this *WordDictionary) AddWord(word string) {
	curr := this

	for i := 0; i < len(word); i++ {
		ch := word[i]
		if curr.children[ch-'a'] == nil {
			curr.children[ch-'a'] = &WordDictionary{}
		}

		curr = curr.children[ch-'a']
	}

	curr.isEnd = true
}

func (this *WordDictionary) Search(word string) bool {
	return dfs(0, this, word)
}

func dfs(j int, root *WordDictionary, word string) bool {
    if root == nil {
        return false
    }

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
			if curr.children[ch-'a'] == nil {
				return false
			}
			curr = curr.children[ch-'a']
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
