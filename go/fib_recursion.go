package fibonacci

func fib_recursion(n int) float64 {
	if n <= 1 {
		return float64(n)
	}
	return fib_recursion(n-1) + fib_recursion(n-2)
}
