package hw19

import (
	"strconv"
	"strings"
)

func Search(text, pattern string) string {
	as := NewAutoSearch(pattern)
	result := as.Search(text)
	return ResultToString(result)
}

func KnuthMorrisPrattSearch(text, pattern string) string {
	kmp := NewKMPSearch(pattern)
	result := kmp.Search(text)
	return ResultToString(result)
}

func ResultToString(result []int) string {
	var sb strings.Builder
	for i := 0; i < len(result); i++ {
		sb.WriteString(strconv.Itoa(result[i]))
		sb.WriteString(" ")
	}
	return strings.Trim(sb.String(), " ")
}

// Left: первые x символов строки.
func Left(s string, x int) string {
	if x > len(s) {
		x = len(s)
	}
	return s[:x]
}

// Right: последние x символов строки.
func Right(s string, x int) string {
	if x > len(s) {
		x = len(s)
	}
	return s[len(s)-x:]
}
