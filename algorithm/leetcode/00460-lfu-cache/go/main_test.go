package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestFunc(t *testing.T) {
	a := assert.New(t)

	//["LFUCache", "put", "put", "get", "put", "get", "get", "put", "get", "get", "get"]
	//[[2], [1, 1], [2, 2], [1], [3, 3], [2], [3], [4, 4], [1], [3], [4]]

	l := Constructor(2)
	l.Put(1, 1)
	l.Put(2, 2)
	a.Equal(l.Get(1), 1)
	l.Put(3, 3)
	a.Equal(l.Get(2), -1)
	a.Equal(l.Get(3), 3)
	l.Put(4, 4)
	a.Equal(l.Get(1), -1)
	a.Equal(l.Get(3), 3)
	a.Equal(l.Get(4), 4)
}

func TestFunc1(t *testing.T) {
	a := assert.New(t)

	l := Constructor(2)
	a.Equal(l.Get(2), -1)
	l.Put(2, 6)
	a.Equal(l.Get(1), -1)
	l.Put(1, 5)
	l.Put(1, 2)
	a.Equal(l.Get(1), 2)
	a.Equal(l.Get(2), 6)
}

func TestFunc2(t *testing.T) {
	// ["LFUCache","put","get"]
	// [[0],[0,0],[0]]
	a := assert.New(t)
	l := Constructor(0)
	l.Put(0, 0)
	a.Equal(l.Get(0), -1)
}
