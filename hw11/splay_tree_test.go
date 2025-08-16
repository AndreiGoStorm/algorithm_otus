package hw11

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestSplayTree(t *testing.T) {
	t.Run("insert items to splay tree", func(t *testing.T) {
		// Arrange
		st := NewSplayTree()

		// Act
		st.Insert(7)
		st.Insert(9)
		st.Insert(3)
		st.Insert(6)
		st.Insert(2)
		st.Insert(1)
		st.Insert(4)
		st.Insert(5)
		st.Insert(10)
		st.Insert(8)

		// Assert
		st.Print(st.root)
		require.Equal(t, 8, st.root.key)
	})

	t.Run("insert random items to splay tree", func(t *testing.T) {
		// Arrange
		key := 500
		st := NewSplayTree()

		// Act
		for i := 1; i < 100000; i++ {
			st.Insert(rnd.Intn(100))
		}
		for i := 0; i < 5; i++ {
			st.Insert(key + i)
		}

		// Assert
		for i := 0; i < 5; i++ {
			node := st.Search(key + i)
			require.Equal(t, node.key, key+i)
		}
		require.Equal(t, 504, st.root.key)
	})

	t.Run("remove key from splay tree", func(t *testing.T) {
		// Arrange
		key := 1
		st := NewSplayTree()
		st.Insert(7)
		st.Insert(9)
		st.Insert(3)
		st.Insert(6)
		st.Insert(2)
		st.Insert(1)

		// Act
		st.Remove(key)

		// Assert
		node := st.Search(key)
		require.Nil(t, node)
		require.Equal(t, 2, st.root.key)
	})

	t.Run("remove many keys to splay tree", func(t *testing.T) {
		// Arrange
		key := 1
		st := NewSplayTree()

		// Act
		for i := key; i < 10000; i++ {
			st.Insert(i)
		}
		for i := 10000; i >= key; i-- {
			st.Remove(i)
		}

		// Assert
		require.Nil(t, st.root)
	})
}
