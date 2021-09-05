package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestFunc(t *testing.T) {
	a := assert.New(t)

	//a.Equal(2, maxDistance([][]int{{1, 0, 1}, {0, 0, 0}, {1, 0, 1}}))
	//a.Equal(4, maxDistance([][]int{{1, 0, 0}, {0, 0, 0}, {0, 0, 0}}))
	a.Equal(-1, maxDistance([][]int{{0, 0, 0, 0}, {0, 0, 0, 0}, {0, 0, 0, 0}, {0, 0, 0, 0}}))
}
