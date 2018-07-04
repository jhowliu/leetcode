/*
We define a harmonious array is an array where the difference between its maximum value and its minimum value is exactly 1.

Example 1:
	Input: [1,3,2,2,5,2,3,7]
	Output: 5
Explanation: The longest harmonious subsequence is [3,2,2,2,3].
Note: The length of the input array will not exceed 20,000.

*/
package solution

func findLHS(nums []int) int {
	counter := map[int]int{}

	for _, num := range nums {
		counter[num]++
	}

	max := 0

	for k, v := range counter {
		count := v
		if value, ok := counter[k-1]; ok {
			count += value
		} else {
			continue
		}

		if count > max {
			max = count
		}
	}

	return max
}
