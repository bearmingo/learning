package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestFunc(t *testing.T) {
	a := assert.New(t)

	a.Equal(7, subarraysWithKDistinct([]int{1, 2, 1, 2, 3}, 2))
	a.Equal(3, subarraysWithKDistinct([]int{1, 2, 1, 3, 4}, 3))
}
