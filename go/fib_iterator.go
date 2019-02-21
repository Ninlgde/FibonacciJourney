package fibonacci

func fib_iterator(n int) float64 {
	a, b := 1.0, 1.0
	for ; n > 1; n-- {
		a, b = a+b, a
	}
	return a
}
