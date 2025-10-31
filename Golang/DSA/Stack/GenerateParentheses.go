package main

func validString(s string) bool {
	stack := make([]rune, 0)

	for a := range s {
		if s[a] == '(' {
			stack = append(stack, '(')
		}

		if s[a] == ')' {
			if len(stack) == 0 {
				return false
			} else if stack[len(stack)-1] != '(' {
				return false
			}
			stack = stack[:len(stack)-1]
		}
	}

	return true
}

func generateHelper(holder string, n int, size int, result []string) []string {

	if n == 0 {
		return result
	}

	if len(holder) == size {
		if validString(holder) {
			result = append(result, holder)
		}
		return result
	}

	leftResults := generateHelper(holder+"(", n-1, size, result)
	rightResults := generateHelper(holder+")", n, size, result)

	return append(result, append(leftResults, rightResults...)...)
}

func GenerateParenthesis(n int) []string {
	return generateHelper("(", n, n*2, make([]string, 0))
}
