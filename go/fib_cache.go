package fibonacci

var cache = make([]float64, 2000)

func fib_recursion_cache(n int) float64 {
	if n <= 1 {
		cache[n] = float64(n)
		return float64(n)
	}
	if cache[n] != 0 {
		return cache[n]
	}
	// 这里先计算了f(n-2)，是因为有些语言有最大堆栈限制，先计算n-2可以使深度提高一倍
	// 甚至可以略微提高效率
	cache[n] = fib_recursion_cache(n-2) + fib_recursion_cache(n-1)
	return cache[n]
}

func fib_cache(n int) float64 {
	cache = make([]float64, 2000)
	return fib_recursion_cache(n)
}
