#include <stdio.h>

long double cache[2600] = {0};

long double fib_recursion_cache(int n)
{
	long double result = 0.0f;
	if (n <= 2) {
		return 1;
	} else if (cache[n] != 0.0f) {
		return cache[n];
	} else {
		result = fib_recursion_cache(n-1) + fib_recursion_cache(n-2);
		cache[n] = result;
		return result;
	}
}

int main(int argc, char const *argv[])
{
	printf("%Lf\n", fib_recursion_cache(2561)); // 2561 is the biggest number
	return 0;
}