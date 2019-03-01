package fibonacci

import (
	"fmt"
	"testing"
)

var n = 10

func TestFibRecursion(t *testing.T) {
	fmt.Println(fib_recursion(10))
}

func TestFibCache(t *testing.T) {
	fmt.Println(fib_cache(n))
}

func TestFibIterate(t *testing.T) {
	fmt.Println(fib_iterate(n))
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
