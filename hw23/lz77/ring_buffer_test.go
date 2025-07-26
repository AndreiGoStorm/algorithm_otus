package lz77

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestGetItemsHashTable(t *testing.T) {
	t.Run("get the oldest element", func(t *testing.T) {
		// Arrange
		rb := NewRingBuffer(5)
		rb.Push('A')
		rb.Push('B')
		rb.Push('C')
		rb.Push('D')
		rb.Push('E')
		rb.Push('F')

		// Act
		b, err := rb.GetOldByte()

		// Assert
		require.NoError(t, err)
		require.Equal(t, uint8('B'), b)
		require.Equal(t, 5, rb.len)
	})

	t.Run("push slice of bytes", func(t *testing.T) {
		// Arrange
		rb := NewRingBuffer(5)

		// Act
		rb.PushBytes([]byte{'A', 'B', 'C', 'D', 'E', 'F', 'G'})

		// Assert
		b, err := rb.GetOldByte()
		require.NoError(t, err)
		require.Equal(t, uint8('C'), b)
	})

	t.Run("get bytes from offset", func(t *testing.T) {
		// Arrange
		rb := NewRingBuffer(6)
		rb.PushBytes([]byte{'A', 'B', 'C', 'D', 'E', 'F'})

		// Act
		bytes, err := rb.GetBytesFromOffset(6, 3)

		// Assert
		require.NoError(t, err)
		require.Equal(t, []byte{'A', 'B', 'C'}, bytes)
		b, err := rb.GetOldByte()
		require.NoError(t, err)
		require.Equal(t, uint8('A'), b)
	})
}
