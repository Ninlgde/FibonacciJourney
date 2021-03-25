def fib_recursion(n):
    if n <= 1:
        return float(n)
    return fib_recursion(n-1) + fib_recursion(n-2)


def fib_recursion_cache(cache, n):
    if n <= 1:
        cache[n] = float(n)
        return cache[n]
    if cache.get(n, None):
        return cache[n]
    cache[n] = fib_recursion_cache(cache, n-1) + fib_recursion_cache(cache, n-2)
    return cache[n]


cache = {}
def fib_cache(n):
    global cache
    return fib_recursion_cache(cache, n)


def fib_iterate(n):
    a, b = 1.0, 0.0
    for _ in range(n-1):
        a, b = a + b, a
    return a


class Matrix(object):
    def __init__(self, a, b, c, d):
        self.a = a
        self.b = b
        self.c = c
        self.d = d
        
    def mul(self, other):
        return Matrix(
            self.a*other.a + self.b*other.c,
            self.a*other.b + self.b*other.d,
            self.c*other.a + self.d*other.c,
            self.c*other.b + self.d*other.d,
        )


def fast_pow_mat(x, n):
    I = Matrix(1, 0, 0, 1)
    while n > 0:
        if n & 1 == 1:
            I = I.mul(x)
        x = x.mul(x)
        n >>= 1 # 每次缩减一半，所以时间复杂度是log(n)
    return I

def fib_matrix(n):
    x = Matrix(1.0, 1.0, 1.0, 0.0)
    r = fast_pow_mat(x, n)

    return r.b


# 根号5提前算好，避免重复计算
sqrt5 = 5 ** 0.5
a = (1 + sqrt5) / 2
b = (1 - sqrt5) / 2

def fast_pow(x, n):
    r = 1.0
    while n > 0:
        if n & 1 == 1:
            r *= x
        x *= x
        n >>= 1
    return r

def fib_formula(n):
    return (fast_pow(a, n) - fast_pow(b, n)) / sqrt5

def fib_formula_builtin1(n):
    return (a ** n - b ** n) / sqrt5


def fib_formula_builtin2(n):
    return (pow(a, n) - pow(b, n)) / sqrt5
