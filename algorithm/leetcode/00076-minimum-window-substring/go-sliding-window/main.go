package main

// https://leetcode-cn.com/problems/minimum-window-substring/

func minWindow(s string, t string) string {
	dict := make(map[byte]int)
	for i := 0; i < len(t); i++ {
		dict[t[i]]++
	}

	target := make(map[byte]int)

	test := func() bool {
		if len(target) < len(dict) {
			return false
		}
		for key, count := range dict {
			if tc, ok := target[key]; ok {
				if tc < count {
					return false
				}
			} else {
				return false
			}
		}
		return true
	}

	var l, r int
	result := s + "end"
	for r = 0; r < len(s); r++ {
		c := s[r]
		if _, ok := dict[c]; ok {
			target[c]++
		}
		for l <= r && test() {
			if r-l+1 < len(result) {
				result = s[l : r+1]
			}
			lc := s[l]
			if _, ok := dict[lc]; ok {
				target[lc]--
			}
			l++
		}
	}

	if len(result) > len(s) {
		return ""
	}
	return result
}
