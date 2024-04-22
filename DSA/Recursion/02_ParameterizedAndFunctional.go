package main

func summationParameterizedHelper(i int, sum int) int {
	if i < 1 {
		return sum
	}

	return summationParameterizedHelper(i-1, sum+i)
}

func summationParameterized(n int) int {
	if n == 0 || n == 1 {
		return n
	}

	return summationParameterizedHelper(n, 0)
}

func summationFunctional(n int) int {
	if n == 0 || n == 1 {
		return n
	}

	return n + summationFunctional(n-1)
}
