// 1314. Matrix Block Sum
package main

import "fmt"

func MatrixBlockSum(mat [][]int, k int) (result [][]int) {
	prefix_sum := [][]int{}

	for i := range mat {
		sum := 0
		prefix_sum = append(prefix_sum, []int{})
		for j := range mat[i] {
			sum += mat[i][j]
			prefix_sum[i] = append(prefix_sum[i], sum)
		}
	}

	fmt.Println(prefix_sum)
	return
}
