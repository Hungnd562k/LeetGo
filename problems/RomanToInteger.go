package problems

import (
	"fmt"
	"strings"
)

func RomanToInt(s string) int {
	strRoman := ""
	quoteChar := "\""
	firstIndex := strings.Index(s, quoteChar)
	lastIndex := strings.LastIndex(s, quoteChar)

	if firstIndex != -1 && lastIndex != -1 && firstIndex < lastIndex {
		strRoman = s[firstIndex+1 : lastIndex]
	}
	fmt.Println(strRoman)
	return -1
}
