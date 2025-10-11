package problems

import (
	"reflect"
	"testing"
)

func TestGenerate(t *testing.T) {
	tests := []struct {
		name    string
		numRows int
		want    [][]int
	}{
		{
			name:    "zero rows",
			numRows: 0,
			want:    [][]int{},
		},
		{
			name:    "one row",
			numRows: 1,
			want: [][]int{
				{1},
			},
		},
		{
			name:    "two rows",
			numRows: 2,
			want: [][]int{
				{1},
				{1, 1},
			},
		},
		{
			name:    "five rows",
			numRows: 5,
			want: [][]int{
				{1},
				{1, 1},
				{1, 2, 1},
				{1, 3, 3, 1},
				{1, 4, 6, 4, 1},
			},
		},
		{
			name:    "large rows",
			numRows: 10, // chỉ test chiều dài hàng cuối thôi, không so chi tiết
			want: [][]int{
				{1},
				{1, 1},
				{1, 2, 1},
				{1, 3, 3, 1},
				{1, 4, 6, 4, 1},
				{1, 5, 10, 10, 5, 1},
				{1, 6, 15, 20, 15, 6, 1},
				{1, 7, 21, 35, 35, 21, 7, 1},
				{1, 8, 28, 56, 70, 56, 28, 8, 1},
				{1, 9, 36, 84, 126, 126, 84, 36, 9, 1},
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := generate(tt.numRows)
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("generate(%d) = %v, want %v", tt.numRows, got, tt.want)
			}
		})
	}
}
