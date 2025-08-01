package bitvector

import (
	"testing"
)

const (
	debugLogs = false
)

func AssertEqual(t *testing.T, a, b interface{}) {
	if a != b {
		t.Errorf("Expected %v, got %v", a, b)
	}
}

func TestBitvectorBasic(t *testing.T) {
	bitvector := ConstructBitvector(debugLogs)
	bitvector.Add(0)
	bitvector.Add(1)
	bitvector.Add(2)
	bitvector.Add(48)
	bitvector.Add(21)

	AssertEqual(t, true, bitvector.Contains(0))
	AssertEqual(t, true, bitvector.Contains(1))
	AssertEqual(t, true, bitvector.Contains(2))
	AssertEqual(t, true, bitvector.Contains(21))
	AssertEqual(t, true, bitvector.Contains(48))
	AssertEqual(t, false, bitvector.Contains(5))
	AssertEqual(t, false, bitvector.Contains(25))
	AssertEqual(t, false, bitvector.Contains(63))
}

func TestBitvectorGreaterChunks(t *testing.T) {
	bitvector := ConstructBitvector(debugLogs)

	bitvector.Add(64)
	AssertEqual(t, true, bitvector.Contains(64))
	AssertEqual(t, false, bitvector.Contains(6767))
	bitvector.Add(6767)
	AssertEqual(t, true, bitvector.Contains(6767))
}

func TestBitvectorAddMany(t *testing.T) {
	bitvector := ConstructBitvector(debugLogs)

	bitvector.AddMany(1, 25, 34, 513, 0, 27, 95, 212, 1025, 891)
	AssertEqual(t, true, bitvector.ContainsMany(25, 513, 0, 212, 891))
	AssertEqual(t, false, bitvector.ContainsMany(25, 2, 513))
}
