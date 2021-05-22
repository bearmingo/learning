package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestFunc(t *testing.T) {
	a := assert.New(t)

	a.Equal(6, longestOnes([]int{1, 1, 1, 0, 0, 0, 1, 1, 1, 1, 0}, 2))
	a.Equal(10, longestOnes([]int{0, 0, 1, 1, 0, 0, 1, 1, 1, 0, 1, 1, 0, 0, 0, 1, 1, 1, 1}, 3))
	a.Equal(3, longestOnes([]int{0, 0, 1, 1, 1, 0, 0}, 0))
	a.Equal(0, longestOnes([]int{0, 0, 0, 0}, 0))
}
