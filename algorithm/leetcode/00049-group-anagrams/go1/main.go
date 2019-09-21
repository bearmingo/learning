package main

import "sort"

type byteSlice []byte

func (c byteSlice) Len() int {
	return len(c)
}

func (c byteSlice) Swap(i, j int) {
	c[i], c[j] = c[j], c[i]
}

func (c byteSlice) Less(i, j int) bool {
	return c[i] < c[j]
}

func sortedString(s string) string {
	b := byteSlice(s)
	sort.Sort(b)

	return string(b)
}

func groupAnagrams(strs []string) [][]string {
	result := make(map[string][]string)

	for _, str := range strs {
		temp := sortedString(str)

		if item, ok := result[temp]; !ok {
			arr := append([]string{}, str)
			result[temp] = arr
		} else {
			result[temp] = append(item, str)
		}
	}

	ret := make([][]string, 0, len(result))
	for _, item := range result {
		ret = append(ret, item)
	}

	return ret
}

func main() {
	print(groupAnagrams([]string{"eat", "tea", "tan", "ate", "nat", "bat"}))
}
