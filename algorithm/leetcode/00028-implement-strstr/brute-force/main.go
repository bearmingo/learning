package main

import "fmt"

func strStr(haystack string, needle string) int {
	if len(needle) == 0 {
		return 0
	}
	if len(needle) > len(haystack) {
		return -1
	}

	for i := 0; i < len(haystack); i++ {
		j := 0
		for ; j < len(needle); j++ {
			if i+j >= len(haystack) {
				return -1
			}
			if haystack[i+j] != needle[j] {
				break
			}
		}
		if j == len(needle) {
			return i
		}
	}

	return -1
}

func main() {
	fmt.Println(strStr("a", "a"))
}
