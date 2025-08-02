package linkedlist

import (
	"go-book/assert"
	"testing"
)

const (
	debugLogs = false
)

func TestLinkedListBasic(t *testing.T) {
	ll := ConstructLinkedList(debugLogs)

	ll.Append(3, 3)
	ll.Append(8, 8)
	ll.Append(21, 21)
	ll.Append(34, 34)

	ll.Insert(5, 5, 1)
	ll.Insert(13, 13, 3)

	val, ok := ll.Get(21)
	assert.True(t, ok)
	assert.Equal(t, int64(21), val)

	val, ok = ll.Get(13)
	assert.True(t, ok)
	assert.Equal(t, int64(13), val)

	val, ok = ll.Get(55)
	assert.False(t, ok)
	assert.Equal(t, int64(0), val)

	ll.Remove(8)
	ll.Remove(5)
	ll.Remove(34)

	val, ok = ll.Get(8)
	assert.False(t, ok)
	assert.Equal(t, int64(0), val)

}

func TestLinkedListMisc(t *testing.T) {
	ll := ConstructLinkedList(debugLogs)

	assert.True(t, ll.IsEmpty())
	assert.Equal(t, 0, ll.Length())

	keys := []int64{0, 1, 1, 2, 3, 5, 8, 13, 21, 34, 55, 89, 144}
	values := []int64{0, 1, 1, 2, 3, 5, 8, 13, 21, 34, 55, 89, 144}
	err := ll.AppendMany(keys, values)

	assert.Nil(t, err)
	assert.False(t, ll.IsEmpty())
	assert.Equal(t, 13, ll.Length())

	ll.Clear()
	assert.True(t, ll.IsEmpty())
	assert.Equal(t, 0, ll.Length())
}
