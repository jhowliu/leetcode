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
