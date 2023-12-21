package main

import "strings"

func main() {
	strings.Replace("a b", " ", "", -1)
}

func isPalindrome(s string) bool {

	s = strings.ToLower(s)
	s = strings.Trim(s, " ")

	if s == "" {
		return true
	}

	filteredS := make([]int32, 0, 10)
	for _, x := range s {
		if x > 96 && x < 123 || x > 47 && x < 58 {
			filteredS = append(filteredS, x)
		}
	}

	n := len(filteredS) - 1
	for i := 0; i < len(filteredS)/2; i++ {
		if filteredS[i] != filteredS[n-i] {
			return false
		}
	}

	return true
}
