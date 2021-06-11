package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestFunc(t *testing.T) {
	a := assert.New(t)

	//a.Equal(3, singleNumber([]int{2, 2, 3, 2}))
	//a.Equal(99, singleNumber([]int{0, 1, 0, 1, 0, 1, 99}))
	a.Equal(-4, singleNumber([]int{-2, -2, 1, 1, 4, 1, 4, 4, -4, -2}))
}
