package fibonacci

import "testing"

func BenchmarkFibRecursion(t *testing.B) {
	for i := 0; i < t.N; i++ {
		fib_recursion(10)
	}
}

var bn = 10

func BenchmarkFibCache(t *testing.B) {
	for i := 0; i < t.N; i++ {
		fib_cache(bn)
	}
}

func BenchmarkFibIterate(t *testing.B) {
	for i := 0; i < t.N; i++ {
		fib_iterate(bn)
	}
}

func BenchmarkFibMatrix(t *testing.B) {
	for i := 0; i < t.N; i++ {
		fib_matrix(bn)
	}
}

func BenchmarkFibFormulaBuiltIn(t *testing.B) {
	for i := 0; i < t.N; i++ {
		fib_formula_builtin(bn)
	}
}

func BenchmarkFibFormula(t *testing.B) {
	for i := 0; i < t.N; i++ {
		fib_formula(bn)
	}
}
