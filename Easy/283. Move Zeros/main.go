package solution

func moveZeros(nums []int) {
	idx := 0

	// 將非0的照順序前移
	for _, num := range nums {
		if num != 0 {
			nums[idx] = num
			idx++
		}
	}

	// 把0補好補滿
	for idx < len(nums) {
		nums[idx] = 0
		idx++
	}
}
