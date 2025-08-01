package bitvector

import (
	"fmt"
	"go-book/debug"
	"strconv"
)

// Bitvector
// A container which supports adding int64 values and testing for their presence.
type Bitvector struct {
	vector []uint64
	*debug.Debugger
}

// ConstructBitvector
// Bitvector constructor. Optional setting for debug logs.
func ConstructBitvector(debugActive bool) *Bitvector {
	vector := []uint64{0}
	debugger := &debug.Debugger{Active: debugActive}
	debugger.DebugLog("\n\n———————— Constructing bit vector ————————\n")
	return &Bitvector{vector: vector, Debugger: debugger}
}

// Add
// Adds n to the bitvector.
func (b *Bitvector) Add(n uint64) {
	multiple, remainder := n/64, n%64
	b.DebugLog("\nADD: %v\nMultiple, remainder, vector: %v, %v, %v\n", n, multiple, remainder, b)
	for uint64(len(b.vector)) <= multiple {
		b.vector = append(b.vector, 0)
	}
	b.vector[multiple] |= 1 << remainder
	b.DebugLog("vector after add: %v\n", b)
}

// AddMany
// Variadic function that will add multiple values to the bitvector.
func (b *Bitvector) AddMany(values ...uint64) {
	for _, n := range values {
		b.Add(n)
	}
}

// Contains
// Returns a boolean signifying whether n is currently present in the bitvector.
func (b *Bitvector) Contains(n uint64) bool {
	multiple, remainder := n/64, n%64
	b.DebugLog("\nCONTAINS: %v\nMultiple, remainder, vector: %v, %v, %v\n", n, multiple, remainder, b)
	if multiple >= uint64(len(b.vector)) {
		return false
	}
	value := b.vector[multiple]
	b.DebugLog("Result: %v\n", value&(1<<remainder) != 0)
	return value&(1<<remainder) != 0
}

// ContainsMany
// Variadic function that returns true if all values are present in the bitvector.
func (b *Bitvector) ContainsMany(values ...uint64) bool {
	for _, n := range values {
		if !b.Contains(n) {
			return false
		}
	}
	return true
}

// Remove
// Sets the presence of n in the bitvector to false.
func (b *Bitvector) Remove(n uint64) {
	multiple, remainder := n/64, n%64
	b.DebugLog("\nREMOVE: %v\nMultiple, remainder, vector: %v, %v, %v\n", n, multiple, remainder, b)
	if multiple < uint64(len(b.vector)) && (b.vector[multiple]&(1<<remainder) != 0) {
		b.vector[multiple] ^= 1 << remainder
	}
	b.DebugLog("vector after remove: %v\n", b)
}

// RemoveMany
// Variadic function that removes multiple values from the bitvector.
func (b *Bitvector) RemoveMany(values ...uint64) {
	for _, n := range values {
		b.Remove(n)
	}
}

// String
// Print out the bitvector in human-readable string format.
func (b *Bitvector) String() string {
	var buf []string
	for word := range b.vector {
		for i := range 64 {
			if b.vector[word]&(1<<i) != 0 {
				buf = append(buf, strconv.Itoa(64*word+i))
			}
		}
	}
	return fmt.Sprintf("%v", buf)
}
