package hw19

import (
	"sort"
	"strings"
)

type AutoSearch struct {
	alphabet string
	pattern  string
	length   int
	delta    [][]int
}

func NewAutoSearch(pattern string) *AutoSearch {
	au := &AutoSearch{pattern: pattern, length: len(pattern)}
	au.createAlphabet()
	au.createDelta()
	return au
}

func (as *AutoSearch) createAlphabet() {
	alfa := as.createRunes()

	sort.Slice(alfa, func(i, j int) bool {
		return alfa[i] < alfa[j]
	})

	var sb strings.Builder
	for c := alfa[0]; c <= alfa[len(alfa)-1]; c++ {
		sb.WriteRune(c)
	}

	as.alphabet = sb.String()
}

func (as *AutoSearch) createRunes() []rune {
	seen := make(map[rune]bool)
	var alfa []rune

	for _, c := range as.pattern {
		if !seen[c] {
			seen[c] = true
			alfa = append(alfa, c)
		}
	}
	return alfa
}

func (as *AutoSearch) createDelta() {
	as.delta = make([][]int, as.length+1)
	for i := range as.delta {
		as.delta[i] = make([]int, len(as.alphabet))
	}

	for i := 0; i <= as.length; i++ {
		for _, c := range as.alphabet {
			k := i + 1
			if k > as.length {
				k--
			}
			line := Left(as.pattern, i) + string(c)
			for k >= 0 && Left(as.pattern, k) != Right(line, k) {
				k--
			}
			idx := c - rune(as.alphabet[0])
			as.delta[i][idx] = k
		}
	}
}

func (as *AutoSearch) Search(text string) []int {
	res := make([]int, 0)
	z := 0
	for {
		z = as.searching(text, z+1)
		if z == -1 {
			return res
		}
		res = append(res, z)
	}
}

func (as *AutoSearch) searching(text string, pos int) int {
	q := 0
	for i := pos; i < len(text); i++ {
		index := int(text[i] - as.alphabet[0])
		if index < 0 || index >= len(as.alphabet) {
			// Символ не из алфавита — можно обработать иначе при необходимости
			q = 0
			continue
		}
		q = as.delta[q][index]
		if q == as.length {
			return i - as.length + 1
		}
	}
	return -1
}
