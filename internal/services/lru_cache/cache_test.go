package lruCache

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestCache(t *testing.T) {
	t.Run("empty cache", func(t *testing.T) {
		c := NewCache(10)

		_, ok := c.Get("aaa")
		require.False(t, ok)

		_, ok = c.Get("bbb")
		require.False(t, ok)
	})

	t.Run("simple", func(t *testing.T) {
		c := NewCache(5)

		wasInCache := c.Set("aaa", 100)
		require.False(t, wasInCache)

		wasInCache = c.Set("bbb", 200)
		require.False(t, wasInCache)

		val, ok := c.Get("aaa")
		require.True(t, ok)
		require.Equal(t, 100, val)

		val, ok = c.Get("bbb")
		require.True(t, ok)
		require.Equal(t, 200, val)

		wasInCache = c.Set("aaa", 300)
		require.True(t, wasInCache)

		val, ok = c.Get("aaa")
		require.True(t, ok)
		require.Equal(t, 300, val)

		val, ok = c.Get("ccc")
		require.False(t, ok)
		require.Nil(t, val)
	})

	t.Run("cache logic", func(t *testing.T) {
		c := NewCache(4)
		c.Set("Russia", 1) // [1]
		c.Set("Canada", 2) // [2,1]
		c.Set("China", 3)  // [3,2,1]
		c.Set("USA", 4)    // [4,3,2,1]

		_, ok := c.Get("Canada")
		require.True(t, ok)
		c.Set("Brazil", 5) // [5,2,4,3]
		_, ok = c.Get("Russia")
		require.False(t, ok)
	})

	t.Run("purge logic", func(t *testing.T) {
		c := NewCache(4)
		c.Set("key1", 100)
		c.Set("key2", 200)
		c.Set("key3", 300)
		c.Clear()
		_, ok := c.Get("key1")
		require.False(t, ok)
	})
}
