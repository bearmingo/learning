#!/usr/bin/env python
# -*- coding=utf-8 -*-

"""
LeetCode:
https://leetcode.com/problems/jump-game

Runtime: 44 ms, faster than 92.15% of Python3 online submissions for Jump Game.
Memory Usage: 14.5 MB, less than 5.28% of Python3 online submissions for Jump Game.
"""

from typing import List

class Solution:
    def canJump(self, nums: List[int]) -> bool:
        # Special cases
        if len(nums) < 2:
            return True

        n = len(nums)

        last = n-1
        for i in range(n-2, -1, -1):
            if i + nums[i] >= last:
                last = i
        return last == 0

if __name__ == "__main__":
    s = Solution()
    assert s.canJump([2,3,1,1,4]) == True