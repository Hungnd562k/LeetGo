package problems

import "strings"

func KeyboardRow(words []string) []string {
	rows := []string{
		"qwertyuiop",
		"asdfghjkl",
		"zxcvbnm",
	}
	letterToRow := make(map[rune]int)
	for i, row := range rows {
		for _, ch := range row {
			letterToRow[ch] = i
		}
	}
	var result []string
	for _, word := range words {
		lower := strings.ToLower(word)
		firstRow := letterToRow[rune(lower[0])]
		valid := true
		for _, ch := range lower {
			if letterToRow[ch] != firstRow {
				valid = false
				break
			}
		}
		if valid {
			result = append(result, word)
		}
	}

	return result
}
