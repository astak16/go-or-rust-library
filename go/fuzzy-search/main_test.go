package fuzzysearch

import (
	"fmt"
	"testing"
)

func TestFuzzySearch(t *testing.T) {
	testData := map[int]struct {
		needle   string
		haystack string
		expected bool
	}{
		0:  {needle: "car", haystack: "cartwheel", expected: true},
		1:  {needle: "cwhl", haystack: "cartwheel", expected: true},
		2:  {needle: "cwheel", haystack: "cartwheel", expected: true},
		3:  {needle: "cartwheel", haystack: "cartwheel", expected: true},
		4:  {needle: "cwheeel", haystack: "cartwheel", expected: false},
		5:  {needle: "lw", haystack: "cartwheel", expected: false},
		6:  {needle: "语言", haystack: "php语言", expected: true},
		7:  {needle: "hp语", haystack: "php语言", expected: true},
		8:  {needle: "Py开发", haystack: "Python开发者", expected: true},
		9:  {needle: "Py 开发", haystack: "Python开发者", expected: false},
		10: {needle: "爪哇进阶", haystack: "爪哇开发进阶", expected: true},
		11: {needle: "格式工具", haystack: "非常简单的格式化工具", expected: true},
		12: {needle: "正则", haystack: "学习正则表达式怎么学习", expected: true},
		13: {needle: "学习正则", haystack: "正则表达式怎么学习", expected: false},
	}

	for i, v := range testData {
		actual := FuzzySearch(v.needle, v.haystack)
		fmt.Println(actual, v.expected, actual == v.expected)
		if actual != v.expected {
			t.Errorf("Test case %d failed: expected %v, got %v", i, v.expected, actual)
		}
	}
}
