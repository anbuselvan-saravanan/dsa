// 152. Maximum Product Subarray
package main

import "math"

func MaxProduct(nums []int) int {
	product := math.MinInt
	length := len(nums)

	prefix_product := 1
	suffix_product := 1

	for i := range length {
		if prefix_product == 0 {
			prefix_product = 1
		}

		if suffix_product == 0 {
			suffix_product = 1
		}

		prefix_product *= nums[i]
		suffix_product *= nums[length-i-1]

		product = max(product, prefix_product, suffix_product)
	}

	return product
}
