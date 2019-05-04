#!/usr/bin/env python
# -*- coding=utf-8 -*-

"""
Leetcode:
https://leetcode.com/problems/wildcard-matching

Runtime: 64 ms, faster than 95.10% of Python3 online submissions for Wildcard Matching.
Memory Usage: 13.3 MB, less than 66.92% of Python3 online submissions for Wildcard Matching.
"""

class Solution:
    def isMatch(self, s: str, p: str) -> bool:
        si = 0
        pi = 0
        starIndex = -1
        match = 0

        while si < len(s):
            # si,pi都前进，如果相等，或者模板字符是?
            if pi < len(p) and (p[pi] == '?' or p[pi] == s[si]):
                si, pi = si+1, pi+1
                continue
            elif pi < len(p) and p[pi] == '*':
                starIndex = pi
                pi = pi + 1
                match = si
                continue
            elif starIndex != -1:
                pi = starIndex + 1
                match = match + 1
                si = match
                continue
            else:
                return False
        
        while pi < len(p) and p[pi] == '*':
            pi += 1
        
        return pi == len(p)


if __name__ == "__main__":
    assert Solution().isMatch("aa", "a") == False
    assert Solution().isMatch("aa", "*") == True
    assert Solution().isMatch("cb", "?a") == False
    assert Solution().isMatch("adceb", "*a*b") == True
    assert Solution().isMatch("acdcb", "a*c?b") == False
    assert Solution().isMatch("mississippi", "m??*ss*?i*pi") == False
    assert Solution().isMatch("ho", "ho**")

    assert Solution().isMatch("aaabbbaabaaaaababaabaaabbabbbbbbbbaabababbabbbaaaaba", "a*******b") == False