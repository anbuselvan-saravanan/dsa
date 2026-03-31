// 904. Fruit Into Baskets
package main

func TotalFruit(fruits []int) (count int) {
	left := 0
	frequency := map[int]int{}

	for right := range fruits {
		frequency[fruits[right]]++

		for len(frequency) > 2 && left < len(fruits) {
			frequency[fruits[left]]--

			if frequency[fruits[left]] == 0 {
				delete(frequency, fruits[left])
			}
			left++
		}

		_count := right - left + 1

		if _count > count {
			count = _count
		}
	}

	return
}
