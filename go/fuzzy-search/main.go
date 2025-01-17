package fuzzysearch

import "fmt"

func FuzzySearch(needle, haystack string) bool {
	needleRunes := []rune(needle)
	haystackRunes := []rune(haystack)

	hlen := len(haystackRunes)
	nlen := len(needleRunes)
	if hlen < nlen {
		return false
	}
	if hlen == nlen {
		return needle == haystack
	}

	j := 0
outer:
	for i := 0; i < nlen; i++ {
		nch := needleRunes[i]
		for j < hlen {
			fmt.Println(i, haystackRunes[j], nch, haystackRunes[j] == nch)
			if haystackRunes[j] == nch {
				j++
				continue outer
			}
			j++
		}
		return false
	}
	return true
}
