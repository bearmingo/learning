package main

import "fmt"

func longestPalindrome(s string) string {
	table := make([][]bool, len(s))
	for i := 0; i < len(s); i++ {
		table[i] = make([]bool, len(s))
	}
	//table := [len(s)][len(s)]bool{}

	maxLen := 0
	start := 0

	foundIn := func(left, right int) bool {
		if left == right {
			if maxLen < 1 {
				maxLen = 1
				start = left
			}
			table[left][right] = true
			return true
		} else if left+1 == right && s[left] == s[right] {
			if maxLen < 2 {
				maxLen = 2
				start = left
			}
			table[left][right] = true
			return true
		} else if s[left] == s[right] && table[left+1][right-1] {
			if maxLen < (right - left + 1) {
				maxLen = right - left + 1
				start = left
			}
			table[left][right] = true
			return true
		}
		return false
	}

	for l := 1; l <= len(s); l++ { // 长度
		for j := 0; j < len(s) && j+l-1 < len(s); j++ { // 起点
			foundIn(j, j+l-1)
		}
	}

	return s[start : start+maxLen]
}

func main() {
	//test := "cbbd"
	//test := "babad"
	test := "ccc"
	ret := longestPalindrome(test)
	fmt.Println(ret)
}
