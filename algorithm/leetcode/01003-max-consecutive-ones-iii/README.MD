# 1004. 最大连续1的个数 III

[1004. 最大连续1的个数 III](https://leetcode-cn.com/problems/max-consecutive-ones-iii/)

解题:

对于数组$A$的区间$[i, j]$，只要他包含不超过$k$个0，就可以认为是一个满足要求，且数组长度是$[j-i+1]$的数组。

因此题目等价与：

找到一个区间$[i, j]$， 且包含的0的个数小于等于$k$。

想要快速判断出一个区间0的个数，可以将数组$A$中的0变成1，1变成0。此时，对数组$A$求出前缀和，记为数组$P$。那么区间$[i, j]$中的0的个数等于:

$$
P[j] - P[i-1]
\tag{1}
$$

当`公式\(1\)1中的结果小于等于$k$，即可找到一个满足要求的数组。找出其中长度最长的，就能解决该问题了。

$$
P[j] - P[i-1] \le k
\tag{2}
$$

## 方法一：滑动窗口

在移动`right`过程中，判断是否满足`公式(2)`。不满足时移动`left`，使得满足`公式(2)`

使用$rsum$和$lsum$，记录right和left分别对应的前缀和。

## REF

https://leetcode-cn.com/problems/max-consecutive-ones-iii/solution/zui-da-lian-xu-1de-ge-shu-iii-by-leetcod-hw12/
