// Package bitvector implements a bitvector type built on top of 64-bit integers.
// Integer k exists in the bitvector if bit (k % 64) is set in words[k / 64].
package bitvector

import (
	"fmt"
	"go-book/debug"
	"strconv"
)

// Bitvector
// A container which supports adding int64 values and testing for their presence.
type Bitvector struct {
	words []uint64
	*debug.Debugger
}

// ConstructBitvector
// Bitvector constructor. Optional setting for debug logs.
func ConstructBitvector(debugActive bool) *Bitvector {
	words := []uint64{0}
	debugger := &debug.Debugger{Active: debugActive}
	debugger.DebugLog("\n\n———————— Constructing bitvector ————————\n")
	return &Bitvector{words: words, Debugger: debugger}
}

// Add
// Adds n to the bitvector.
func (b *Bitvector) Add(n uint64) {
	word, bit := n/64, n%64
	b.DebugLog("\nADD: %v\nWord, bit, bitvector: %v, %v, %v\n", n, word, bit, b)
	for uint64(len(b.words)) <= word {
		b.words = append(b.words, 0)
	}
	b.words[word] |= 1 << bit
	b.DebugLog("bitvector after add: %v\n", b)
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
	word, bit := n/64, n%64
	b.DebugLog("\nCONTAINS: %v\nWord, bit, bitvector: %v, %v, %v\n", n, word, bit, b)
	if word >= uint64(len(b.words)) {
		return false
	}
	value := b.words[word]
	b.DebugLog("Result: %v\n", value&(1<<bit) != 0)
	return value&(1<<bit) != 0
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
	word, bit := n/64, n%64
	b.DebugLog("\nREMOVE: %v\nWord, bit, bitvector: %v, %v, %v\n", n, word, bit, b)
	if word < uint64(len(b.words)) && (b.words[word]&(1<<bit) != 0) {
		b.words[word] ^= 1 << bit
	}
	b.DebugLog("bitvector after remove: %v\n", b)
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
	for word := range b.words {
		for i := range 64 {
			if b.words[word]&(1<<i) != 0 {
				buf = append(buf, strconv.Itoa(64*word+i))
			}
		}
	}
	return fmt.Sprintf("%v", buf)
}
