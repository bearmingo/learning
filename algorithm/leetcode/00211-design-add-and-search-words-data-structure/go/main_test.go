package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestFunc(t *testing.T) {
	a := assert.New(t)

	wd := Constructor()
	wd.AddWord("bad")
	wd.AddWord("dad")
	wd.AddWord("mad")
	a.False(wd.Search("pad"))
	a.True(wd.Search("bad"))
	a.True(wd.Search(".ad"))
	a.True(wd.Search("b.."))
}

func TestFunc2(t *testing.T) {
	a := assert.New(t)

	wd := Constructor()
	wd.AddWord("bad")
	wd.AddWord("bada")
	a.True(wd.Search("bad"))
	a.True(wd.Search("bada"))
}

func TestFunc5(t *testing.T) {
	a := assert.New(t)

	wd := Constructor()
	wd.AddWord("bada")
	wd.AddWord("bad")
	a.True(wd.Search("bad"))
	a.True(wd.Search("bada"))
}

func TestFunc4(t *testing.T) {
	a := assert.New(t)

	wd := Constructor()
	wd.AddWord("bada")
	wd.AddWord("badcc")
	a.True(wd.Search("bada"))
	a.True(wd.Search("badcc"))
	a.True(wd.Search("bad.c"))
}
