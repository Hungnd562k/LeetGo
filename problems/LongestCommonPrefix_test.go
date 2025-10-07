package problems

import "testing"

func TestLongestCommonPrefix(t *testing.T) {
	tests := []struct {
		input    []string
		expected string
	}{
		{[]string{"flower", "flow", "flight"}, "fl"},
		{[]string{"dog", "racecar", "car"}, ""},
		{[]string{"interspecies", "interstellar", "interstate"}, "inters"},
		{[]string{"a"}, "a"},
		{[]string{"ab", "a"}, "a"},
		{[]string{"prefix", "prefixes", "prefecture"}, "pref"},
	}

	for _, tt := range tests {
		result := longestCommonPrefix(tt.input)
		if result != tt.expected {
			t.Errorf("For input %v, expected %q but got %q", tt.input, tt.expected, result)
		}
	}
}
