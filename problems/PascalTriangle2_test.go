package problems

import (
	"reflect"
	"testing"
)

func TestGetRow(t *testing.T) {
	tests := []struct {
		name     string
		rowIndex int
		want     []int
	}{
		{"Row 0", 0, []int{1}},
		{"Row 1", 1, []int{1, 1}},
		{"Row 2", 2, []int{1, 2, 1}},
		{"Row 3", 3, []int{1, 3, 3, 1}},
		{"Row 4", 4, []int{1, 4, 6, 4, 1}},
		{"Row 5", 5, []int{1, 5, 10, 10, 5, 1}},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := getRow(tt.rowIndex)
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("GetRow(%d) = %v, want %v", tt.rowIndex, got, tt.want)
			}
		})
	}
}
