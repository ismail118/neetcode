package main

func main() {
	containsDuplicate([]int{1, 2, 3, 1})
}

// O(n2)
//func containsDuplicate(nums []int) bool {
//	for i, num := range nums {
//		for _, x := range nums[i+1:] {
//			if num == x {
//				return true
//			}
//		}
//	}
//
//	return false
//}

// O(n)
func containsDuplicate(nums []int) bool {
	hashSet := make(map[int]struct{})

	for _, x := range nums {
		_, ok := hashSet[x]
		if ok {
			return true
		} else {
			hashSet[x] = struct{}{}
		}
	}

	return false
}
