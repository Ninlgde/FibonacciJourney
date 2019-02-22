#include <stdio.h>

long double fib_iterator(int n)
{
	if (n <= 2) {
		return 1;
	}
	long double a = 1.0f, b = 1.0f, temp;
	int i;
	for (i = 2; i < n; ++i)
	{
		temp = a + b;
		b = a;
		a = temp;
	}

	return a;
}

int main(int argc, char const *argv[])
{
	printf("%Lf\n", fib_iterator(10000)); // 2561 is the biggest number
	return 0;
}