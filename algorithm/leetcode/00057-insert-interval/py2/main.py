#!/usr/bin/python3
# -*- coding=utf-8 -*-

from typing import List

class Solution:
    def insert(self, intervals: List[List[int]], newInterval: List[int]) -> List[List[int]]:
        ret = []
        l = len(intervals)
        index = 0
        # add all intervals ending before newInterval starts
        while index < l and intervals[index][1] < newInterval[0]:
            pass
        ret.extend(intervals[:index])
        
        # merge all overlapping intervals to one considering newInterval
        while index < l and intervals[index][0] <= newInterval[1]:
            newInterval = [
                min(intervals[index][0], newInterval[0]),
                max(intervals[index][1], newInterval[1])]
            index += 1

        # add the union of intervals we got
        ret.append(newInterval)
        # add all the rest
        ret.extend(intervals[index:])

        return ret

if __name__ == "__main__":
    s = Solution()
    print(s.insert([[1,3],[6,9]], [2,5]))
    print(s.insert([[1,2],[3,5],[6,7],[8,10],[12,16]], [4,8]))