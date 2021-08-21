package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestFunc(t *testing.T) {
	a := assert.New(t)

	obj := Constructor()
	obj.Push(1)
	obj.Push(2)
	a.Equal(obj.Pop(), 1)
	a.Equal(obj.Peek(), 2)
	a.False(obj.Empty())

	a.Equal(obj.Pop(), 2)
	a.True(obj.Empty())
}
