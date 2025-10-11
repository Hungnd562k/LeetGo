package problems

func searchInsert(nums []int, target int) int {
	if len(nums) == 0 {
		return 0
	}

	for i, _ := range nums[0:] {
		if target <= nums[i] {
			if i == 0 {
				return 0
			} else {
				return i
			}
		} else {
			if i == len(nums)-1 {
				return i + 1
			} else {
				continue
			}
		}
	}
	return -1
}
