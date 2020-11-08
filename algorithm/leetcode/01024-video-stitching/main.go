package main

import "fmt"

// https://leetcode-cn.com/problems/video-stitching/

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func videoStitching(clips [][]int, T int) int {
	furthest := make([]int, T)

	for i := 0; i < len(clips); i++ {
		for j := clips[i][0]; j <= clips[i][1]; j++ {
			if j >= T {
				break
			}
			furthest[j] = max(furthest[j], clips[i][1])
		}
	}

	var end, last, ans int
	for i := 0; i < T; i++ {
		last = max(last, furthest[i])
		if last == i {
			return -1
		}
		if end == i {
			ans++
			end = last
		}
	}
	return ans
}

func main() {
	//fmt.Println(videoStitching([][]int{{0, 2}, {4, 6}, {8, 10}, {1, 9}, {1, 5}, {5, 9}}, 10))
	fmt.Println(videoStitching([][]int{{0, 1}, {6, 8}, {0, 2}, {5, 6}, {0, 4}, {0, 3}, {6, 7}, {1, 3}, {4, 7}, {1, 4}, {2, 5}, {2, 6}, {3, 4}, {4, 5}, {5, 7}, {6, 9}}, 9))
}
