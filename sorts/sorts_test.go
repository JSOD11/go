package sorts

import (
	"math/rand"
	"slices"
	"sort"
	"testing"
)

var sortFuncs = []sortFunction{
	newBuiltInSort(),
	newMergeSort(),
	newInsertionSort(),
	newBubbleSort(),
	newQuickSort(),
}

var (
	testSlices  [][]int
	numSlices   = 10
	sliceLength = 100000
	maxValue    = 100
)

func randIntSlice(sliceLength int) []int {
	slice := make([]int, sliceLength)
	for i := range sliceLength {
		slice[i] = rand.Intn(maxValue)
	}
	return slice
}

func init() {
	for range numSlices {
		testSlices = append(testSlices, randIntSlice(sliceLength))
	}
}

func TestSorts(t *testing.T) {

	for _, slice := range testSlices {
		//fmt.Printf("\nSlice: %v\n", slice)
		var sortedSlice = make([]int, len(slice))
		copy(sortedSlice, slice)
		sort.Ints(sortedSlice)
		for _, sortFunc := range sortFuncs {
			var sliceCopy = make([]int, len(slice))
			copy(sliceCopy, slice)

			s := sortFunc.Sort(sliceCopy)

			ok := slices.Equal(sortedSlice, s)
			if !ok {
				t.Errorf("%v failed: expected: %v, got: %v", sortFunc, sortedSlice, s)
			}
			//fmt.Printf("%v: %v\n", sortFunc, s)
		}
	}
}

func BenchmarkBuiltInSort(b *testing.B) {
	builtIn := newBuiltInSort()
	slice := randIntSlice(sliceLength)
	for b.Loop() {
		builtIn.Sort(slice)
	}
}

func BenchmarkMergeSort(b *testing.B) {
	merge := newMergeSort()
	slice := randIntSlice(sliceLength)
	for b.Loop() {
		merge.Sort(slice)
	}
}

func BenchmarkInsertionSort(b *testing.B) {
	insertion := newInsertionSort()
	slice := randIntSlice(sliceLength)
	for b.Loop() {
		insertion.Sort(slice)
	}
}

func BenchmarkBubbleSort(b *testing.B) {
	bubble := newBubbleSort()
	slice := randIntSlice(sliceLength)
	for b.Loop() {
		bubble.Sort(slice)
	}
}

func BenchmarkQuickSort(b *testing.B) {
	quick := newQuickSort()
	slice := randIntSlice(sliceLength)
	for b.Loop() {
		quick.Sort(slice)
	}
}
