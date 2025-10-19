package problems

func MissingNumber(nums []int) int {
	n := len(nums)
	xor := 0
	for i := 0; i < n; i++ {
		xor ^= nums[i]
		xor ^= i + 1
	}

	return xor
}
