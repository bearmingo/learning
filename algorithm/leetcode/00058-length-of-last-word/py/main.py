#!/usr/bin/python3
# -*- coding=utf-8 -*-

class Solution:
    def lengthOfLastWord(self, s: str) -> int:
        result = s.rsplit(None, 1) # Just need one (1) word split on spaces (None)
        return len(result[-1]) if result else 0

if __name__ == "__main__":
    print(Solution().lengthOfLastWord("hello world"))
    print(Solution().lengthOfLastWord("a "))