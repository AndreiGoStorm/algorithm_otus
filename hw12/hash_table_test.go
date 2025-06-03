package hw12

import (
	"crypto/rand"
	"math/big"
	"strconv"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestGetItemsHashTable(t *testing.T) {
	t.Run("get empty item for not existed key", func(t *testing.T) {
		// Arrange
		ht := NewHashTable()

		// Act
		item := ht.Get("cat")

		// Assert
		require.Nil(t, item)
	})

	t.Run("get item for existed key", func(t *testing.T) {
		// Arrange
		ht := NewHashTable()
		ht.Insert("cat", "black")

		// Act
		item := ht.Get("cat")

		// Assert
		require.NotNil(t, item)
		require.Equal(t, item.value.(string), "black")
	})

	t.Run("get item with updated value", func(t *testing.T) {
		// Arrange
		ht := NewHashTable()
		ht.Insert("cat", "grey")
		ht.Insert("dog", "black")
		ht.Insert("dog", "white")

		// Act
		item := ht.Get("dog")

		// Assert
		require.NotNil(t, item)
		require.Equal(t, item.value.(string), "white")
	})
}

func TestRemoveItemsHashTable(t *testing.T) {
	t.Run("remove not existed item", func(t *testing.T) {
		// Arrange
		ht := NewHashTable()
		ht.Insert("cat", "grey")

		// Act
		ht.Remove("cat")
		item := ht.Get("cat")

		// Assert
		require.Nil(t, item)
	})

	t.Run("remove existed item", func(t *testing.T) {
		// Arrange
		ht := NewHashTable()
		ht.Insert("cat", "grey")
		ht.Insert("dog", "black")

		// Act
		ht.Remove("dog")
		item := ht.Get("dog")

		// Assert
		require.Nil(t, item)
	})
}

func TestInsertItemsHashTable(t *testing.T) {
	t.Run("insert item to hashtable", func(t *testing.T) {
		// Arrange
		ht := NewHashTable()

		// Act
		ht.Insert("cat", "grey")
		item := ht.Get("cat")

		// Assert
		require.NotNil(t, item)
		require.Equal(t, item.key, "cat")
		require.Equal(t, item.value.(string), "grey")
	})

	t.Run("insert 10000 items to hashtable", func(t *testing.T) {
		// Arrange
		ht := NewHashTable()
		for i := 0; i < 10000; i++ {
			rnd, _ := rand.Int(rand.Reader, big.NewInt(100))
			ht.Insert("cat"+strconv.Itoa(i), rnd)
		}

		// Act
		item := ht.Get("cat0")

		// Assert
		require.NotNil(t, item)
		require.Equal(t, item.key, "cat0")
	})
}
