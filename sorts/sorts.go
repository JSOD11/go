package sorts

import (
	"sort"
)

type sortFunction interface {
	String() string
	Sort([]int) []int
}

type quickSort struct {
	name string
}

func newQuickSort() *quickSort {
	return &quickSort{name: "quickSort"}
}

func (s *quickSort) String() string {
	return s.name
}

func (s *quickSort) Sort(a []int) []int {
	quickSortWithBounds(a, 0, len(a)-1)
	return a
}

func quickSortWithBounds(a []int, l, r int) {
	if l < r {
		pivotIndex := partition(a, l, r)
		quickSortWithBounds(a, l, pivotIndex-1)
		quickSortWithBounds(a, pivotIndex+1, r)
	}
}

func partition(a []int, l, r int) int {
	pivot := a[r]
	i := l - 1

	for j := l; j < r; j++ {
		if a[j] < pivot {
			i++
			a[i], a[j] = a[j], a[i]
		}
	}

	a[i+1], a[r] = a[r], a[i+1]
	return i + 1
}

type bubbleSort struct {
	name string
}

func newBubbleSort() *bubbleSort {
	return &bubbleSort{name: "bubbleSort"}
}

func (s *bubbleSort) String() string {
	return s.name
}

func (s *bubbleSort) Sort(a []int) []int {
	N := len(a)
	swapped := true
	for swapped {
		swapped = false
		for i := 0; i < N-1; i++ {
			if a[i] > a[i+1] {
				a[i], a[i+1] = a[i+1], a[i]
				swapped = true
			}
		}
	}
	return a
}

type insertionSort struct {
	name string
}

func newInsertionSort() *insertionSort {
	return &insertionSort{name: "insertionSort"}
}

func (s *insertionSort) String() string {
	return s.name
}

func (s *insertionSort) Sort(a []int) []int {
	N := len(a)
	for i := 0; i < N; i++ {
		minSeen, minIndex := a[i], i
		for j := i; j < N; j++ {
			if a[j] < minSeen {
				minSeen = a[j]
				minIndex = j
			}
		}
		a[i], a[minIndex] = a[minIndex], a[i]
	}
	return a
}

type mergeSort struct {
	name string
}

func newMergeSort() *mergeSort {
	return &mergeSort{name: "mergeSort"}
}

func (s *mergeSort) String() string {
	return s.name
}

func (s *mergeSort) Sort(a []int) []int {
	l, r := 0, len(a)
	if l == r-1 {
		return a
	}
	m := (l + r) / 2
	return s.merge(s.Sort(a[:m]), s.Sort(a[m:]))
}

func (s *mergeSort) merge(a, b []int) []int {
	i, j := 0, 0
	var result []int
	for i < len(a) && j < len(b) {
		if a[i] <= b[j] {
			result = append(result, a[i])
			i++
		} else {
			result = append(result, b[j])
			j++
		}
	}
	if i < len(a) {
		result = append(result, a[i:]...)
	} else {
		result = append(result, b[j:]...)
	}
	return result
}

type builtInSort struct {
	name string
}

func newBuiltInSort() *builtInSort {
	return &builtInSort{name: "builtInSort"}
}

func (s *builtInSort) String() string {
	return s.name
}

func (s *builtInSort) Sort(a []int) []int {
	sort.Ints(a)
	return a
}
