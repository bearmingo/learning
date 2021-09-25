package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestFunc(t *testing.T) {
	a := assert.New(t)

	a.Equal("bacd", smallestStringWithSwaps("dcab", [][]int{{0, 3}, {1, 2}}))
	a.Equal("abcd", smallestStringWithSwaps("dcab", [][]int{{0, 3}, {1, 2}, {0, 2}}))
	a.Equal("abc", smallestStringWithSwaps("cba", [][]int{{0, 1}, {1, 2}}))
}
