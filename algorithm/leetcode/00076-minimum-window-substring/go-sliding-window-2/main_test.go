package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestFunc(t *testing.T) {
	a := assert.New(t)

	//a.Equal("BANC", minWindow("ADOBECODEBANC", "ABC"))
	a.Equal("a", minWindow("a", "a"))
}
