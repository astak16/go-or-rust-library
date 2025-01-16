package string_similarity

import (
	"errors"
	"strings"
)

func CompareTwoStrings(first, second string) float64 {
	first = strings.ReplaceAll(first, " ", "")
	second = strings.ReplaceAll(second, " ", "")

	if first == second {
		return 1
	}
	if len(first) < 2 || len(second) < 2 {
		return 0
	}

	firstBigrams := make(map[string]int)
	for i := 0; i < len(first)-1; i++ {
		bigram := first[i : i+2]
		if _, ok := firstBigrams[bigram]; ok {
			firstBigrams[bigram]++
		} else {
			firstBigrams[bigram] = 1
		}
	}

	intersectionSize := 0
	for i := 0; i < len(second)-1; i++ {
		bigram := second[i : i+2]
		if count, ok := firstBigrams[bigram]; ok && count > 0 {
			firstBigrams[bigram]--
			intersectionSize++
		}
	}

	return 2.0 * float64(intersectionSize) / float64(len(first)+len(second)-2)
}

type Ratings struct {
	Target any
	Rating float64
}

type BestMatch struct {
	Ratings        []Ratings
	BestMatch      Ratings
	BestMatchIndex int
}

func FindBestMatch(args ...interface{}) (*BestMatch, error) {
	if !areArgsValid(args...) {
		return nil, errors.New("bad arguments: first argument should be a string, second should be an array of strings")
	}
	mainString := args[0]
	targetStrings := args[1]
	ratings := []Ratings{}
	bestMatchIndex := 0
	for i, target := range targetStrings.([]string) {
		currentRating := CompareTwoStrings(mainString.(string), target)
		ratings = append(ratings, Ratings{Target: target, Rating: currentRating})
		if currentRating > ratings[bestMatchIndex].Rating {
			bestMatchIndex = i
		}
	}

	bestMatch := ratings[bestMatchIndex]
	return &BestMatch{Ratings: ratings, BestMatch: bestMatch, BestMatchIndex: bestMatchIndex}, nil
}

func areArgsValid(args ...interface{}) bool {
	if len(args) < 2 {
		return false
	}
	mainString := args[0]
	targetStrings := args[1]
	if _, ok := mainString.(string); !ok {
		return false
	}
	if _, ok := targetStrings.([]string); !ok {
		return false
	}
	if len(targetStrings.([]string)) == 0 {
		return false
	}
	return true
}
