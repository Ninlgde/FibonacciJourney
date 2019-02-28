package fibonacci

import (
	"fmt"
	"testing"
)

var n = 1000

func TestFibRecursion(t *testing.T) {
	fmt.Println(fib_recursion(35))
}

func TestFibCache(t *testing.T) {
	fmt.Println(fib_cache(n))
}

func TestFibIterator(t *testing.T) {
	fmt.Println(fib_iterator(n))
}

func TestFibMatrix(t *testing.T) {
	fmt.Println(fib_matrix(n))
}

func TestFibFormulaBuiltIn(t *testing.T) {
	fmt.Println(fib_formula_builtin(n))
}

func TestFibFormula(t *testing.T) {
	fmt.Println(fib_formula(n))
}
