#include<stdio.h>

double fib_recursion(double n)
{
	if (n <= 2) {
		return 1;
	} else {
		return fib_recursion(n-1) + fib_recursion(n-2);
	}
}

int main(int argc, char const *argv[])
{
	printf("%f\n", fib_recursion(50));
	return 0;
}