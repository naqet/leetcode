package main

import (
	"fmt"
	"strconv"
	"strings"

	"reflect"
	"testing"
)

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

type Codec struct {
	i int
}

func Constructor() Codec {
	return Codec{}
}

func (this *Codec) serialize(root *TreeNode) string {
	return dfs(root)
}

func dfs(root *TreeNode) string {
	if root == nil {
		return "n,"
	}

	return fmt.Sprintf("%d,", root.Val) + dfs(root.Left) + dfs(root.Right)
}

func (this *Codec) deserialize(data string) *TreeNode {
	values := strings.Split(data, ",")
	return this.create(values)
}

func (this *Codec) create(values []string) *TreeNode {
	if values[this.i] == "n" {
		this.i++
		return nil
	}

	val, _ := strconv.Atoi(values[this.i])

	node := &TreeNode{val, nil, nil}

	this.i++

	node.Left = this.create(values)
	node.Right = this.create(values)
	return node
}

func walk(root *TreeNode, path *[]int) {
	if root == nil {
		return
	}

	walk(root.Left, path)

	*path = append(*path, root.Val)

	walk(root.Right, path)
}

func TestLogic(t *testing.T) {
	root := &TreeNode{2, &TreeNode{1, nil, nil}, &TreeNode{3, nil, nil}}
	c := Constructor()

	res := c.serialize(root)
	if res != "2,1,n,n,3,n,n," {
		t.Error(res)
		return
	}

	path := []int{}
	node := c.deserialize(res)

	walk(node, &path)
	if !reflect.DeepEqual(path, []int{1, 2, 3}) {
		t.Error(path)
		return
	}
}
