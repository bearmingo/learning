package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestFunc(t *testing.T) {
	a := assert.New(t)
	a.Equal(3, lengthOfLongestSubstring("abcabcbb"))
	a.Equal(1, lengthOfLongestSubstring("bbbbb"))
	a.Equal(3, lengthOfLongestSubstring("pwwkew"))
	a.Equal(0, lengthOfLongestSubstring(""))
	a.Equal(1, lengthOfLongestSubstring(" "))
	a.Equal(1, lengthOfLongestSubstring("  "))
	a.Equal(5, lengthOfLongestSubstring("nfpdmpi"))
	a.Equal(5, lengthOfLongestSubstring("tmmzuxt"))
	a.Equal(3, lengthOfLongestSubstring("abcabcbb"))
}
