package solutiongodp

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestFunc(t *testing.T) {
	a := assert.New(t)

	a.Equal(5, maxTurbulenceSize([]int{9, 4, 2, 10, 7, 8, 8, 1, 9}))
	a.Equal(2, maxTurbulenceSize([]int{4, 8, 12, 16}))
	a.Equal(1, maxTurbulenceSize([]int{1}))
}
