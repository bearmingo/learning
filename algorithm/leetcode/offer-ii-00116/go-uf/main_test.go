package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestFunc(t *testing.T) {
	a := assert.New(t)

	a.Equal(2, findCircleNum([][]int{{1, 1, 0}, {1, 1, 0}, {0, 0, 1}}))
	a.Equal(3, findCircleNum([][]int{{1, 0, 0}, {0, 1, 0}, {0, 0, 1}}))

}
