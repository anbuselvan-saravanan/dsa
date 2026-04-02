// 974. Subarray Sums Divisible by K
package main

/*
*
Sum of frequencies of remainders.
*
*/
func SubarraysDivByK(nums []int, k int) (count int) {
	hash_map := make(map[int]int, len(nums))

	hash_map[0] = 1
	prefix_sum := 0
	for right := range nums {
		prefix_sum += nums[right]

		remainder := prefix_sum % k

		// If the remainder is negative, k should be added to the remainder to make it positive
		if remainder < 0 {
			remainder += k
		}

		count += hash_map[remainder]
		hash_map[remainder]++
	}
	return
}
