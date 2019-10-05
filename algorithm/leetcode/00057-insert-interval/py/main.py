#!/usr/bin/python3
# -*- coding=utf-8 -*-

from typing import List

class Solution:
    def insert(self, intervals: List[List[int]], newInterval: List[int]) -> List[List[int]]:
        ret = []
        if len(intervals) == 0:
            ret.append(newInterval)
            return ret

        inserted = False
        for i, item in enumerate(intervals):
            if item[0] > newInterval[0]:
                intervals.insert(i, newInterval)
                inserted = True
                break
        if not inserted:
            intervals.append(newInterval)

        ret = []
        (b, e) = intervals[0]
        for i in intervals[1:]:
            # |---e---| or |------|--e
            if i[0] <= e:
                if e < i[1]:
                    e = i[1]
            else:
                # e |---|
                ret.append([b, e])
                b, e = i
        ret.append([b, e])
        return ret

if __name__ == "__main__":
    s = Solution()
    print(s.insert([[1,3],[6,9]], [2,5]))
    print(s.insert([[1,2],[3,5],[6,7],[8,10],[12,16]], [4,8]))