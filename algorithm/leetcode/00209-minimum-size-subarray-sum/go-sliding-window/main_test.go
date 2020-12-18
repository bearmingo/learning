package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestFunc(t *testing.T) {
	a := assert.New(t)

	a.Equal(2, minSubArrayLen(7, []int{2, 3, 1, 2, 4, 3}))
	a.Equal(5, minSubArrayLen(15, []int{1, 2, 3, 4, 5}))
	a.Equal(1, minSubArrayLen(4, []int{1, 4, 4}))
}
