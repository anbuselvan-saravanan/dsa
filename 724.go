// 724. Find Pivot Index
package main

func PivotIndex(nums []int) int {
	total_sum := 0
	for right := range nums {
		total_sum += nums[right]
	}

	left_sum := 0
	for right := range nums {
		right_sum := total_sum - nums[right] - left_sum
		if left_sum == right_sum {
			return right
		}
		left_sum += nums[right]
	}
	return -1
}
