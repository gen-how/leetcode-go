package main

// https://leetcode.com/problems/binary-search
func search(nums []int, target int) int {
	// Searches for the interval of [l, r).
	l, r := 0, len(nums)
	for l < r {
		// m := l + (r-l)/2
		m := (l + r) / 2
		if nums[m] > target {
			r = m
		} else if nums[m] < target {
			l = m + 1
		} else {
			return m
		}
	}
	return -1
}
