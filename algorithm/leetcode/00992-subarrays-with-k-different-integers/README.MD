# K 个不同整数的子数组

## 问题解析

关键：

1. 「最多存在$K$个不同整数的子区间的个数」与「恰好存在$K$个不同整数的子区间的个数」的差恰好等于「最多存在$K−1$个不同整数的子区间的个数」
2. 求最多n个不同整数的子数组时，每次增加合法区间的$left-right$个，该过程正如我们在纸上手算的过程，总过程类似等差数列：$1+2+……+n$ ，即每次新增的区间点，除了增加自己以外，还要和前面的区间点组成子数组，所以为$n-1+1$，即每次加$n$。

## REF

https://leetcode-cn.com/problems/subarrays-with-k-different-integers/solution/k-ge-bu-tong-zheng-shu-de-zi-shu-zu-by-l-ud34/
