package bitvector

import (
	"go-book/assert"
	"testing"
)

const (
	debugLogs = false
)

// Test Add and Contains bitvector behavior for numbers less than 64.
func TestBitvectorBasic(t *testing.T) {
	bitvector := ConstructBitvector(debugLogs)

	assert.False(t, bitvector.Contains(0))

	bitvector.Add(0)
	bitvector.Add(1)
	bitvector.Add(2)
	bitvector.Add(48)
	bitvector.Add(21)

	assert.True(t, bitvector.Contains(0))
	assert.True(t, bitvector.Contains(1))
	assert.True(t, bitvector.Contains(2))
	assert.True(t, bitvector.Contains(21))
	assert.True(t, bitvector.Contains(48))
	assert.False(t, bitvector.Contains(5))
	assert.False(t, bitvector.Contains(25))
	assert.False(t, bitvector.Contains(63))
}

// Test Add and Contains for numbers greater than 64.
func TestBitvectorGreaterChunks(t *testing.T) {
	bitvector := ConstructBitvector(debugLogs)

	bitvector.Add(64)
	assert.True(t, bitvector.Contains(64))
	assert.False(t, bitvector.Contains(6767))
	bitvector.Add(6767)
	assert.True(t, bitvector.Contains(6767))
}

// Test AddMany and ContainsMany.
func TestBitvectorAddMany(t *testing.T) {
	bitvector := ConstructBitvector(debugLogs)

	bitvector.AddMany(1, 25, 34, 513, 0, 27, 95, 212, 1025, 891)
	assert.True(t, bitvector.ContainsMany(25, 513, 0, 212, 891))
	assert.False(t, bitvector.ContainsMany(25, 2, 513))
}

// Test interleaving AddMany, ContainsMany, Remove, RemoveMany.
func TestBitvectorRemove(t *testing.T) {
	bitvector := ConstructBitvector(debugLogs)

	bitvector.AddMany(1, 25, 34, 513, 0, 27, 95, 212, 1025, 891)
	assert.True(t, bitvector.ContainsMany(25, 34, 212))

	bitvector.Remove(34)
	assert.False(t, bitvector.Contains(34))

	bitvector.AddMany(65, 14, 901, 542, 318, 96)

	assert.True(t, bitvector.ContainsMany(1, 25, 513, 0, 212, 14, 96))

	bitvector.RemoveMany(1, 25, 513, 0, 212, 14, 96, 100, 101, 102)
	assert.False(t, bitvector.ContainsMany(102, 1, 25, 212, 1025))
}
