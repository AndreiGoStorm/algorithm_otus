package power

import "math/big"

// алгоритм возведения в степень O(N/2+LogN) = O(N).
func powerN(a float64, n int64) float64 {
	if n == 0 {
		return 1
	}
	if n == 1 {
		return a
	}
	x := powerN(a, n/2)
	if n%2 == 0 {
		return x * x
	}

	return a * x * x
}

func NBigFloat(a *big.Float, n int64) *big.Float {
	if n == 0 {
		return big.NewFloat(1)
	}
	if n == 1 {
		return new(big.Float).Set(a)
	}
	x := NBigFloat(a, n/2)
	xx := new(big.Float).Mul(x, x)
	if n%2 == 0 {
		return xx
	}

	return new(big.Float).Mul(a, xx)
}
