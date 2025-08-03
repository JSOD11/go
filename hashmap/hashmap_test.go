package hashmap

import (
	"go-book/assert"
	"testing"
)

const (
	debugLogs = true
)

func TestHashMapBasic(t *testing.T) {
	h := ConstructHashMap(debugLogs)
	h.Put(3, 3)
	h.Put(5, 5)
	h.Put(8, 8)
	h.Put(25, 1)
	h.Put(150, 3)

	val, found := h.Get(3)
	assert.True(t, found)
	assert.Equal(t, int64(3), val)

	val, found = h.Get(8)
	assert.True(t, found)
	assert.Equal(t, int64(8), val)

	val, found = h.Get(10)
	assert.False(t, found)
	assert.Equal(t, int64(0), val)

	val, found = h.Get(100)
	assert.False(t, found)
	assert.Equal(t, int64(0), val)

	val, found = h.Get(150)
	assert.True(t, found)
	assert.Equal(t, int64(3), val)

	h.Put(235, 121)

	val, found = h.Get(235)
	assert.True(t, found)
	assert.Equal(t, int64(121), val)

	h.Remove(8)

	val, found = h.Get(8)
	assert.False(t, found)
	assert.Equal(t, int64(0), val)
}
