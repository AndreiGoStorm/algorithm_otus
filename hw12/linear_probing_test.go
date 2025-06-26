package hw12

import (
	"crypto/rand"
	"math/big"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestGetLinearProbing(t *testing.T) {
	t.Run("get empty item for not existed key", func(t *testing.T) {
		// Arrange
		ht := NewLinearProbing(10)

		// Act
		item := ht.Get(1)

		// Assert
		require.Nil(t, item)
	})

	t.Run("get item for existed key", func(t *testing.T) {
		// Arrange
		ht := NewLinearProbing(10)
		key := 265
		ht.Insert(key)

		// Act
		item := ht.Get(key)

		// Assert
		require.NotNil(t, item)
		require.Equal(t, item.key, key)
		require.Equal(t, item.deleted, false)
	})
}

func TestRemoveLinearProbing(t *testing.T) {
	t.Run("remove existed item", func(t *testing.T) {
		// Arrange
		ht := NewLinearProbing(20)
		key := 169
		ht.Insert(key)

		// Act
		ht.Remove(key)

		// Assert
		item := ht.Get(key)
		require.Nil(t, item)
		require.Equal(t, ht.size, 0)
	})

	t.Run("remove item in rehashed table", func(t *testing.T) {
		// Arrange
		capacity := 10
		ht := NewLinearProbing(capacity)
		key := 100
		for i := 0; i < capacity; i++ {
			ht.Insert(i + key)
		}

		// Act
		ht.Remove(key)

		// Assert
		item := ht.Get(key)
		require.Nil(t, item)
		require.Equal(t, ht.size, 9)
		require.Equal(t, ht.capacity, 21)
	})
}

func TestInsertLinearProbing(t *testing.T) {
	t.Run("insert item to hashtable", func(t *testing.T) {
		// Arrange
		ht := NewLinearProbing(20)
		key := 125

		// Act
		ht.Insert(key)
		item := ht.Get(key)

		// Assert
		require.Equal(t, item.key, key)
		require.Equal(t, item.deleted, false)
	})

	t.Run("insert items with rehash", func(t *testing.T) {
		// Arrange
		capacity := 10
		ht := NewLinearProbing(capacity)

		// Act
		for i := 0; i < 1000; i++ {
			rnd, _ := rand.Int(rand.Reader, big.NewInt(100))
			ht.Insert(int(rnd.Int64()))
		}

		// Assert
		require.Equal(t, ht.size, 1000)
		require.Equal(t, ht.capacity, 2815)
	})
}
