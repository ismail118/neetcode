package main

import (
	"fmt"
	"sort"
)

func main() {
	nums := []int{-1, 0, 3, 4, 9, 12}
	num := BinarySearch(nums, 9)
	fmt.Println(num)
}

func BinarySearch(nums []int, target int) int {
	// for make sure slice sorted
	sort.Ints(nums)

	l := 0
	r := len(nums) - 1

	for {
		m := (l + r) / 2
		if nums[m] == target {
			return nums[m]
		} else if nums[m] > target {
			r = m - 1
		} else {
			l = m + 1
		}
		if l == r {
			break
		}
	}

	return -1
}
