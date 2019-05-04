#!/usr/bin/env python
# -*- coding=utf-8 -*-

"""
LeetCode:
https://leetcode.com/problems/multiply-strings

Runtime: 164 ms, faster than 38.04% of Python3 online submissions for Multiply Strings.
Memory Usage: 13.1 MB, less than 6.04% of Python3 online submissions for Multiply Strings.
"""


class Solution:
    def multiply(self, num1: str, num2: str) -> str:
        res = [0 for _ in range(len(num1) + len(num2))]
        for i, c1 in enumerate(reversed(num1)):
            n1 = ord(c1) - ord('0')
            for j, c2 in enumerate(reversed(num2)):
                n2 = ord(c2) - ord('0')

                index = i + j
                tmp = n1 * n2 + res[index]
                
                res[index] = tmp % 10
                res[index+1] += int(tmp / 10)

        resStr = ""
        start = False
        for x in reversed(res):
            if not start and x != 0:
                start = True
            if start:
                resStr += chr(x + ord('0'))

        return resStr if len(resStr) != 0 else "0" 

if __name__ == "__main__":
    print(Solution().multiply("123", "456"))