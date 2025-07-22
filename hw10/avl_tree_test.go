package hw10

import (
	"crypto/rand"
	"math/big"
	"strings"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestAVLTree(t *testing.T) {
	t.Run("insert items to avl tree", func(t *testing.T) {
		// Arrange
		avl := NewAVLTree()

		// Act
		avl.Insert(7)
		avl.Insert(9)
		avl.Insert(3)
		avl.Insert(6)
		avl.Insert(2)
		avl.Insert(1)
		avl.Insert(4)
		avl.Insert(5)
		avl.Insert(10)
		avl.Insert(8)

		// Assert
		str := avl.Sort(avl.root)
		require.Equal(t, strings.Trim(str, " "), "1 2 3 4 5 6 7 8 9 10")
	})

	t.Run("insert random items to avl tree", func(t *testing.T) {
		// Arrange
		key := 400
		avl := NewAVLTree()

		// Act
		for i := 1; i < 100000; i++ {
			rnd, _ := rand.Int(rand.Reader, big.NewInt(1000))
			avl.Insert(int(rnd.Int64()))
		}
		for i := 0; i < 5; i++ {
			avl.Insert(key + i)
		}

		// Assert
		for i := 0; i < 5; i++ {
			node := avl.Search(key + i)
			require.Equal(t, node.key, key+i)
		}
	})

	t.Run("remove key from avl tree", func(t *testing.T) {
		// Arrange
		key := 7
		avl := NewAVLTree()
		avl.Insert(7)
		avl.Insert(9)
		avl.Insert(3)
		avl.Insert(6)
		avl.Insert(2)

		// Act
		avl.Remove(key)

		// Assert
		node := avl.Search(key)
		require.Nil(t, node)
	})

	t.Run("remove many keys to avl tree", func(t *testing.T) {
		// Arrange
		key := 1
		avl := NewAVLTree()

		// Act
		for i := key; i < 10000; i++ {
			avl.Insert(i)
		}
		for i := 10000; i >= key; i-- {
			avl.Remove(i)
		}

		// Assert
		require.Nil(t, avl.root)
	})
}
