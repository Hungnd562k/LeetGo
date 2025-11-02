package problems

import "testing"

func TestRomanToInt(t *testing.T) {
	result := RomanToInt("s = \"MCMXCIV\"")
	t.Errorf("Khong lam gi ca, %v", result)
}
