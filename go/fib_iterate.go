package fibonacci

// To iterate is human,to recurse divine.
func fib_iterate(n int) float64 {
	a, b := 1.0, 0.0
	for ; n > 1; n-- {
		a, b = a+b, a // => t=a;a=a+b;b=t
	}
	return a
}
