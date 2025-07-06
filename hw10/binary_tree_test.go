package hw10

import (
	"crypto/rand"
	"math/big"
	"strings"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestBinaryTree(t *testing.T) {
	t.Run("insert items to binary tree", func(t *testing.T) {
		// Arrange
		bt := NewBinaryTree()

		// Act
		bt.Insert(7)
		bt.Insert(9)
		bt.Insert(3)
		bt.Insert(6)
		bt.Insert(2)
		bt.Insert(1)
		bt.Insert(4)
		bt.Insert(5)
		bt.Insert(10)
		bt.Insert(8)

		// Assert
		str := bt.Sort(bt.root)
		require.Equal(t, strings.Trim(str, " "), "1 2 3 4 5 6 7 8 9 10")
	})

	t.Run("insert random items to binary tree", func(t *testing.T) {
		// Arrange
		key := 325
		bt := NewBinaryTree()

		// Act
		for i := 1; i < 10000; i++ {
			rnd, _ := rand.Int(rand.Reader, big.NewInt(1000))
			bt.Insert(int(rnd.Int64()))
		}
		bt.Insert(key)

		// Assert
		node := bt.Search(key)
		require.Equal(t, node.key, key)
	})

	t.Run("search parents in binary tree", func(t *testing.T) {
		// Arrange
		key := 9
		bt := NewBinaryTree()

		// Act
		bt.Insert(7)
		bt.Insert(9)
		bt.Insert(3)
		bt.Insert(6)

		// Assert
		node := bt.Search(key)
		parent := bt.SearchParent(node, bt.root)
		require.Equal(t, node.key, parent.right.key)
	})

	t.Run("remove key from binary tree", func(t *testing.T) {
		// Arrange
		key := 10
		bt := NewBinaryTree()
		bt.Insert(7)
		bt.Insert(9)
		bt.Insert(3)
		bt.Insert(6)
		bt.Insert(2)
		bt.Insert(1)
		bt.Insert(4)
		bt.Insert(5)
		bt.Insert(key)
		bt.Insert(8)

		// Act
		bt.Remove(key)

		// Assert
		node := bt.Search(key)
		require.Nil(t, node)
	})

	t.Run("remove many keys to binary tree", func(t *testing.T) {
		// Arrange
		key := 1
		bt := NewBinaryTree()

		// Act
		for i := key; i < 10000; i++ {
			bt.Insert(i)
		}
		for i := 10000; i >= key; i-- {
			bt.Remove(i)
		}

		// Assert
		require.Nil(t, bt.root)
	})
}
