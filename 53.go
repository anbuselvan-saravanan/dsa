// 53. Maximum Subarray
package main

func MaxSubArray(nums []int) (sum int) {
	right := 0
	sub_array_sum := nums[right]
	sum = max(sub_array_sum, nums[right])

	if sub_array_sum < 0 {
		sub_array_sum = 0
	}

	for right = right + 1; right < len(nums); right++ {
		sub_array_sum += nums[right]

		sum = max(sub_array_sum, sum)

		if sub_array_sum < 0 {
			sub_array_sum = 0
		}
	}
	return
}
