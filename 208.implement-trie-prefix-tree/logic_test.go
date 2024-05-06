package main

import (
	"testing"
)

type Trie struct {
	children [26]*Trie
	isEnd    bool
}

func Constructor() Trie {
	return Trie{}
}

func (this *Trie) Insert(word string) {
	curr := this
	for _, ch := range word {
		charIdx := ch - 'a'

		if curr.children[charIdx] == nil {
			curr.children[charIdx] = &Trie{}
		}

		curr = curr.children[charIdx]
	}


	curr.isEnd = true
}

func (this *Trie) Search(word string) bool {
	curr := this
	for _, ch := range word {
		charIdx := ch - 'a'

		if curr.children[charIdx] == nil {
            return false
		}

		curr = curr.children[charIdx]
	}

	return curr.isEnd
}

func (this *Trie) StartsWith(prefix string) bool {
	result := true
	curr := this
	for _, ch := range prefix {
		charIdx := ch - 'a'

		if curr.children[charIdx] == nil {
			result = false
			break
		}

		curr = curr.children[charIdx]
	}

	return result
}

func TestLogic(t *testing.T) {
	trie := Constructor()
	trie.Insert("hello")
    res := trie.Search("hello") 

	if !res {
		t.Error()
		return
	}

	res = trie.Search("helloa")
	if res {
		t.Error()
		return
	}
}
