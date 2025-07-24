package main

import "github.com/shopspring/decimal"

// Output: n^exp mod p
func Exponentiate(n decimal.Decimal, exp decimal.Decimal, p decimal.Decimal) decimal.Decimal {
	return n.Pow(exp).Mod(p)

}

// Output:
// if k >= c * x then s = k - c * x mod q
// if k < c * x then s = (k - c * x) + q mod q
func Solve(k decimal.Decimal, c decimal.Decimal, x decimal.Decimal, q decimal.Decimal) decimal.Decimal {
	cx := c.Mul(x)
	if k.GreaterThanOrEqual(cx) {
		return k.Sub(cx).Mod(q)
	}
	diff := k.Sub(cx)
	result := diff.Mod(q)
	if result.Sign() < 0 {
		result = result.Add(q)
	}
	return result
}

// Output:
// r1 = alpha^s * y1^c
// r2 = beta^s * y2^c
func Verify(r1 decimal.Decimal, r2 decimal.Decimal, y1 decimal.Decimal, y2 decimal.Decimal, alpha decimal.Decimal, beta decimal.Decimal, s decimal.Decimal, c decimal.Decimal, p decimal.Decimal) bool {
	cond1 := r1.Cmp(alpha.Pow(s).Mul(y1.Pow(c).Mod(p)).Mod(p))
	cond2 := r2.Cmp(beta.Pow(s).Mul(y2.Pow(c).Mod(p)).Mod(p))
	return cond1 == 0 && cond2 == 0
}
