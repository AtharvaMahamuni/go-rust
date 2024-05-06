package main

import "strconv"

func EvalRPN(tokens []string) int {

	stack := make([]int, 0)
	for _, token := range tokens {

		if token != "+" && token != "-" && token != "*" && token != "/" {
			num, _ := strconv.Atoi(token)
			stack = append(stack, num)
		} else {
			top := stack[len(stack)-1]
			topPrev := stack[len(stack)-2]
			stack = stack[:len(stack)-2]

			var exp int
			if token == "+" {
				exp = topPrev + top
			} else if token == "-" {
				exp = topPrev - top
			} else if token == "*" {
				exp = topPrev * top
			} else {
				exp = topPrev / top
			}

			stack = append(stack, exp)
		}

	}

	return stack[len(stack)-1]
}
