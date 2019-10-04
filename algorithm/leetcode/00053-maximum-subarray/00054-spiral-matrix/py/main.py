# -*- coding=utf-8 -*-

from typing import List


class Solution:
    def spiralOrder(self, matrix: List[List[int]]) -> List[int]:
        return matrix and [*matrix.pop(0)] + self.spiralOrder([*zip(*matrix)][::-1])


if __name__ == "__main__":
    s = Solution()
    testcase = [
        [1, 2, 3, 4],
        [5, 6, 7, 8],
        [9,10,11,12]
    ]

    print(s.spiralOrder(testcase))