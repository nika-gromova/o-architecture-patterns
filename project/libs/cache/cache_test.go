package cache

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestCache(t *testing.T) {
	t.Run("should store values correctly", func(t *testing.T) {
		c := New()

		c.Set("test", 1000)

		value, ok := c.Get("test")
		require.True(t, ok)
		require.Equal(t, 1000, value)
	})
	t.Run("should return false if value not found", func(t *testing.T) {
		c := New()

		c.Set("test", 1000)

		value, ok := c.Get("test-1000")
		require.False(t, ok)
		require.Empty(t, value)
	})
}
