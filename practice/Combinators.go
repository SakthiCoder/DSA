package practice

import "dsa/recursion"

func NCRCombinators(n, r int) int {
	return recursion.Factorial(n) / (recursion.Factorial(r) * recursion.Factorial((n - r)))
}
