package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestFunc(t *testing.T) {
	a := assert.New(t)

	// 输入：["RLEIterator","next","next","next","next"], [[[3,8,0,9,2,5]],[2],[1],[1],[2]]
	// 输出：[null,8,8,5,-1]

	obj := Constructor([]int{3, 8, 0, 9, 2, 5})

	a.Equal(obj.Next(2), 8)
	a.Equal(obj.Next(1), 8)
	a.Equal(obj.Next(1), 5)
	a.Equal(obj.Next(2), -1)
}
