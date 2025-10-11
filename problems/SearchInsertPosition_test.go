package problems

import "testing"

func TestSearchInsert(t *testing.T) {
	tests := []struct {
		nums     []int
		target   int
		expected int
	}{
		{[]int{1, 3, 5, 6}, 5, 2}, // target có trong mảng
		{[]int{1, 3, 5, 6}, 2, 1}, // target chèn vào giữa
		{[]int{1, 3, 5, 6}, 7, 4}, // target lớn nhất => cuối
		{[]int{1, 3, 5, 6}, 0, 0}, // target nhỏ nhất => đầu
		{[]int{}, 10, 0},          // mảng rỗng
	}

	for _, tt := range tests {
		result := searchInsert(tt.nums, tt.target)
		if result != tt.expected {
			t.Errorf("searchInsert(%v, %d) = %d; expected %d",
				tt.nums, tt.target, result, tt.expected)
		}
	}
}
