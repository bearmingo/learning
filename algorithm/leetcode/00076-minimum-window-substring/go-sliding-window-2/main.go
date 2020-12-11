package main

// https://leetcode-cn.com/problems/minimum-window-substring/

func minWindow(s string, t string) string {
	sLen := len(s)
	tLen := len(t)
	// left, right 记录窗口移动的位置
	// distance 满足需求的字符的个数
	// begin 记录查找到最优解时，字符串起始位置，
	// minLen 最优解的长度
	left, right, distance, begin, minLen := 0, 0, 0, 0, sLen+1

	// tLen == distance时，表示窗口中的字符串是满足条件(但不一定最优解)

	var (
		tFreq   [128]int // 记录目标字符个数统计
		winFreq [128]int // 窗口中字符的个数统计
	)

	if sLen == 0 || tLen == 0 || sLen < tLen {
		return ""
	}

	// 统计目标字符个数
	for _, v := range t {
		tFreq[v]++
	}

	for right < sLen {
		rightChar := s[right]

		// 不在目标中的字符串，直接移动right
		if tFreq[rightChar] == 0 {
			right++
			continue
		}

		// 是目标中的字符，且没有超过需要的个数时，
		if winFreq[rightChar] < tFreq[rightChar] {
			distance++
		}
		winFreq[rightChar]++
		right++

		for distance == tLen {
			if right-left < minLen {
				minLen = right - left
				begin = left
			}

			leftChar := s[left]
			if tFreq[leftChar] == 0 {
				left++
				continue
			}

			// 当前窗口中的leftChar，已经是最少需要的个数来。再少一个就不满足条件来，
			// 这时distance需要减下
			if winFreq[leftChar] == tFreq[leftChar] {
				distance--
			}
			winFreq[leftChar]--
			left++
		}
	}

	// 没找到
	if minLen == sLen+1 {
		return ""
	}

	return s[begin : begin+minLen]
}
