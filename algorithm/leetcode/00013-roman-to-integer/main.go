package main

import "fmt"

func romanToInt(s string) int {
    ret := 0
    
    for i := 0; i < len(s); i++ {
        ch := s[i]
        var chn byte
        if i + 1 < len(s) {
            chn = s[i + 1]
        }
        switch {
        case ch == 'M':
			ret += 1000
		case ch == 'D':
			ret += 500
        case ch == 'C' && chn == 'M':
            ret += 900
            i++
        case ch == 'C' && chn == 'D':
            ret += 400
            i++
        case ch == 'C':
            ret += 100
		case ch == 'X' && chn == 'C':
			ret += 90
			i++
		case ch == 'X' && chn == 'L':
			ret += 40
			i++
		case ch == 'L':
			ret += 50
		case ch == 'X':
			ret += 10
		case ch == 'I' && chn == 'X':
			ret += 9
			i++
		case ch == 'I' && chn == 'V':
			ret += 4
			i++
		case ch == 'V':
			ret += 5
		case ch == 'I':
			ret += 1
		}
	}
	
	return ret
}

func main() {
	fmt.Print(romanToInt("MCMXCIV"))
}