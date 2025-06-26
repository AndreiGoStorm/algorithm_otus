package hw12

import (
	"crypto/rand"
	"math/big"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestQuadraticProbing(t *testing.T) {
	t.Run("insert items for quadratic probing", func(t *testing.T) {
		// Arrange
		ht := NewQuadraticProbing(1000, 1, 2)

		// Act
		for i := 0; i < 10000; i++ {
			rnd, _ := rand.Int(rand.Reader, big.NewInt(100))
			ht.Insert(int(rnd.Int64()))
		}

		// Assert
		require.Equal(t, ht.size, 10000)
		require.Equal(t, ht.capacity, 16015)
	})
}
