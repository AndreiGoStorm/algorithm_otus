package prime

import "math"

func IsPrime(n int64) bool {
	if n == 2 {
		return true
	}
	if n%2 == 0 {
		return false
	}

	sqrt := int64(math.Sqrt(float64(n)))
	for i := int64(3); i <= sqrt; i += 2 {
		if n%i == 0 {
			return false
		}
	}
	return true
}

func Primes(n int64) int64 {
	count := int64(0)
	for i := int64(2); i <= n; i++ {
		if IsPrime(i) {
			count++
		}
	}

	return count
}

var (
	count  int64
	primes []int64
)

func IsPrimeWithMemory(n int64) bool {
	sqrt := int64(math.Sqrt(float64(n)))
	for i := int64(0); primes[i] <= sqrt; i++ {
		if n%primes[i] == 0 {
			return false
		}
	}
	return true
}

func PrimesWithMemory(n int64) int64 {
	if n < 2 {
		return 0
	}
	count = 1
	primes = make([]int64, 0, n/10)
	primes = append(primes, 2)

	for i := int64(3); i <= n; i++ {
		if IsPrimeWithMemory(i) {
			count++
			primes = append(primes, i)
		}
	}

	return count
}

func Eratosphen(n int64) int64 {
	prime := make([]bool, n+1)
	cnt := int64(0)
	for i := int64(2); i <= n; i++ {
		if !prime[i] {
			cnt++
			for j := i * i; j <= n; j += i {
				prime[j] = true
			}
		}
	}
	return cnt
}
