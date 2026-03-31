package main

import "fmt"

func main() {
	_target := [][]int{
		{1, 2, 1, 2, 3},
		{1, 2, 1, 3, 4},
	}

	_k := []int{
		2,
		3,
	}

	target := _target[1]
	k := _k[1]
	fmt.Println("992. Subarrays with K Different Integers")
	fmt.Println("Input")
	fmt.Println("nums: ", target)
	fmt.Println("k: ", k)
	fmt.Println("Output: ", SubarraysWithKDistinct(target, k))
}
