#!/usr/bin/env python
# -*- coding=utf-8 -*-

"""
LeetCode:
https://leetcode.com/problems/jump-game

Runtime: 56 ms, faster than 35.30% of Python3 online submissions for Jump Game.
Memory Usage: 14.6 MB, less than 5.28% of Python3 online submissions for Jump Game.
"""

from typing import List

class Solution:
    def canJump(self, nums: List[int]) -> bool:
        # Special cases
        if len(nums) < 2:
            return True

        level = 0
        currentMax = 0  # 当前层最大的偏移量
        i = 0 # 当前层中，最小的偏移量

        # 判断下一层中，可用的数据是否为空
        while currentMax - i + 1 > 0:
            level += 1

            nextMax = 0
            # 搜索当前层中，最大可跳转到的位置
            while i <= currentMax:
                nextMax = max(nextMax, nums[i]+i)
                # Reach the ends, we found it
                if nextMax >= len(nums) - 1:
                    return True
                i += 1
            currentMax = nextMax
        
        return False

if __name__ == "__main__":
    s = Solution()
    assert s.canJump([2,3,1,1,4]) == True