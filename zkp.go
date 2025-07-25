package main

import "github.com/shopspring/decimal"

type ZKP struct {
	p     decimal.Decimal
	q     decimal.Decimal
	alpha decimal.Decimal
	beta  decimal.Decimal
}

func NewZKP(p, q, alpha, beta decimal.Decimal) *ZKP {
	return &ZKP{
		p:     p,
		q:     q,
		alpha: alpha,
		beta:  beta,
	}
}

// Output: n^exp mod p
func (zkp *ZKP) Exponentiate(n decimal.Decimal, exp decimal.Decimal) decimal.Decimal {
	return n.Pow(exp).Mod(zkp.p)

}

// Output:
// if k >= c * x then s = k - c * x mod q
// if k < c * x then s = (k - c * x) + q mod q
func (zkp *ZKP) Solve(k decimal.Decimal, c decimal.Decimal, x decimal.Decimal) decimal.Decimal {
	cx := c.Mul(x)
	if k.GreaterThanOrEqual(cx) {
		return k.Sub(cx).Mod(zkp.q)
	}
	diff := k.Sub(cx)
	result := diff.Mod(zkp.q)
	if result.Sign() < 0 {
		result = result.Add(zkp.q)
	}
	return result
}

// Output:
// r1 = alpha^s * y1^c
// r2 = beta^s * y2^c
func (zkp *ZKP) Verify(r1 decimal.Decimal, r2 decimal.Decimal, y1 decimal.Decimal, y2 decimal.Decimal, s decimal.Decimal, c decimal.Decimal) bool {
	cond1 := r1.Cmp(zkp.alpha.Pow(s).Mul(y1.Pow(c).Mod(zkp.p)).Mod(zkp.p))
	cond2 := r2.Cmp(zkp.beta.Pow(s).Mul(y2.Pow(c).Mod(zkp.p)).Mod(zkp.p))
	return cond1 == 0 && cond2 == 0
}
