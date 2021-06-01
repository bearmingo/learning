package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestFunc(t *testing.T) {
	a := assert.New(t)

	a.Equal(1, singleNumber([]int{2, 2, 1}))
	a.Equal(4, singleNumber([]int{4, 1, 2, 1, 2}))
}
