#!/usr/bin/env python
# -*- coding=utf-8 -*-

"""
LeetCode:

https://leetcode.com/problems/jump-game

超时
"""

from typing import List

class Solution:
    def canJump(self, nums: List[int]) -> bool:
        if len(nums) <= 1:
            return True

        memo = [None for _ in range(len(nums))]
        memo[-1] = True

        for i in range(len(nums)-2, -1, -1):
            far = min(len(nums)- 1, nums[i] + i)
            for j in range(i+1, far+1):
                if memo[j] == True:
                    memo[i] = True
                    break
            
        return memo[0] is True

if __name__ == "__main__":
    s = Solution()
    assert s.canJump([2,3,1,1,4]) == True