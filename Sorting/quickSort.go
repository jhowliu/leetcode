package sorting

func quickSort(nums []int, left int, right int) {
	if left < right {
		mid := partition(nums, left, right)
		// Sorting left of pivot
		quickSort(nums, left, mid-1)
		quickSort(nums, mid+1, right)
	}
}

func partition(nums []int, left int, right int) int {
	// pick the last element as pivot
	pivot := nums[right]

	i, j := left-1, left

	for j < right {
		// 若nums[j]大於Pivot則不會進入此判斷式，i會停留在前一個位置，
		// 這代表nums[i+1]一定是大於Pivot，直到找到下一個小於等於Pivot的元素，將兩者交換。
		if nums[j] <= pivot {
			i++
			swap(nums, i, j)
		}
		j++
	}

	// swap the larger element(i+1) and pivot
	swap(nums, i+1, right)
	// return index of pivot
	return i + 1
}
