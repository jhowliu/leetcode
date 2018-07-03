package sorting

func insertionSort(nums []int) {
	n := len(nums)
	for i := 1; i < n; i++ {
		key := nums[i]
		j := i - 1
		for ; j >= 0 && nums[j] > key; j-- {
			nums[j+1] = nums[j]
		}
		nums[j+1] = key
	}
}

func insertionSortRecursively(nums []int, n int) {
	if n <= 1 {
		return
	}

	insertionSortRecursively(nums, n-1)

	key := nums[n-1]
	j := n - 2

	for ; j >= 0 && nums[j] > key; j-- {
		nums[j+1] = nums[j]
	}
	nums[j+1] = key
}

// using binary search to reduce comparisons.
func binaryInsertionSort(nums []int, n int) {
	for i := 1; i < n; i++ {
		key := nums[i]
		j := i - 1
		ix := binarySearch(nums, 0, i, key)

		for ; j >= ix; j-- {
			nums[j+1] = nums[j]
		}
		nums[ix] = key
	}
}

func binarySearch(nums []int, left int, right int, target int) int {
	// in the last stack (not found the same number)
	if left >= right {
		if target > nums[left] {
			return left + 1
		}
		return left
	}
	mid := (left + right) / 2

	if nums[mid] == target {
		return mid + 1
	}

	if nums[mid] >= target {
		return binarySearch(nums, left, mid-1, target)
	}

	return binarySearch(nums, mid+1, right, target)
}
