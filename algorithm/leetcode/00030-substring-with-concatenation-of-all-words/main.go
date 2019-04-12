package main

import "fmt"

// 在submit1上进行优化
// 1. 使用begin和end偏移量减少位置计算
//
// Leetcode测试结果为:
//     Runtime: 4ms, faster then: 100%
//     Memory: 3.3M, less then: 75%

func findSubstring(s string, words []string) []int {
	// Special cases
	if len(words) == 0 {
		return []int{}
	}

	// create words map
	wordsMap := make(map[string]int)
	for _, word := range words {
		wordsMap[word]++
	}

	var ret []int
	wordLen := len(words[0])
	wordNum := len(words)
	sLen := wordLen * wordNum //

	for i := 0; i < wordLen; i++ {
		searchedWordsMap := make(map[string]int)
		begin := i
		end := i + wordLen
		for len(s)-begin >= sLen {
			w := s[end-wordLen : end]

			if v, ok := wordsMap[w]; ok {
				// 存在这个单词
				// 如列表中中已经存在足够的单词，再添加一个就超了。
				if c, ok2 := searchedWordsMap[w]; ok2 && c == v {
					// 移动到上次出现这个单词后，从记录中除去上次出现这个单词之前的记录
					for y := begin; y < end; y += wordLen {
						tmpWord := s[y : y+wordLen]
						if tmpWord != w {
							searchedWordsMap[tmpWord]--
						} else {
							begin = y + wordLen
							break
						}
					}
				} else {
					searchedWordsMap[w] += 1
				}

			} else {
				// 这个单词不存在，移动到这个单词后开始搜索，并清空搜索记录
				if len(searchedWordsMap) != 0 {
					searchedWordsMap = make(map[string]int)
				}
				begin = end
				end = begin + wordLen
				continue
			}

			// 找到一个
			if end-begin == sLen {
				ret = append(ret, begin)
				searchedWordsMap[s[begin:begin+wordLen]]--
				begin += wordLen
				end += wordLen
			} else {
				end += wordLen
			}
		}
	}

	return ret
}

func main() {
	//fmt.Println(findSubstring("wordgoodgoodgoodbestword", []string{"word", "good", "best", "word"}))
	fmt.Println(findSubstring("barfoothefoobarman", []string{"foo", "bar"}))
}
