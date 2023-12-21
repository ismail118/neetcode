package main

import "strings"

func main() {

}

func isAnagram(s string, t string) bool {

	if len(s) != len(t) {
		return false
	}

	makeHashMap := func(x string) map[string]int {
		hashMap := make(map[string]int)

		letters := strings.Split(x, "")
		for _, letter := range letters {
			_, ok := hashMap[letter]
			if ok {
				hashMap[letter] = hashMap[letter] + 1
			} else {
				hashMap[letter] = 1
			}
		}

		return hashMap
	}

	hashMapS := makeHashMap(s)
	hashMapT := makeHashMap(t)

	for k, x := range hashMapS {
		y := hashMapT[k]
		if x != y {
			return false
		}
	}

	return true
}
