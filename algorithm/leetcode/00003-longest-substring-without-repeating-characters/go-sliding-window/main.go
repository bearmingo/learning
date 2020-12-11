package main

func lengthOfLongestSubstring(s string) int {
	var l, r, max int

	chs := make([]int, 256)
	for k := 0; k < len(chs); k++ {
		chs[k] = -1
	}

	for r = 0; r < len(s); r++ {
		c := s[r]
		last := chs[c]
		if last != -1 {
			for t := l; t <= last; t++ {
				chs[s[t]] = -1
			}
			l = last + 1

			chs[c] = r
			continue
		}
		chs[c] = r

		temp := r - l + 1
		if temp > max {
			max = temp
		}
	}

	return max
}
