package main

import "fmt"

func main() {
	nums := []int{8, 3, 1, 6, 7, 5}
	bubbleSort(nums)
	fmt.Println(nums)
}

func bubbleSort(nums []int) {
	for i := 0; i < len(nums); i++ {
		for j := i + 1; j < len(nums); j++ {
			if nums[i] > nums[j] {
				nums[i], nums[j] = nums[j], nums[i]
			}
		}
	}
}
