package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestFunc(t *testing.T) {
	a := assert.New(t)

	a.Equal(3, totalFruit([]int{1, 2, 1}))
	a.Equal(3, totalFruit([]int{0, 1, 2, 2}))
	a.Equal(4, totalFruit([]int{1, 2, 3, 2, 2}))
	a.Equal(5, totalFruit([]int{3, 3, 3, 1, 2, 1, 1, 2, 3, 3, 4}))
}
