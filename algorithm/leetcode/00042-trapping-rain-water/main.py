#!/usr/bin/env python
# -*- coding=utf-8 -*-

"""
Leetcode:
https://leetcode.com/problems/trapping-rain-water/

Runtime: 40 ms, faster than 99.98% of Python3 online submissions for Trapping Rain Water.
Memory Usage: 13.3 MB, less than 5.11% of Python3 online submissions for Trapping Rain Water.
"""

from typing import List

class Solution:
    def trap(self, height: List[int]) -> int:
        left, right = 0, len(height) - 1

        total = 0
        max_left = 0
        max_right = 0
        while (left <= right):
            hl, hr = height[left], height[right]
            if hl < hr:
                if max_left < hl:
                    max_left = hl
                else:
                    total += max_left - hl
                left += 1
            else:
                if max_right < hr:
                    max_right = hr
                else:
                    total += max_right - hr
                right -= 1

        return total


if __name__ == "__main__":
    print(Solution().trap([0,1,0,2,1,0,1,3,2,1,2,1]))