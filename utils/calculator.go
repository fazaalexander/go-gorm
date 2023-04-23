package utils

func sum(a, b int) int {
	return a + b
}

func sub(a, b int) int {
	return a - b
}

func div(a, b int) int {
	if b == 0 {
		return 0
	}
	return a / b
}

func multiple(a, b int) int {
	return a * b
}
