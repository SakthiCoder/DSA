package recursion

import "fmt"

func Example1(n int) {
	if n > 0 {
		fmt.Print(n, " ")
		Example1(n - 1)
		fmt.Print(n, " ")
	}
}

func Factorial(n int) int {

	if n == 0 || n == 1 {
		return 1
	}

	return Factorial(n-1) * n
}

func Fibonacci(n int) int {

	if n == 1 || n == 0 {
		return n
	}

	return Fibonacci(n-1) + Fibonacci(n-2)
}
