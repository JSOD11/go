package bitvector

import "fmt"

type Bitvector struct {
	vector []int64
	debug  bool
}

func (b *Bitvector) debugLog(out string, args ...interface{}) {
	if b.debug {
		fmt.Printf(out, args...)
	}
}

func ConstructBitvector(debug bool) *Bitvector {
	vector := []int64{0}
	if debug {
		fmt.Printf("\n\n———————— Constructing bit vector ————————\n\n\n\n")
	}
	return &Bitvector{vector, debug}
}

func (b *Bitvector) Add(n int64) {
	multiple := n / 64
	remainder := n % 64
	b.debugLog("\nMultiple, remainder, vector: %v, %v, %v\n", multiple, remainder, b.vector)
	for i := int64(len(b.vector)) + 1; i < multiple+2; i++ {
		b.vector = append(b.vector, 0)
	}
	b.vector[multiple] += 1 << remainder
	b.debugLog("Vector: %v\n\n\n\n", b.vector)
}

func (b *Bitvector) AddMany(values ...int64) {
	for _, n := range values {
		b.Add(n)
	}
}

func (b *Bitvector) Contains(n int64) bool {
	b.debugLog("\nSearching for: %v in vector %v\n", n, b.vector)
	multiple := n / 64
	remainder := n % 64
	b.debugLog("Multiple, remainder, vector: %v, %v, %v\n", multiple, remainder, b.vector)
	if multiple >= int64(len(b.vector)) {
		return false
	}
	value := b.vector[multiple]
	b.debugLog("Value: %v, bool: %v\n", value, value&(1<<remainder) != 0)
	return value&(1<<remainder) != 0
}

func (b *Bitvector) ContainsMany(values ...int64) bool {
	for _, n := range values {
		if !b.Contains(n) {
			return false
		}
	}
	return true
}

func (b *Bitvector) PrintVector() {
	fmt.Printf("\nPrintVector: %v", b.vector)
}
