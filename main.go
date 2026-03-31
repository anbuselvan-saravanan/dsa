package main

import "fmt"

func main() {
	problem := 209

	switch problem {
	case 992:
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

	case 902:
		_fruits := [][]int{
			{1, 2, 1},
			{0, 1, 2, 2},
			{1, 2, 3, 2, 2},
		}

		fruit := _fruits[2]
		fmt.Println("904. Fruit Into Baskets")
		fmt.Println("Input")
		fmt.Println("fruits: ", fruit)
		fmt.Println("Output: ", TotalFruit(fruit))

	case 209:
		_target := [][]int{
			{2, 3, 1, 2, 4, 3},
			{1, 4, 4},
			{1, 1, 1, 1, 1, 1, 1, 1},
		}

		_k := []int{
			7,
			4,
			11,
		}

		target := _target[0]
		k := _k[0]
		fmt.Println("209. Minimum Size Subarray Sum")
		fmt.Println("Input")
		fmt.Println("nums: ", target)
		fmt.Println("k: ", k)
		fmt.Println("Output: ", MinSubArrayLen(k, target))

	}
}
