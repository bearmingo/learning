package main

func hasGroupsSizeX(deck []int) bool {
	count := make(map[int]int)
	for _, d := range deck {
		count[d]++
	}

	g := -1
	for _, v := range count {
		if v == 0 {
			continue
		}
		if g == -1 {
			g = v
		} else {
			g = gcd(g, v)
		}
	}

	return g >= 2
}

func gcd(x, y int) int {
	if x == 0 {
		return y
	}
	return gcd(y%x, x)
}
