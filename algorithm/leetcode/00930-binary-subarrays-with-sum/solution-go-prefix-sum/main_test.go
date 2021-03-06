package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestFunc(t *testing.T) {
	a := assert.New(t)

	a.Equal(4, numSubarraysWithSum([]int{1, 0, 1, 0, 1}, 2))
	a.Equal(15, numSubarraysWithSum([]int{0, 0, 0, 0, 0}, 0))
}
