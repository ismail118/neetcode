package main

func main() {

}

func isValid(s string) bool {

	hashMap := map[string]string{
		")": "(",
		"]": "[",
		"}": "{",
	}

	var stackL []string

	for _, x := range s {
		l, ok := hashMap[string(x)]
		if ok {
			if len(stackL) == 0 {
				return false
			}
			if stackL[len(stackL)-1] != l {
				return false
			} else {
				stackL = stackL[:len(stackL)-1]
			}
		} else {
			stackL = append(stackL, string(x))
		}
	}

	if len(stackL) > 0 {
		return false
	}

	return true
}
