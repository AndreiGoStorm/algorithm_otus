package fibo

import (
	"math"
	"math/big"

	"algorithm_otus/hw3/power"
)

func fiboGoldenRatio(k int64) *big.Int {
	f := big.NewInt(0)
	if k <= 1 {
		return f.SetInt64(k)
	}

	sqrt := math.Sqrt(float64(5))
	sqrtFloat := new(big.Float).SetFloat64(sqrt)

	fi := (1.0 + sqrt) / 2.0
	fiFloat := new(big.Float).SetFloat64(fi)

	x := power.NBigFloat(fiFloat, k)
	div := new(big.Float).Quo(x, sqrtFloat)
	result := new(big.Float).Add(div, new(big.Float).SetFloat64(0.5))

	result.Int(f)
	return f
}

func fiboMatrix(k int64) *big.Int {
	if k == 0 {
		return big.NewInt(0)
	}

	init := [2][2]*big.Int{
		{big.NewInt(1), big.NewInt(1)},
		{big.NewInt(1), big.NewInt(0)},
	}

	matrix := matrixPow(init, k-1)
	return matrix[0][0]
}

func matrixPow(a [2][2]*big.Int, k int64) [2][2]*big.Int {
	if k <= 1 {
		return a
	}

	x := matrixPow(a, k/2)
	if k%2 == 0 {
		return matrixMul(x, x)
	}
	return matrixMul(a, matrixMul(x, x))
}

func matrixMul(a, b [2][2]*big.Int) [2][2]*big.Int {
	var res [2][2]*big.Int
	res[0][0] = new(big.Int).Add(
		new(big.Int).Mul(a[0][0], b[0][0]),
		new(big.Int).Mul(a[0][1], b[1][0]),
	)
	res[0][1] = new(big.Int).Add(
		new(big.Int).Mul(a[0][0], b[0][1]),
		new(big.Int).Mul(a[0][1], b[1][1]),
	)
	res[1][0] = new(big.Int).Set(res[0][1])
	res[1][1] = new(big.Int).Add(
		new(big.Int).Mul(a[1][0], b[0][1]),
		new(big.Int).Mul(a[1][1], b[1][1]),
	)
	return res
}
