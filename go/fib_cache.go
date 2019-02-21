package fibonacci

var cache = make([]float64, 2000)

func fib_recursion_cache(n int) float64 {
	if n <= 1 {
		cache[n] = 1
		return 1
	}
	if cache[n] != 0 {
		return cache[n]
	}
	cache[n] = fib_recursion_cache(n-1) + fib_recursion_cache(n-2)
	return cache[n]
}

func fib_cache(n int) float64 {
	cache = make([]float64, 2000)
	return fib_recursion_cache(n)
}
