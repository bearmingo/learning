#!/usr/bin/env python
# -*- coding=utf-8 -*-

"""
Runtime: 60ms, faster than: 95.49%
Memory usage: 13.2MB, less than 5.14%
"""

class Solution:
    def combinationSum(self, candidates: List[int], target: int) -> List[List[int]]:
        results = []
        
        candidates.sort()
        
        def sumNext(total, nums, index):
            for i in range(index, len(candidates)):
                c = candidates[i]
                tmp = total + c
                if tmp < target:
                    nums.append(c)
                    sumNext(tmp, nums, i)
                    nums.pop()
                elif tmp == target:
                    results.append(nums + [c])
                    break
                elif tmp > target:
                    break
        sumNext(0, [], 0)
        return results

