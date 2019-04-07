package main

import "fmt"

// KMP

// 获取跳转表
func getNext(p string) []int {
	next := make([]int, len(p))

	// 第一个匹配失败时，源字符串向后一动一维
	next[0] = -1

	j := 0
	k := -1

	for j < len(p)-1 {
		if k == -1 || p[j] == p[k] {
			k++
			j++
			next[j] = k

		} else {
			k = next[k]
		}
	}
	return next
}

func strStr(haystack string, needle string) int {
	if len(needle) == 0 {
		return 0
	}
	if len(needle) > len(haystack) {
		return -1
	}

	next := getNext(needle)

	i, j := 0, 0
	for i < len(haystack) && j < len(needle) {
		if j == -1 || haystack[i] == needle[j] {
			i++
			j++
		} else {
			j = next[j]
		}
	}

	if j == len(needle) {
		return i - j
	}

	return -1
}

func main() {
	fmt.Println(strStr("hello", "ll"))
}
