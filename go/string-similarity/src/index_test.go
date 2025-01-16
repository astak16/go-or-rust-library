package string_similarity

import (
	"math"
	"testing"
)

func TestCompareTwoStrings(t *testing.T) {
	testData := map[int]struct {
		first    string
		second   string
		expected float64
	}{
		0:  {first: "french", second: "quebec", expected: 0},
		1:  {first: "france", second: "france", expected: 1},
		2:  {first: "fRaNce", second: "france", expected: 0.2},
		3:  {first: "healed", second: "sealed", expected: 0.8},
		4:  {first: "web applications", second: "applications of the web", expected: 0.7878787878787878},
		5:  {first: "this will have a typo somewhere", second: "this will huve a typo somewhere", expected: 0.92},
		6:  {first: "Olive-green table for sale, in extremely good condition.", second: "For sale: table in very good  condition, olive green in colour.", expected: 0.6060606060606061},
		7:  {first: "Olive-green table for sale, in extremely good condition.", second: "For sale: green Subaru Impreza, 210,000 miles", expected: 0.2558139534883721},
		8:  {first: "Olive-green table for sale, in extremely good condition.", second: "Wanted: mountain bike with at least 21 gears.", expected: 0.1411764705882353},
		9:  {first: "this has one extra word", second: "this has one word", expected: 0.7741935483870968},
		10: {first: "a", second: "a", expected: 1},
		11: {first: "a", second: "b", expected: 0},
		12: {first: "", second: "", expected: 1},
		13: {first: "a", second: "", expected: 0},
		14: {first: "", second: "a", expected: 0},
		15: {first: "apple event", second: "apple    event", expected: 1},
		16: {first: "iphone", second: "iphone x", expected: 0.9090909090909091},
	}
	for i, v := range testData {
		actual := CompareTwoStrings(v.first, v.second)
		if actual != v.expected {
			t.Errorf("Test case %d failed: expected %f, got %f", i, v.expected, actual)
		}
	}
}

func TestFindBestMatch(t *testing.T) {
	badArgsErrorMsg := "bad arguments: first argument should be a string, second should be an array of strings"
	_, err := FindBestMatch("one", []string{"two", "three"})
	if err != nil {
		t.Errorf("Test case 0 failed: expected nil, got %v", err)
	}

	tests := []struct {
		name     string
		args     []interface{}
		wantErr  bool
		errorMsg string
	}{
		{
			name:     "Empty slice",
			args:     []interface{}{[]string{}},
			wantErr:  true,
			errorMsg: badArgsErrorMsg,
		},
		{
			name:     "Empty string in slice",
			args:     []interface{}{[]string{""}},
			wantErr:  true,
			errorMsg: badArgsErrorMsg,
		},
		{
			name:     "Invalid type int slice",
			args:     []interface{}{[]int{8}},
			wantErr:  true,
			errorMsg: badArgsErrorMsg,
		},
		{
			name:     "Valid string slice",
			args:     []interface{}{[]string{"hello", "something"}},
			wantErr:  true,
			errorMsg: badArgsErrorMsg,
		},
		{
			name:     "Mixed types with empty slice",
			args:     []interface{}{[]interface{}{"hello", []string{}}},
			wantErr:  true,
			errorMsg: badArgsErrorMsg,
		},
		{
			name:     "Mixed types with nested interface",
			args:     []interface{}{[]interface{}{"hello", []interface{}{2, "world"}}},
			wantErr:  true,
			errorMsg: badArgsErrorMsg,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			_, err := FindBestMatch(tt.args...)
			if (err == nil) != !tt.wantErr {
				t.Errorf("FindBestMatch() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if err != nil && err.Error() != tt.errorMsg {
				t.Errorf("FindBestMatch() error message = %v, want %v", err.Error(), tt.errorMsg)
			}
		})
	}
}

func TestFindBestMatchRatings(t *testing.T) {
	mainString := "healed"
	targetStrings := []string{"mailed", "edward", "sealed", "theatre"}

	expectedRatings := []Ratings{
		{Target: "mailed", Rating: 0.4},
		{Target: "edward", Rating: 0.2},
		{Target: "sealed", Rating: 0.8},
		{Target: "theatre", Rating: 0.36363636363636365},
	}

	matches, err := FindBestMatch(mainString, targetStrings)
	if err != nil {
		t.Fatalf("FindBestMatch failed: %v", err)
	}

	// 检查结果长度
	if len(matches.Ratings) != len(expectedRatings) {
		t.Errorf("Expected %d ratings, got %d", len(expectedRatings), len(matches.Ratings))
	}

	// 逐个比较评分
	for i, expected := range expectedRatings {
		got := matches.Ratings[i]

		// 比较目标字符串
		if got.Target != expected.Target {
			t.Errorf("Rating[%d].Target = %v, want %v", i, got.Target, expected.Target)
		}

		// 比较评分值（使用浮点数比较）
		if !almostEqual(got.Rating, expected.Rating, 0.000001) {
			t.Errorf("Rating[%d].Rating = %v, want %v", i, got.Rating, expected.Rating)
		}
	}
}

func TestFindBestMatchResults(t *testing.T) {
	tests := []struct {
		name          string
		mainString    string
		targetStrings []string
		wantBest      Ratings
		wantIndex     int
	}{
		{
			name:          "find best match for 'healed'",
			mainString:    "healed",
			targetStrings: []string{"mailed", "edward", "sealed", "theatre"},
			wantBest:      Ratings{Target: "sealed", Rating: 0.8},
			wantIndex:     2,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			matches, err := FindBestMatch(tt.mainString, tt.targetStrings)
			if err != nil {
				t.Fatalf("FindBestMatch failed: %v", err)
			}

			// 测试最佳匹配
			if matches.BestMatch.Target != tt.wantBest.Target {
				t.Errorf("BestMatch.Target = %v, want %v", matches.BestMatch.Target, tt.wantBest.Target)
			}
			if !almostEqual(matches.BestMatch.Rating, tt.wantBest.Rating, 0.000001) {
				t.Errorf("BestMatch.Rating = %v, want %v", matches.BestMatch.Rating, tt.wantBest.Rating)
			}

			// 测试最佳匹配索引
			if matches.BestMatchIndex != tt.wantIndex {
				t.Errorf("BestMatchIndex = %v, want %v", matches.BestMatchIndex, tt.wantIndex)
			}
		})
	}
}

// almostEqual 用于比较浮点数
func almostEqual(a, b, tolerance float64) bool {
	return math.Abs(a-b) <= tolerance
}
