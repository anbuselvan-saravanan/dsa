// 209. Minimum Size Subarray Sum
package main

import "math"

func MinSubArrayLen(target int, nums []int) (length int) {
	length = math.MaxInt
	left := 0
	right := 0
	sum := 0

	for right = range nums {
		sum += nums[right]

		for sum >= target {
			length = min(length, right-left+1)
			sum -= nums[left]
			left++
		}
	}

	if length == math.MaxInt {
		length = 0
	}
	return
}
