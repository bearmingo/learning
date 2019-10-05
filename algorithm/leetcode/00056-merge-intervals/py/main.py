#!/usr/bin/python3
# -*- coding=utf-8 -*-

from typing import List

class Solution:
    def merge(self, intervals: List[List[int]]) -> List[List[int]]:
        if len(intervals) == 0:
            return []
        
        intervals.sort(key=lambda i: i[0])
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
    print(s.merge([[1,3],[2,6],[8,10],[15,18]]))
    print(s.merge([[1,4],[4,5]]))
    print(s.merge([[1,4],[0,4]]))

