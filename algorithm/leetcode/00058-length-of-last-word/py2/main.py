#!/usr/bin/python3
# -*- coding=utf-8 -*-

class Solution:
    def lengthOfLastWord(self, s: str) -> int:
        if len(s) == 0:
            return 0
    
        num = 0
        for _, ch in enumerate(reversed(s)):
            if ch == ' ':
                if num == 0:
                    continue
                else:
                    break
            num += 1
        return num


if __name__ == "__main__":
    print(Solution().lengthOfLastWord("hello world"))
    print(Solution().lengthOfLastWord("a "))