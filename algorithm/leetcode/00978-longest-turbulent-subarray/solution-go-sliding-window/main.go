package main

func max(a, b int) int {
	if a > b {
		return a
	} else {
		return b
	}
}

func maxTurbulenceSize(arr []int) int {
	left, right := 0, 0
	ret := 1

	for right < len(arr)-1 {
		if left == right {
			if arr[left] == arr[left+1] {
				left++
			}
			right++
		} else {
			if arr[right-1] < arr[right] && arr[right] > arr[right+1] {
				right++
			} else if arr[right-1] > arr[right] && arr[right] < arr[right+1] {
				right++
			} else {
				left = right
			}
		}
		ret = max(ret, right-left+1)
	}

	return ret
}
