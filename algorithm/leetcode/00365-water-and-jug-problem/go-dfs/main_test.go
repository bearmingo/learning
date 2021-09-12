package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestFunc(t *testing.T) {
	a := assert.New(t)

	// a.True(canMeasureWater(5, 3, 4))
	a.False(canMeasureWater(2, 6, 5))
}
