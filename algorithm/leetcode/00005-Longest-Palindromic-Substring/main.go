package main

func foundIn(s string, left, right int, start , maxLen *int) {
    for left >= 0 && right < len(s) && s[left] == s[right] {
        left--;
        right++;
    }

    if *maxLen < right - left - 1 {
        *start = left + 1
        *maxLen = right - left - 1
    }
}

func longestPalindrome(s string) string {
    if len(s) < 2 {
        return s
    }

    n := len(s)
    maxLen := 0
    start := 0

    for i := 0; i < n - 1; i++ {
        foundIn(s, i, i, &start, &maxLen)
        foundIn(s, i, i + 1, &start, &maxLen)
    }

    return s[start:start+maxLen]
}
