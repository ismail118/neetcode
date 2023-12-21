package main

import "fmt"

func main() {

}

func productExceptSelf(nums []int) []int {
	prefixCount := make([]int, len(nums))
	postfixCount := make([]int, len(nums))

	i := 0
	j := len(nums) - 1
	for i < len(nums) {
		if i > 0 {
			prefixCount[i] = nums[i] * prefixCount[i-1]
			postfixCount[j] = nums[j] * postfixCount[j+1]
		} else {
			// initiate count
			prefixCount[i] = nums[i]
			postfixCount[j] = nums[j]
		}
		i++
		j--
	}

	fmt.Println(prefixCount)
	fmt.Println(postfixCount)

	result := make([]int, len(nums))
	for i = 0; i < len(nums); i++ {
		if i == 0 {
			result[i] = 1 * postfixCount[i+1]
		} else if i == len(nums)-1 {
			result[i] = 1 * prefixCount[i-1]
		} else {
			result[i] = prefixCount[i-1] * postfixCount[i+1]
		}
	}

	return result
}
