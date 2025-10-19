package problems

import (
	"reflect"
	"testing"
)

func TestKeyboardRow(t *testing.T) {
	tests := []struct {
		input    []string
		expected []string
	}{
		{[]string{"Hello", "Alaska", "Dad", "Peace"}, []string{"Alaska", "Dad"}},
		{[]string{"omk"}, []string{}},
		{[]string{"adsdf", "sfd"}, []string{"adsdf", "sfd"}},
	}

	for _, tt := range tests {
		result := KeyboardRow(tt.input)
		if result == nil {
			result = []string{}
		}
		if !reflect.DeepEqual(result, tt.expected) {
			t.Errorf("For input %v, expected %q but got %q", tt.input, tt.expected, result)
		}
	}
}
