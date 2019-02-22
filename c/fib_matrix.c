//
//  main.c
//  fab
//
//  Created by Malik Ma on 2019/2/22.
//  Copyright © 2019 Malik Ma. All rights reserved.
//

#include <stdio.h>
#include <string.h>

#define MAT_SIZE 4*sizeof(long double)

void matmul(long double a[2][2], long double b[2][2], long double mul[2][2])
{
    long double result[2][2];
    result[0][0] = a[0][0] * b[0][0] + a[0][1] * b[1][0];
    result[0][1] = a[0][0] * b[0][1] + a[0][1] * b[1][1];
    result[1][0] = a[1][0] * b[0][0] + a[1][1] * b[1][0];
    result[1][1] = a[1][0] * b[0][1] + a[1][1] * b[1][1];
    memcpy(mul, result, MAT_SIZE);
}

void fast_pow_mat(long double x[2][2], int n, long double result[2][2])
{
    long double a[2][2] = {{1, 0}, {0, 1}};
    long double b[2][2];
    memcpy(b, x, MAT_SIZE);
    
    while(n > 0) {
        if (n & 1)
            matmul(a, b, a);
        matmul(b, b, b);
        n >>= 1;
    }
    memcpy(result, a, MAT_SIZE);
}


int main(int argc, const char * argv[]) {
    long double a[2][2] = {{1, 1}, {1, 0}};
    long double result[2][2];
    fast_pow_mat(a, 10000, result);
    printf("%Lf\n", result[0][1]);
    return 0;
}
