package problems

import "testing"

// Định nghĩa cấu trúc cho các trường hợp kiểm thử
type testCase struct {
	roman    string // Đầu vào: Chuỗi số La Mã
	expected int    // Đầu ra mong đợi: Số nguyên tương ứng
}

// Hàm Test chính
func TestRomanToInteger(t *testing.T) {
	// Danh sách các trường hợp kiểm thử
	tests := []testCase{
		{"III", 3},        // Trường hợp cơ bản
		{"LVIII", 58},     // Trường hợp phức tạp hơn (L = 50, V = 5, III = 3)
		{"MCMXCIV", 1994}, // Trường hợp có quy tắc trừ: M=1000, CM=900, XC=90, IV=4
		{"IX", 9},         // Quy tắc trừ cơ bản (10 - 1)
		{"IV", 4},         // Quy tắc trừ cơ bản (5 - 1)
		{"XL", 40},        // Quy tắc trừ
		{"XC", 90},        // Quy tắc trừ
		{"CD", 400},       // Quy tắc trừ
		{"CM", 900},       // Quy tắc trừ
		{"V", 5},
		{"DCCCLXXXIV", 884}, // Trường hợp dài
		{"MMXXV", 2025},     // Trường hợp số lớn
		{"CMLII", 952},
	}

	// Lặp qua từng trường hợp kiểm thử
	for _, test := range tests {
		// Gọi hàm cần kiểm tra
		result := RomanToInt(test.roman)

		// So sánh kết quả thực tế với kết quả mong đợi
		if result != test.expected {
			t.Errorf("RomanToInt(%s) sai.\nKết quả thực tế: %d, Kết quả mong đợi: %d",
				test.roman, result, test.expected)
		}
	}
}
