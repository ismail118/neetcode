package main

func main() {
	groupAnagrams([]string{"eat", "tea", "tan", "ate", "nat", "bat"})
}

func groupAnagrams(strs []string) [][]string {
	hashMap := make(map[[26]int32][]string)

	for _, s := range strs {
		var key [26]int32
		for _, c := range s {
			i := c - 'a'
			key[i]++
		}
		hashMap[key] = append(hashMap[key], s)
	}

	res := make([][]string, 0, 10)

	for _, x := range hashMap {
		res = append(res, x)
	}

	return res
}
