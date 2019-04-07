package main

import (
	"bytes"
	"fmt"
)

func min(l, r int) int {
	if l > r {
		return r
	} else {
		return l
	}
}

func manacherSearch(s string) string {
	buf := bytes.Buffer{}
	buf.WriteString("$#")
	for _, ch := range s {
		buf.WriteRune(ch)
		buf.WriteByte('#')
	}

	s2 := buf.String()

	p := make([]int, len(s2)) // 保存回文半径的数组
	id := 0                   // id为能延伸到最右端的位置的那个回文子串的中心点位置
	mx := 0                   // mx是回文串能够延伸到最右端的位置
	resCenter := 0            // 搜索到最长的回文的中心点
	resLen := 0               // 搜索到最长回文的半径

	for i := 1; i < len(s2); i++ {
		if mx > i {
			p[i] = min(p[2*id-i], mx-i)
		} else {
			p[i] = 1
		}

		// 匹配超出部分的回文
		for l, r := i-p[i], i+p[i]; l > 0 && r < len(s2) && s2[l] == s2[r]; l, r = l-1, r+1 {
			p[i]++
		}

		// 更新延伸到最右边的回文中心点和右边的位置
		if mx < i+p[i] {
			mx = i + p[i]
			id = i
		}

		// 记录最长的回文
		if p[i] > resLen {
			resCenter = i
			resLen = p[i]
		}
	}

	b := (resCenter - resLen) / 2
	return s[b : b+resLen-1]
}

func longestPalindrome(s string) string {
	return manacherSearch(s)
}

func main() {
	//test := "ccc"
	test := "babad"
	//test := "cbbd"
	ret := longestPalindrome(test)
	fmt.Println(ret)
}
