package problems

import (
	"reflect"
	"testing"
)

func TestRemoveDuplicates(t *testing.T) {
	tests := []struct {
		name     string
		input    []int
		expected []int
		wantLen  int
	}{
		{
			name:     "Basic",
			input:    []int{1, 1, 2},
			expected: []int{1, 2},
			wantLen:  2,
		},
		{
			name:     "All unique",
			input:    []int{1, 2, 3, 4},
			expected: []int{1, 2, 3, 4},
			wantLen:  4,
		},
		{
			name:     "All duplicates",
			input:    []int{5, 5, 5, 5},
			expected: []int{5},
			wantLen:  1,
		},
		{
			name:     "Empty slice",
			input:    []int{},
			expected: []int{},
			wantLen:  0,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			nums := append([]int(nil), tt.input...) // copy để không thay input gốc
			gotLen := removeDuplicates(nums)

			if gotLen != tt.wantLen {
				t.Errorf("expected length %d, got %d", tt.wantLen, gotLen)
			}

			// So sánh phần đầu của mảng
			if !reflect.DeepEqual(nums[:gotLen], tt.expected) {
				t.Errorf("expected nums=%v, got=%v", tt.expected, nums[:gotLen])
			}
		})
	}
}
