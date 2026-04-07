package main

import "fmt"

func main() {
	problem := 724

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

	case 560:
		_target := [][]int{
			{1, 1, 1},
			{1, 2, 3},
			{1, 2, 3, -3, 1, 1, 1, 4, 2, -3},
		}

		_k := []int{
			2,
			3,
			3,
		}

		target := _target[0]
		k := _k[0]
		fmt.Println("560. Subarray Sum Equals K")
		fmt.Println("Input")
		fmt.Println("nums: ", target)
		fmt.Println("k: ", k)
		fmt.Println("Output: ", SubarraySum(target, k))

	case 1314:
		_target := [][][]int{
			{{1, 2, 3}, {4, 5, 6}, {7, 8, 9}},
			{{1, 2, 3}, {4, 5, 6}, {7, 8, 9}},
		}

		_k := []int{
			1,
			2,
		}

		target := _target[0]
		k := _k[0]
		fmt.Println("1314. Matrix Block Sum")
		fmt.Println("Input")
		fmt.Println("nums: ", target)
		fmt.Println("k: ", k)
		fmt.Println("Output: ", MatrixBlockSum(target, k))

	case 238:
		_target := [][]int{
			{1, 2, 3, 4},
			{-1, 1, 0, -3, 3},
		}

		target := _target[0]
		fmt.Println("238. Product of Array Except Self")
		fmt.Println("Input")
		fmt.Println("nums: ", target)
		fmt.Println("Output: ", ProductExceptSelf(target))

	case 523:
		_target := [][]int{
			{23, 2, 4, 6, 7},
			{23, 2, 6, 4, 7},
			{5, 0, 0, 0},
		}

		_k := []int{
			6,
			6,
			3,
		}

		target := _target[2]
		k := _k[2]
		fmt.Println("523. Continuous Subarray Sum")
		fmt.Println("Input")
		fmt.Println("nums: ", target)
		fmt.Println("k: ", k)
		fmt.Println("Output: ", CheckSubarraySum(target, k))

	case 974:
		_target := [][]int{
			{4, 5, 0, -2, -3, 1},
			{5},
			{-1, 2, 9},
		}

		_k := []int{
			5,
			9,
			2,
		}

		target := _target[0]
		k := _k[0]
		fmt.Println("974. Subarray Sums Divisible by K")
		fmt.Println("Input")
		fmt.Println("nums: ", target)
		fmt.Println("k: ", k)
		fmt.Println("Output: ", SubarraysDivByK(target, k))

	case 724:
		_input := [][]int{
			{1, 7, 3, 6, 5, 6},
			{1, 2, 3},
			{2, 1, -1},
		}

		input := _input[0]
		fmt.Println("724. Find Pivot Index")
		fmt.Println("Input")
		fmt.Println("nums: ", input)
		fmt.Println("Output: ", PivotIndex(input))
	}
}
