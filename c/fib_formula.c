//
//  main.c
//  fab
//
//  Created by Malik Ma on 2019/2/22.
//  Copyright © 2019 Malik Ma. All rights reserved.
//

#include <stdio.h>
#include <string.h>
#include <math.h>

#define SQRT5 sqrt(5)

long double fast_pow(long double x, int n)
{
    long double a = 1.0f;
    
    while(n > 0) {
        if (n & 1)
            a *= x;
        x *= x;
        n >>= 1;
    }
    return a;
}


int main(int argc, const char * argv[]) {
    long double a = (1 + SQRT5) / 2;
    long double b = (1 - SQRT5) / 2;
    long double result = (fast_pow(a, 10000) - fast_pow(b, 10000)) / SQRT5;
    printf("%Lf\n", result);
    return 0;
}
