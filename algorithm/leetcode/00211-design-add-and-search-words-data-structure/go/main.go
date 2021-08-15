package main

import (
	"strings"
)

type (
	kind uint8
	node struct {
		label          byte
		prefix         string
		parent         *node
		staticChildren []*node
		isLeaf         bool
	}
)

type WordDictionary struct {
	root *node
}

func (n *node) addStaticNode(str string, isLeaf bool) *node {
	c := &node{
		label:  str[0],
		prefix: str,
		parent: n,
		isLeaf: isLeaf,
	}
	if n.staticChildren == nil {
		n.staticChildren = make([]*node, 0)
	}
	n.staticChildren = append(n.staticChildren, c)
	return c
}

func (n *node) findStaticNodeWithLabel(label byte) *node {
	if n.staticChildren == nil {
		return nil
	}
	for _, c := range n.staticChildren {
		if c.label == label {
			return c
		}
	}
	return nil
}

func (n *node) addNode(str string) {
	parent := n
	for {
		idx := strings.IndexRune(str, '.')
		if idx == -1 {
			parent.addStaticNode(str, true)
			break
		}

		if idx != 0 {
			prefix := str[:idx]
			parent = parent.addStaticNode(prefix, false)
		}
	}
}

func (n *node) findNode(str string) bool {

	idx := 0
	minLen := min(len(n.prefix), len(str))
	for ; idx < minLen && (str[idx] == '.' || str[idx] == n.prefix[idx]); idx++ {
	}

	if idx == len(n.prefix) {
		if idx == len(str) {
			return n.isLeaf
		}

		if label := str[idx]; label != '.' {
			n = n.findStaticNodeWithLabel(str[idx])
			if n == nil {
				return false
			}
			return n.findNode(str[idx:])
		}

		for _, cn := range n.staticChildren {
			if cn.findNode(str[idx:]) {
				return true
			}
		}
	}

	return false
}

/** Initialize your data structure here. */
func Constructor() WordDictionary {
	wd := WordDictionary{}
	wd.root = &node{prefix: ""}
	return wd
}

func (this *WordDictionary) AddWord(word string) {
	if len(word) == 0 {
		return
	}

	n := this.root
	str := word

	for {
		idx := 0
		minLen := min(len(n.prefix), len(str))
		for ; idx < minLen && (n.prefix[idx] == str[idx]); idx++ {
		}

		if idx == len(n.prefix) {
			if idx == len(str) {
				n.isLeaf = true
				break
			}

			str = str[idx:]
			c := n.findStaticNodeWithLabel(str[0])
			if c == nil {
				n.addStaticNode(str, true)
				break
			}

			n = c
			continue
		}

		// separate the node
		newPrefix := n.prefix[:idx]
		newNode := &node{
			label:  newPrefix[0],
			prefix: newPrefix,
			parent: n.parent,
		}
		newNode.staticChildren = make([]*node, 0)
		newNode.staticChildren = append(newNode.staticChildren, n)
		for i, pc := range newNode.parent.staticChildren {
			if pc != n {
				continue
			}
			newNode.parent.staticChildren[i] = newNode
			break
		}
		n.parent = newNode
		n.label = n.prefix[idx]
		n.prefix = n.prefix[idx:]

		if len(str) == idx {
			newNode.isLeaf = true
		} else {
			newNode.addStaticNode(str[idx:], true)
		}
		break
	}
}

func (this *WordDictionary) Search(word string) bool {
	return this.root.findNode(word)
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

/**
 * Your WordDictionary object will be instantiated and called as such:
 * obj := Constructor();
 * obj.AddWord(word);
 * param_2 := obj.Search(word);
 */
