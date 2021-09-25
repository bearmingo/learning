package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestFunc(t *testing.T) {
	a := assert.New(t)

	a.False(equationsPossible([]string{"a==b", "b!=a"}))
	a.True(equationsPossible([]string{"b==a", "a==b"}))
	a.True(equationsPossible([]string{"a==b", "b==c", "a==c"}))
	a.False(equationsPossible([]string{"a==b", "b!=c", "c==a"}))
	a.True(equationsPossible([]string{"c==c", "b==d", "x!=z"}))
}
