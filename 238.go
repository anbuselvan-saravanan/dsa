// 238. Product of Array Except Self
package main

func ProductExceptSelf(nums []int) []int {
	result := make([]int, len(nums))

	result[0] = 1
	product := 1
	for i := 1; i < len(nums); i++ {
		product *= nums[i-1]
		result[i] = product
	}
	product = 1
	for i := len(nums) - 1; i >= 0; i-- {
		result[i] *= product
		product *= nums[i]
	}
	return result
}
