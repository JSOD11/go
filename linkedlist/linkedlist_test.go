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

	ll.Append(3)
	ll.Append(8)
	ll.Append(21)
	ll.Append(34)

	ll.Insert(5, 1)
	ll.Insert(13, 3)

	assert.True(t, ll.Contains(21))
	assert.True(t, ll.Contains(13))
	assert.False(t, ll.Contains(55))

	ll.Remove(8)
	ll.Remove(5)
	ll.Remove(34)

	assert.False(t, ll.Contains(8))
	assert.False(t, ll.Contains(34))
}

func TestLinkedListMisc(t *testing.T) {
	ll := ConstructLinkedList(debugLogs)

	assert.True(t, ll.IsEmpty())
	assert.Equal(t, 0, ll.Length())
	ll.AppendMany(0, 1, 1, 2, 3, 5, 8, 13, 21, 34, 55, 89, 144)
	assert.False(t, ll.IsEmpty())
	assert.Equal(t, 13, ll.Length())

	ll.Clear()
	assert.True(t, ll.IsEmpty())
	assert.Equal(t, 0, ll.Length())
}
