// https://leetcode.com/problems/remove-element
package main

// Uses two indices to remove elements.
func removeElement(nums []int, val int) int {
	slow := 0
	for fast := range nums {
		if val != nums[fast] {
			nums[slow] = nums[fast]
			slow++
		}
	}
	return slow
}

/* Naive version
func removeElement(nums []int, val int) int {
	size := len(nums)
	for i := 0; i < size; i++ {
		if nums[i] == val {
			for j := i + 1; j < size; j++ {
				nums[j-1] = nums[j]
			}
			// Rechecks the overwritten `nums[i]`
			i--
			// Original `num[i]` removed
			size--
		}
	}
	return size
}
*/
