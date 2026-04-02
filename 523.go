// 523. Continuous Subarray Sum
package main

/*
*
Look for the already existing remainder that matches with current prefix sum remainder.
If the same remainder comes again then there is k in between these two indices
*
*/
func CheckSubarraySum(nums []int, k int) bool {
	hash_map := map[int]int{0: -1}

	prefix_sum := 0
	for right := range nums {
		prefix_sum += nums[right]

		remainder := prefix_sum % k
		value, ok := hash_map[remainder]

		if ok {
			if right-value > 1 {
				return true
			}
		} else {
			hash_map[remainder] = right
		}
	}
	return false
}
