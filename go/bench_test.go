package fibonacci

import "testing"

func BenchmarkFibRecursion(t *testing.B) {
	for i := 0; i < t.N; i++ {
		fib_recursion(35)
	}
}

var bn = 1000

func BenchmarkFibCache(t *testing.B) {
	for i := 0; i < t.N; i++ {
		fib_cache(bn)
	}
}

func BenchmarkFibIterator(t *testing.B) {
	for i := 0; i < t.N; i++ {
		fib_iterator(bn)
	}
}

func BenchmarkFibMatrix(t *testing.B) {
	for i := 0; i < t.N; i++ {
		fib_matrix(bn)
	}
}

func BenchmarkFibFormulaBuildIn(t *testing.B) {
	for i := 0; i < t.N; i++ {
		fib_formula_buildin(bn)
	}
}

func BenchmarkFibFormula(t *testing.B) {
	for i := 0; i < t.N; i++ {
		fib_formula(bn)
	}
}
