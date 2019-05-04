#!/usr/bin/env python
# -*- coding=utf-8 -*-

"""
LeetCode:
https://leetcode.com/problems/jump-game-ii

Runtime: 68 ms, faster than 29.35% of Python3 online submissions for Jump Game II.
Memory Usage: 14.5 MB, less than 5.83% of Python3 online submissions for Jump Game II.
"""

from typing import List

class Solution:
    def jump(self, nums: List[int]) -> int:
        """BFS(深度优先算法)
        将搜索方式理解为BFS算法，以[2,3,1,1,4]为例
             0  1  2, 3  4
            [2, 3, 1, 1, 4]
        1层： 从$i=0$元素为2，可到达的位置是$i \in [1,2]$
        2层：$i \in [1,2]$, i=1时可跳3不，到达末尾，
        """
        # Special cases
        if len(nums) < 2:
            return 0

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
                    return level
                i += 1
            currentMax = nextMax
        
        return 0


if __name__ == "__main__":
    s = Solution()
    assert s.jump([2,3,1,1,4]) == 2
    #assert s.jump([2,3,0,1,4]) == 2
    #assert s.jump([5,9,3,2,1,0,2,3,3,1,0,0]) == 3