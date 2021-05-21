package solutiongodp

func max(a, b int) int {
	if a > b {
		return a
	} else {
		return b
	}
}

func maxTurbulenceSize(arr []int) int {
	if len(arr) < 2 {
		return 1
	}

	incr, desc := 1, 1

	ret := 1

	for i := 1; i < len(arr); i++ {
		if arr[i-1] < arr[i] {
			incr = desc + 1
			desc = 1
		} else if arr[i-1] > arr[i] {
			desc = incr + 1
			incr = 1
		} else {
			desc = 1
			incr = 1
		}
		ret = max(ret, max(incr, desc))
	}

	return ret
}
