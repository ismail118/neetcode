package main

import "fmt"

func main() {
	nums := []int{8, 3, 1, 6, 7, 5}
	quickShort(nums, 0, len(nums)-1)
	fmt.Println(nums)
}

func quickShort(nums []int, minIdx, maxIdx int) {

	if minIdx >= maxIdx {
		return
	}

	pivot := nums[maxIdx]

	leftIdx := minIdx
	rightIdx := maxIdx

	for leftIdx < rightIdx {
		for nums[leftIdx] <= pivot && leftIdx < rightIdx {
			leftIdx++
		}
		for nums[rightIdx] >= pivot && leftIdx < rightIdx {
			rightIdx--
		}
		nums[leftIdx], nums[rightIdx] = nums[rightIdx], nums[leftIdx]
	}

	nums[leftIdx], nums[maxIdx] = nums[maxIdx], nums[leftIdx]

	quickShort(nums, minIdx, leftIdx-1)
	quickShort(nums, leftIdx+1, maxIdx)
}
