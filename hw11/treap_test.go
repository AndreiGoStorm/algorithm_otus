package hw11

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestTreap(t *testing.T) {
	t.Run("insert items to treap", func(t *testing.T) {
		// Arrange
		var root *Treap

		// Act
		root = root.Insert(7)
		root = root.Insert(9)
		root = root.Insert(3)
		root = root.Insert(6)
		root = root.Insert(2)
		root = root.Insert(1)
		root = root.Insert(4)
		root = root.Insert(5)
		root = root.Insert(10)
		root = root.Insert(8)

		// Assert
		root.Print(root)
		require.Equal(t, 10, root.Count(root))
	})

	t.Run("insert random items to treap", func(t *testing.T) {
		// Arrange
		key := 1000
		var root *Treap

		// Act
		for i := 1; i < 100000; i++ {
			root = root.Insert(rnd.Intn(100))
		}
		for i := 0; i < 5; i++ {
			root = root.Insert(key + i)
		}

		// Assert
		for i := 0; i < 5; i++ {
			treap := root.Search(key + i)
			require.Equal(t, treap.x, key+i)
		}
		require.Equal(t, 100004, root.Count(root))
	})

	t.Run("remove key from treap", func(t *testing.T) {
		// Arrange
		key := 5
		var root *Treap
		root = root.Insert(2)
		root = root.Insert(1)
		root = root.Insert(5)
		root = root.Insert(10)

		// Act
		root = root.Remove(key)

		// Assert
		treap := root.Search(key)
		require.Nil(t, treap)
		require.Equal(t, 3, root.Count(root))
	})

	t.Run("remove many keys to treap", func(t *testing.T) {
		// Arrange
		key := 500000
		var root *Treap
		root = root.Insert(key)

		// Act
		for i := 0; i < 100000; i++ {
			root = root.Insert(i)
		}
		for i := 100000; i >= 0; i-- {
			root = root.Remove(i)
		}

		// Assert
		treap := root.Search(key)
		require.Equal(t, key, treap.x)
		require.Equal(t, 1, root.Count(root))
	})
}
