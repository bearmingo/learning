package main

func groupAnagrams(strs []string) [][]string {
	assis := make(map[[26]byte]int)
	result := make([][]string, 0)

	for _, str := range strs {
		temp := [26]byte{}
		for _, ch := range str {
			temp[ch-'a']++
		}
		if index, ok := assis[temp]; ok {
			result[index] = append(result[index], str)
		} else {
			result = append(result, []string{str})
			assis[temp] = len(result) - 1
		}
	}

	return result
}

func main() {
	print(groupAnagrams([]string{"eat", "tea", "tan", "ate", "nat", "bat"}))
}
