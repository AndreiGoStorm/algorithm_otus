package prime

import "math"

func isPrime(n int64) bool {
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

func primes(n int64) int64 {
	count := int64(0)
	for i := int64(2); i <= n; i++ {
		if isPrime(i) {
			count++
		}
	}

	return count
}

var (
	primesCount   int64
	primesNumbers []int64
)

func isPrimeWithMemory(n int64) bool {
	sqrt := int64(math.Sqrt(float64(n)))
	for i := int64(0); primesNumbers[i] <= sqrt; i++ {
		if n%primesNumbers[i] == 0 {
			return false
		}
	}
	return true
}

func primesWithMemory(n int64) int64 {
	if n < 2 {
		return 0
	}
	primesCount = 1
	primesNumbers = make([]int64, 0, n/10)
	primesNumbers = append(primesNumbers, 2)

	for i := int64(3); i <= n; i++ {
		if isPrimeWithMemory(i) {
			primesCount++
			primesNumbers = append(primesNumbers, i)
		}
	}

	return primesCount
}

func eratosphen(n int64) int64 {
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
