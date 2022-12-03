package main

import (
	"sort"
	"strings"
	"unicode"
)

//
// https://leetcode.com/problems/sort-characters-by-frequency/submissions/854083374/
//

type Meta struct {
	Char  rune
	Count int
}

const (
	COUNT = 26*2 + 10

	DIGIT_OFFSET = 0
	LOWER_OFFSET = 10
	UPPER_OFFSET = 26 + 10
)

func FrequencySort(s string) string {
	m := map[rune]*Meta{}

	for _, c := range s {
		meta := m[c]

		if meta == nil {
			meta = &Meta{Char: c, Count: 1}
			m[c] = meta
		} else {
			m[c].Count++
		}
	}

	o := make([]*Meta, COUNT)

	for k, v := range m {
		index := 0

		if unicode.IsUpper(k) {
			index = UPPER_OFFSET + int(k-'A')
		} else if unicode.IsLower(k) {
			index = LOWER_OFFSET + int(k-'a')
		} else if unicode.IsDigit(k) {
			index = DIGIT_OFFSET + int(k-'0')
		}

		o[index] = v
	}

	sort.SliceStable(o, func(i, j int) bool {
		a := o[i]
		b := o[j]

		if a == nil {
			return false
		}

		if b == nil {
			return true
		}

		return a.Count > b.Count
	})

	builder := strings.Builder{}

	for idx := 0; idx < COUNT; idx++ {
		c := o[idx]

		if c == nil {
			break
		}

		builder.WriteString(strings.Repeat(string(c.Char), c.Count))
	}

	return builder.String()
}
