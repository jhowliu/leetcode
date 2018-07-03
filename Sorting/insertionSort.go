package sorting

func insertionSort(nums []int) {
	n := len(nums)
	for i := 1; i < n; i++ {
		tmp := nums[i]
		j := i - 1
		for ; j >= 0 && nums[j] > tmp; j-- {
			nums[j+1] = nums[j]
		}
		nums[j+1] = tmp
	}
}
