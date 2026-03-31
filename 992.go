// 992. Subarrays with K Different Integers
package main

func SubarraysWithKDistinct(nums []int, k int) (count int) {
	frequency := map[int]int{}
	left := 0

	for right := range nums {
		frequency[nums[right]]++

		for len(frequency) > k && left < len(nums) {
			frequency[nums[left]]--

			if frequency[nums[left]] == 0 {
				delete(frequency, nums[left])
			}

			left++
		}

		count += right - left + 1
	}

	left = 0
	frequency = map[int]int{}
	for right := range nums {
		frequency[nums[right]]++

		for len(frequency) > k-1 && left < len(nums) {
			frequency[nums[left]]--

			if frequency[nums[left]] == 0 {
				delete(frequency, nums[left])
			}
			left++
		}

		count -= right - left + 1
	}
	return
}
