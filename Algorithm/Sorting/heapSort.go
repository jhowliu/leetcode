package sorting

// To heapify subtree with node i

func heapify(nums []int, n int, i int) {
	root, left, right := i, 2*i+1, 2*i+2

	if left < n && nums[left] > nums[root] {
		root = left
	}

	if right < n && nums[right] > nums[root] {
		root = right
	}

	if root != i {
		swap(nums, root, i)

		// make sure the changed substree
		heapify(nums, n, root)
	}
}

func buildMaxHeap(nums []int) {
	n := len(nums)

	for i := n/2 - 1; i >= 0; i-- {
		heapify(nums, n, i)
	}
}

func swap(nums []int, i int, j int) {
	nums[i], nums[j] = nums[j], nums[i]
}

func heapSort(nums []int) {
	n := len(nums)
	end := n

	buildMaxHeap(nums)

	for i := 0; i < n; i++ {
		// swap largest number with last numbers
		swap(nums, 0, end-1)
		// avoid to heapify the sorted number
		end--

		heapify(nums, end, 0)
	}
}
