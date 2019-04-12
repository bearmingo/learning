package main

import "fmt"

// Ref:
// https://leetcode.windliang.cc/leetCode-30-Substring-with-Concatenation-of-All-Words.html?q=

func findSubstring(s string, words []string) []int {
	// Special cases
	if len(words) == 0 {
		return []int{}
	}

	// create words map
	wordsMap := make(map[string]int)
	for _, word := range words {
		if v, ok := wordsMap[word]; ok {
			wordsMap[word] = v + 1
		} else {
			wordsMap[word] = 1
		}
	}

	var ret []int
	wordLen := len(words[0])

	for i := 0; i < wordLen; i++ {
		searchedWordsMap := make(map[string]int)
		foundWordsNum := 0
		for j := i; j+wordLen*len(words)-1 < len(s); j += wordLen {
			hasRemoved := false
			for foundWordsNum < len(words) {
				w := s[j+foundWordsNum*wordLen : j+(foundWordsNum+1)*wordLen]
				if _, ok := wordsMap[w]; ok {
					if _, ok2 := searchedWordsMap[w]; ok2 {
						searchedWordsMap[w] += 1
					} else {
						searchedWordsMap[w] = 1
					}

					// 一直移除单词，直到次数符合了
					removeNum := 0
					if searchedWordsMap[w] > wordsMap[w] {
						hasRemoved = true

						for searchedWordsMap[w] > wordsMap[w] {
							firstWord := s[j+removeNum*wordLen : j+(removeNum+1)*wordLen]
							searchedWordsMap[firstWord]--
							removeNum++
						}

						foundWordsNum = foundWordsNum - removeNum + 1
						i = i + (removeNum-1)*wordLen
						break
					}
				} else {
					searchedWordsMap = make(map[string]int)
					j = j + foundWordsNum*wordLen
					foundWordsNum = 0
					break
				}
				foundWordsNum++
			}
			if foundWordsNum == len(words) {
				ret = append(ret, j)
			}

			if foundWordsNum > 0 && !hasRemoved {
				firstWord := s[j : j+wordLen]
				searchedWordsMap[firstWord]--
				foundWordsNum--
			}
		}
	}

	return ret
}

func main() {
	fmt.Println(findSubstring("wordgoodgoodgoodbestword", []string{"word", "good", "best", "word"}))
	fmt.Println(findSubstring("barfoothefoobarman", []string{"foo", "bar"}))
}
