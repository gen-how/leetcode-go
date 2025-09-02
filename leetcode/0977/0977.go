// https://leetcode.com/problems/squares-of-a-sorted-array
package main

func abs(n int) int {
	if n < 0 {
		return -n
	}
	return n
}

// Uses two pointer moving in opppsite directions.
func sortedSquares(nums []int) []int {
	result := make([]int, len(nums))
	k := len(nums) - 1
	for l, r := 0, len(nums)-1; l <= r; {
		if abs(nums[l]) < abs(nums[r]) {
			result[k] = nums[r] * nums[r]
			r--
		} else {
			result[k] = nums[l] * nums[l]
			l++
		}
		k--
	}
	return result
}

/* Naive version
func sortedSquaresNaive(nums []int) []int {
	for i, v := range nums {
		nums[i] = v * v
	}
	slices.Sort(nums)
	return nums
}
*/
