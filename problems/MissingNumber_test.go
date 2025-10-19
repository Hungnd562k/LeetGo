package problems

import (
	"math/rand"
	"testing"
)

func TestMissingNumber_TableDriven(t *testing.T) {
	tests := []struct {
		name string
		nums []int
		want int
	}{
		{"example1", []int{3, 0, 1}, 2},
		{"example2", []int{0, 1}, 2},
		{"single0", []int{0}, 1},
		{"single1", []int{1}, 0},
		{"unsorted", []int{9, 6, 4, 2, 3, 5, 7, 0, 1}, 8},
		{"missing middle", []int{0, 2, 3, 4, 5}, 1},
		{"missing last", []int{0, 1, 2, 3}, 4},
		{"missing first", []int{1, 2, 3, 4}, 0},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			got := MissingNumber(tc.nums)
			if got != tc.want {
				t.Fatalf("MissingNumber(%v) = %d; want %d", tc.nums, got, tc.want)
			}
		})
	}
}

// Test lớn với dữ liệu sinh ngẫu nhiên nhưng seed cố định để test luôn lặp lại.
func TestMissingNumber_LargeRandom(t *testing.T) {
	const n = 10000
	seed := int64(42)
	r := rand.New(rand.NewSource(seed))

	// tạo slice 0..n
	all := make([]int, n+1)
	for i := 0; i <= n; i++ {
		all[i] = i
	}

	// chọn 1 index để loại ra
	missingIdx := r.Intn(n + 1)
	missingVal := all[missingIdx]

	// xây slice nums chứa tất cả trừ missingIdx, rồi shuffle
	nums := make([]int, 0, n)
	for i := 0; i <= n; i++ {
		if i == missingIdx {
			continue
		}
		nums = append(nums, i)
	}
	// shuffle predictable bằng cùng seed
	r.Shuffle(len(nums), func(i, j int) { nums[i], nums[j] = nums[j], nums[i] })

	got := MissingNumber(nums)
	if got != missingVal {
		t.Fatalf("LargeRandom: got %d, want %d (missing index %d)", got, missingVal, missingIdx)
	}
}

// Ví dụ benchmark (tùy ý)
func BenchmarkMissingNumber(b *testing.B) {
	const n = 100000
	// chuẩn bị dữ liệu một lần
	nums := make([]int, n)
	for i := 0; i < n; i++ {
		nums[i] = i // giả sử thiếu số n
	}
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_ = MissingNumber(nums)
	}
}
