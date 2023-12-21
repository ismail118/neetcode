package main

func main() {

}

func topKFrequent(nums []int, k int) []int {
	hashMap := make(map[int]int)

	for _, num := range nums {
		hashMap[num] = hashMap[num] + 1
	}

	var keys []int
	for i := 0; i < k; i++ {
		var max, key int
		for y, x := range hashMap {
			if x > max {
				max = x
				key = y
			}
		}
		keys = append(keys, key)
		delete(hashMap, key)

	}

	return keys
}
