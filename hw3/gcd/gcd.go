package gcd

import "math/big"

func gcd(m, n *big.Int) *big.Int {
	if n.Cmp(big.NewInt(0)) == 0 {
		return m
	}

	if m.Cmp(big.NewInt(0)) == 0 {
		return n
	}

	if m.Cmp(n) == 0 {
		return m
	}

	if m.Cmp(n) > 0 {
		return gcd(new(big.Int).Rem(m, n), n)
	}

	return gcd(new(big.Int).Rem(n, m), m)
}
