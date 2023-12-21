package main

import "strconv"

func main() {

}

func evalRPN(tokens []string) int {

	if len(tokens) == 1 {
		num, err := strconv.Atoi(tokens[0])
		if err != nil {
			panic(err)
		}
		return num
	}
	validOpr := map[string]struct{}{
		"+": {},
		"-": {},
		"*": {},
		"/": {},
	}

	calculator := func(a, b int, operator string) int {
		switch operator {
		case "+":
			return a + b
		case "-":
			return a - b
		case "*":
			return a * b
		case "/":
			return a / b
		default:
			return 0
		}
	}

	stack := make([]int, 0, 10)
	var result int

	for _, token := range tokens {
		if _, ok := validOpr[token]; ok {
			result = calculator(stack[len(stack)-2], stack[len(stack)-1], token)
			stack = stack[:len(stack)-2]
			stack = append(stack, result)
		} else {
			num, err := strconv.Atoi(token)
			if err != nil {
				panic(err)
			}
			stack = append(stack, num)
		}
	}

	return result
}
