package cookbook

import (
	"regexp"
	"strings"
)

func mostCommonWord(paragraph string, banned []string) string {
	bannedMap := map[string]bool{}
	for _, item := range banned {
		bannedMap[item] = true
	}

	wordCnt := map[string]int{}
	maxCnt := 0
	rexp, _ := regexp.Compile(`\w+`)
	for _, item := range rexp.FindAllString(paragraph, -1) {
		item = strings.ToLower(item)
		if bannedMap[item] {
			continue
		}

		wordCnt[item]++
		if wordCnt[item] > maxCnt {
			maxCnt = wordCnt[item]
		}
	}

	for key, cnt := range wordCnt {
		if cnt == maxCnt {
			return key
		}
	}
	return ""
}
