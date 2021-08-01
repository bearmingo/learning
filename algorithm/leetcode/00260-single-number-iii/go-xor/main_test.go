package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestFunc(t *testing.T) {
	a := assert.New(t)

	a.Equal([]int{5, 3}, singleNumber([]int{1, 2, 1, 3, 2, 5}))
	a.Equal([]int{0, -1}, singleNumber([]int{-1, 0}))
	a.Equal([]int{0, 1}, singleNumber([]int{1, 0}))

}
