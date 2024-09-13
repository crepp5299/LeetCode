func majorityElement(nums []int) int {
	max, count := nums[0], 1
	for i := 1; i < len(nums); i++ {
		if nums[i] != max {
			if count > 0 {
				count--
			} else {
				max = nums[i]
				count = 1
			}
		} else {
			count++
		}
	}
	return max
}