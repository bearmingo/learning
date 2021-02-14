package main

func index(ch byte) int {
	return int(ch - 'a')
}

func findAnagrams(s string, p string) []int {
	record := make([]int, 26)

	chs := make([]int, 26)
	for i := 0; i < len(p); i++ {
		chs[index(p[i])]++
	}

	isSubString := func() bool {
		for i, n := range record {
			if chs[i] != n {
				return false
			}
		}
		return true
	}

	subStringBegins := make([]int, 0)

	var left, right int
	for right = 0; right < len(s); right++ {
		idx := s[right] - 'a'
		record[idx]++
		for left+len(p)-1 <= right {
			if isSubString() {
				subStringBegins = append(subStringBegins, left)
			}
			lidx := s[left] - 'a'
			record[lidx]--
			left++
		}
	}

	return subStringBegins
}
