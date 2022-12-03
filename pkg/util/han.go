package util

import "unicode"

func HanCount(s string) int {
	var count int
	for _, c := range s {
		if unicode.Is(unicode.Han, c) || c == '　' {
			count++
		}
	}
	return count
}
