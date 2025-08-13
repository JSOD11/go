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
