package fibonacci

type mat22 struct {
	a float64 // 0 0
	b float64 // 0 1
	c float64 // 1 0
	d float64 // 1 1
}

func (a *mat22) mul(b *mat22) *mat22 {
	//a[0][0] * b[0][0] + a[0][1] * b[1][0];
	//a[0][0] * b[0][1] + a[0][1] * b[1][1];
	//a[1][0] * b[0][0] + a[1][1] * b[1][0];
	//a[1][0] * b[0][1] + a[1][1] * b[1][1];
	return &mat22{
		a.a*b.a + a.b*b.c,
		a.a*b.b + a.b*b.d,
		a.c*b.a + a.d*b.c,
		a.c*b.b + a.d*b.d,
	}
}

func fast_pow_mat22(x *mat22, n int) *mat22 {
	I := &mat22{1, 0, 0, 1}
	for n > 0 {
		if n&1 == 1 {
			I = I.mul(x)
		}
		x = x.mul(x)
		n >>= 1
	}
	return I
}

func fib_matrix(n int) float64 {
	x := &mat22{1, 1, 1, 0}
	r := fast_pow_mat22(x, n)
	return r.a
}
