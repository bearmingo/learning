package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestFunc(t *testing.T) {
	a := assert.New(t)

	a.True(hasGroupsSizeX([]int{1, 2, 3, 4, 4, 3, 2, 1}))

}
