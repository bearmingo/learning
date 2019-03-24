package main

import (
	"fmt"
)

var digitToChs = [][]string {
	[]string{}, 		// 0
	[]string{},			// 1
	[]string{"a", "b", "c"}, // 2
	[]string{"d", "e", "f"},
	[]string{"g", "h", "i"},
	[]string{"j", "k", "l"},
	[]string{"m", "n", "o"},
	[]string{"p", "q", "r", "s"},
	[]string{"t", "u", "v"},
	[]string{"w", "x", "y", "z"},
}

func letterCombinations(digits string) []string {
	if len(digits) == 1 {
		return digitToChs[digits[0] - '0']
	}

    ret := make([]string, 0)
    
	subs := letterCombinations(digits[1:])
	for _, ch := range digitToChs[digits[0] - '0'] {
		for _, sub := range subs {
			ret = append(ret, ch + sub)
		}
	}
	
	return ret
}

func main() {
	fmt.Println(letterCombinations("23"))
}