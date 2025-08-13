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
}

func TestSorts(t *testing.T) {

	var (
		testSlices  [][]int
		numSlices   = 10
		sliceLength = 20
		maxValue    = 100
	)

	for range numSlices {
		var slice []int
		for range sliceLength {
			slice = append(slice, rand.Intn(maxValue))
		}
		testSlices = append(testSlices, slice)
	}

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

var s = []int{23, 14, 5, 16, 82, 90, 95, 21, 45, 75, 28, 85, 91}

func BenchmarkBuiltInSort(b *testing.B) {
	for i := 0; i < b.N; i++ {
		sortFuncs[0].Sort(s)
	}
}

func BenchmarkMergeSort(b *testing.B) {
	for i := 0; i < b.N; i++ {
		sortFuncs[1].Sort(s)
	}
}

func BenchmarkInsertionSort(b *testing.B) {
	for i := 0; i < b.N; i++ {
		sortFuncs[2].Sort(s)
	}
}
