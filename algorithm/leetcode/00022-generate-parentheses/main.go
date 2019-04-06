package main

import "fmt"

/*
func addNextParenthesis(results *[]string, str string, m, n int) {
	if n == 0 && m == 0 {
		*results = append(*results, str)
		return
	}

	if n > 0 {
		addNextParenthesis(results, str+"(", m+1, n-1)
	}

	if m > 0 {
		addNextParenthesis(results, str+")", m-1, n)
	}
}

func generateParenthesis(n int) []string {
	results := make([]string, 0)
	addNextParenthesis(&results, "", 0, n)

	return results
}
*/
func generateParenthesis(n int) []string {
	var res []string
	var generate func(string, int, int)
	generate = func(s string, l, r int) {
		if len(s) == 2*n {
			res = append(res, s)
			return
		}
		if l < n {
			generate(s+"(", l+1, r)
		}
		if r < l {
			generate(s+")", l, r+1)
		}
	}
	generate("", 0, 0)
	return res
}

func main() {
	ret := generateParenthesis(3)
	for _, item := range ret {
		fmt.Println(item)
	}
}
