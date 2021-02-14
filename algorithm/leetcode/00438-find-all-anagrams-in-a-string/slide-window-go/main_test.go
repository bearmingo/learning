package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestFunc(t *testing.T) {
	a := assert.New(t)

	a.Equal([]int{0, 6}, findAnagrams("cbaebabacd", "abc"))
	a.Equal([]int{0, 1, 2}, findAnagrams("abab", "ab"))
}
