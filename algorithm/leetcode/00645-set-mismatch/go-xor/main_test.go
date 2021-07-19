package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestFunc(t *testing.T) {
	a := assert.New(t)

	a.Equal([]int{2, 3}, findErrorNums([]int{1, 2, 2, 4}))

}
