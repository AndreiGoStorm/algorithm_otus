package hw21

func gcd(m, n int64) int64 {
	if m == 0 || n == 0 {
		return 1
	}
	if m == n {
		return m
	}
	if m&1 == 0 && n&1 == 0 {
		return 2 * gcd(m>>1, n>>1)
	}
	if m&1 == 1 && n&1 == 1 {
		if m > n {
			return gcd(n, (m-n)>>1)
		} else {
			return gcd(m, (n-m)>>1)
		}
	}
	if m&1 == 1 {
		return gcd(m, n>>1)
	}
	if n&1 == 1 {
		return gcd(m>>1, n)
	}
	return 1
}

func garland(tree [][]int64) int64 {
	start := len(tree) - 2
	for i := start; i >= 0; i-- {
		for j := 0; j < len(tree[i]); j++ {
			tree[i][j] += max(tree[i+1][j], tree[i+1][j+1])
		}
	}
	return tree[0][0]
}

func fiveAndEighth(N, x5, x55, x8, x88 int64) int64 {
	if N < 2 {
		return x5 + x55 + x8 + x88
	}

	return fiveAndEighth(N-1, x8+x88, x5, x5+x55, x8)
}

func bigIsland(matrix [][]int64) (count int64) {
	N := int64(len(matrix))
	var walk func(x, y int64)
	walk = func(x, y int64) {
		if x < 0 || x >= N || y < 0 || y >= N {
			return
		}
		if matrix[x][y] == 0 {
			return
		}
		matrix[x][y] = 0
		walk(x, y-1)
		walk(x, y+1)
		walk(x-1, y)
		walk(x+1, y)
	}

	for i := int64(0); i < N; i++ {
		for j := int64(0); j < N; j++ {
			if matrix[i][j] == 1 {
				walk(i, j)
				count++
			}
		}
	}
	return count
}
