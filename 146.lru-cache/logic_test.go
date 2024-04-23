package main

import (
	"testing"
)

type Node struct {
	val  int
	key  int
	next *Node
	prev *Node
}

type LRUCache struct {
	capacity int
	length   int
	head     *Node
	tail     *Node
	hash     map[int]*Node
}

func Constructor(capacity int) LRUCache {
	return LRUCache{capacity: capacity, hash: map[int]*Node{}}
}

func (this *LRUCache) Get(key int) int {
	node, ok := this.hash[key]

	if !ok {
		return -1
	}

	this.detach(node)
	this.prepend(node)

	return node.val
}

func (this *LRUCache) Put(key int, value int) {
	node, ok := this.hash[key]
	if ok {
		this.detach(node)
		this.prepend(node)
		node.val = value
		return
	}

	newNode := &Node{val: value, key: key}
	this.prepend(newNode)
	this.trimCache()
	this.hash[key] = newNode
}

func (this *LRUCache) prepend(node *Node) {
	this.length++

	if this.head == nil {
		this.head = node
		this.tail = node
		return
	}

	node.next = this.head
	this.head.prev = node
	this.head = node
}

func (this *LRUCache) detach(node *Node) {
	if node == nil {
		return
	}

	if this.tail == node {
		this.tail = node.prev
	}

	if this.head == node {
		this.head = node.next
	}

	if node.prev != nil {
		node.prev.next = node.next
	}

	if node.next != nil {
		node.next.prev = node.prev
	}

	node.prev = nil
	node.next = nil

	this.length--
}

func (this *LRUCache) trimCache() {
	if this.length <= this.capacity {
		return
	}

	tail := this.tail

	this.detach(tail)

	delete(this.hash, tail.key)
}

type singleTest struct {
	input    string
	expected bool
}

var tests = []singleTest{}

func TestLogic(t *testing.T) {
	lru := Constructor(1)
	lru.Put(1, 1)
	lru.Put(2, 2)
	lru.Put(3, 3)

	if val := lru.Get(1); val != -1 {
		t.Error(val)
		return
	}

	if val := lru.Get(2); val != -1 {
		t.Error(val)
		return
	}

	if val := lru.Get(100); val != -1 {
		t.Error(val)
		return
	}
}
