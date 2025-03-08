package sort

import (
	"math/rand"
	"testing"
	"time"
)

func TestSortingAlgorithms(t *testing.T) {
	tests := []struct {
		name     string
		sortFunc SortFunc[int]
	}{
		{"QuickSort", QuickSort[int]},
		{"MergeSort", MergeSort[int]},
		{"InsertionSort", InsertionSort[int]},
		{"BubbleSort", BubbleSort[int]},
		{"HeapSort", HeapSort[int]},
	}

	testCases := []struct {
		name string
		arr  []int
		want []int
	}{
		{
			name: "Empty array",
			arr:  []int{},
			want: []int{},
		},
		{
			name: "Single element",
			arr:  []int{1},
			want: []int{1},
		},
		{
			name: "Already sorted",
			arr:  []int{1, 2, 3, 4, 5},
			want: []int{1, 2, 3, 4, 5},
		},
		{
			name: "Reverse sorted",
			arr:  []int{5, 4, 3, 2, 1},
			want: []int{1, 2, 3, 4, 5},
		},
		{
			name: "Random order",
			arr:  []int{3, 1, 4, 1, 5, 9, 2, 6, 5, 3, 5},
			want: []int{1, 1, 2, 3, 3, 4, 5, 5, 5, 6, 9},
		},
		{
			name: "Duplicate elements",
			arr:  []int{3, 3, 3, 1, 1, 2, 2},
			want: []int{1, 1, 2, 2, 3, 3, 3},
		},
	}

	less := func(a, b int) bool {
		return a < b
	}

	for _, algo := range tests {
		t.Run(algo.name, func(t *testing.T) {
			for _, tc := range testCases {
				t.Run(tc.name, func(t *testing.T) {
					// Create a copy of the input array
					arr := make([]int, len(tc.arr))
					copy(arr, tc.arr)

					// Sort the array
					algo.sortFunc(arr, less)

					// Check if the array is sorted correctly
					if !IsSorted(arr, less) {
						t.Errorf("Array is not sorted: %v", arr)
					}

					// Check if the sorted array matches the expected result
					if len(arr) != len(tc.want) {
						t.Errorf("got len %d, want len %d", len(arr), len(tc.want))
					}

					for i := range arr {
						if arr[i] != tc.want[i] {
							t.Errorf("at index %d: got %d, want %d", i, arr[i], tc.want[i])
						}
					}
				})
			}
		})
	}
}

func BenchmarkSortingAlgorithms(b *testing.B) {
	sizes := []int{10, 100, 1000, 10000}
	algorithms := []struct {
		name     string
		sortFunc SortFunc[int]
	}{
		{"QuickSort", QuickSort[int]},
		{"MergeSort", MergeSort[int]},
		{"InsertionSort", InsertionSort[int]},
		{"BubbleSort", BubbleSort[int]},
		{"HeapSort", HeapSort[int]},
	}

	less := func(a, b int) bool {
		return a < b
	}

	rand.Seed(time.Now().UnixNano())

	for _, size := range sizes {
		// Generate a random array of the current size
		arr := make([]int, size)
		for i := range arr {
			arr[i] = rand.Intn(1000000)
		}

		for _, algo := range algorithms {
			b.Run(algo.name+"-"+string(rune(size)), func(b *testing.B) {
				for i := 0; i < b.N; i++ {
					// Create a copy of the array for each iteration
					testArr := make([]int, len(arr))
					copy(testArr, arr)
					algo.sortFunc(testArr, less)
				}
			})
		}
	}
}