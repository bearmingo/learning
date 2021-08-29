package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestFunc(t *testing.T) {
	a := assert.New(t)

	// [[],[5],[7],[5],[7],[4],[5],[],[],[],[]]
	// 输出：[null,null,null,null,null,null,null,5,7,5,4]

	obj := Constructor()
	obj.Push(5)
	obj.Push(7)
	obj.Push(5)
	obj.Push(7)
	obj.Push(4)
	obj.Push(5)

	a.Equal(obj.Pop(), 5)
	a.Equal(obj.Pop(), 7)
	a.Equal(obj.Pop(), 5)
	a.Equal(obj.Pop(), 4)
}
