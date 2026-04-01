// 560. Subarray Sum Equals K
package main

/*
*

		s-k		k
	|-------||------|
	|_______________|
					prefix-sum
	number of subarray gives sum k == number of subarray gives sum s-k

*
*/
func SubarraySum(nums []int, k int) (count int) {
	hash_map := map[int]int{0: 1}
	sum := 0
	for right := range nums {
		sum += nums[right]

		s_k := sum - k
		count += hash_map[s_k]
		hash_map[sum]++
	}

	return
}
