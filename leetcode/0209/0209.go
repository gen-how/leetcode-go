// https://leetcode.com/problems/minimum-size-subarray-sum
package main

const maxInt32 int = 1<<31 - 1

// Uses sliding window.
func minSubArrayLen(target int, nums []int) int {
	l := 0
	sum := 0
	minLen := maxInt32
	for r := range nums {
		sum += nums[r]
		for sum >= target {
			minLen = min(minLen, r-l+1)
			sum -= nums[l]
			l++
		}
	}

	// This means not found.
	if minLen == maxInt32 {
		return 0
	}
	return minLen
}
