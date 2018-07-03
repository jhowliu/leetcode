package sorting

func mergeSort(nums []int, left int, right int) {
	if left < right {
		mid := (left + right) / 2

		mergeSort(nums, left, mid)
		mergeSort(nums, mid+1, right)

		merge(nums, mid, left, right)
	}
}

func merge(nums []int, mid int, left int, right int) {
	i, j, k := left, mid+1, left

	tmp := make([]int, len(nums))

	copy(tmp, nums)

	for i < mid+1 && j <= right {
		if tmp[i] <= tmp[j] {
			nums[k] = tmp[i]
			i++
		} else {
			nums[k] = tmp[j]
			j++
		}
		k++
	}
	// append remaining elements
	for i < mid+1 {
		nums[k] = tmp[i]
		k++
		i++
	}
	// append remaining elements
	for j <= right {
		nums[k] = tmp[j]
		k++
		j++
	}
}
