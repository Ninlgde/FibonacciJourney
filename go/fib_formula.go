package fibonacci

import "math"

var sqrt5 = math.Pow(5, 0.5)

func fast_pow_float64(x float64, n int) float64 {
	r := 1.0
	for n > 0 {
		if n&1 == 1 {
			r *= x
		}
		x *= x
		n >>= 1
	}
	return r
}

func fib_formula(n int) float64 {
	n++ // fib_formula(0) = 0 start with 1
	a := (1 + sqrt5) / 2
	b := (1 - sqrt5) / 2
	return (fast_pow_float64(a, n) - fast_pow_float64(b, n)) / sqrt5
}

func fib_formula_builtin(n int) float64 {
	n++ // fib_formula(0) = 0 start with 1
	a := (1 + sqrt5) / 2
	b := (1 - sqrt5) / 2
	return (math.Pow(a, float64(n)) - math.Pow(b, float64(n))) / sqrt5

}
