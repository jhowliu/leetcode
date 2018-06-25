/*
A non-empty array A consisting of N integers is given.
The array contains an odd number of elements,
and each element of the array can be paired with another element that has the same value,
except for one element that is left unpaired.

For Example:
	[1, 2, 3, 5, 1, 2, 3] => 5
*/
package solution

func findOddOccurencesNumber(nums []int) int {
	result := nums[0]
	for _, num := range nums[1:] {
		result ^= num
	}

	return result
}
