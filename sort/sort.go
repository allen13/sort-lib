package sort

// SortFunc represents a sorting function that can sort a slice of any ordered type
type SortFunc[T any] func([]T, func(T, T) bool)

// Less is a comparison function that returns true if a should come before b
type Less[T any] func(a, T, b T) bool

// QuickSort implements the quicksort algorithm
func QuickSort[T any](arr []T, less func(T, T) bool) {
	if len(arr) <= 1 {
		return
	}

	pivot := arr[len(arr)-1]
	i := 0

	for j := 0; j < len(arr)-1; j++ {
		if less(arr[j], pivot) {
			arr[i], arr[j] = arr[j], arr[i]
			i++
		}
	}

	arr[i], arr[len(arr)-1] = arr[len(arr)-1], arr[i]

	QuickSort(arr[:i], less)
	QuickSort(arr[i+1:], less)
}

// MergeSort implements the merge sort algorithm
func MergeSort[T any](arr []T, less func(T, T) bool) {
	if len(arr) <= 1 {
		return
	}

	mid := len(arr) / 2
	left := make([]T, mid)
	right := make([]T, len(arr)-mid)

	copy(left, arr[:mid])
	copy(right, arr[mid:])

	MergeSort(left, less)
	MergeSort(right, less)

	merge(arr, left, right, less)
}

// merge merges two sorted slices into one
func merge[T any](arr, left, right []T, less func(T, T) bool) {
	i, j, k := 0, 0, 0

	for i < len(left) && j < len(right) {
		if less(left[i], right[j]) {
			arr[k] = left[i]
			i++
		} else {
			arr[k] = right[j]
			j++
		}
		k++
	}

	for i < len(left) {
		arr[k] = left[i]
		i++
		k++
	}

	for j < len(right) {
		arr[k] = right[j]
		j++
		k++
	}
}

// InsertionSort implements the insertion sort algorithm
func InsertionSort[T any](arr []T, less func(T, T) bool) {
	for i := 1; i < len(arr); i++ {
		key := arr[i]
		j := i - 1

		for j >= 0 && less(key, arr[j]) {
			arr[j+1] = arr[j]
			j--
		}

		arr[j+1] = key
	}
}

// BubbleSort implements the bubble sort algorithm
func BubbleSort[T any](arr []T, less func(T, T) bool) {
	n := len(arr)
	for i := 0; i < n-1; i++ {
		swapped := false
		for j := 0; j < n-i-1; j++ {
			if less(arr[j+1], arr[j]) {
				arr[j], arr[j+1] = arr[j+1], arr[j]
				swapped = true
			}
		}
		if !swapped {
			break
		}
	}
}

// HeapSort implements the heap sort algorithm
func HeapSort[T any](arr []T, less func(T, T) bool) {
	n := len(arr)

	// Build max heap
	for i := n/2 - 1; i >= 0; i-- {
		heapify(arr, n, i, less)
	}

	// Extract elements from heap one by one
	for i := n - 1; i > 0; i-- {
		arr[0], arr[i] = arr[i], arr[0]
		heapify(arr, i, 0, less)
	}
}

// heapify maintains the heap property
func heapify[T any](arr []T, n, i int, less func(T, T) bool) {
	largest := i
	left := 2*i + 1
	right := 2*i + 2

	if left < n && less(arr[largest], arr[left]) {
		largest = left
	}

	if right < n && less(arr[largest], arr[right]) {
		largest = right
	}

	if largest != i {
		arr[i], arr[largest] = arr[largest], arr[i]
		heapify(arr, n, largest, less)
	}
}

// IsSorted checks if a slice is sorted according to the less function
func IsSorted[T any](arr []T, less func(T, T) bool) bool {
	for i := 1; i < len(arr); i++ {
		if less(arr[i], arr[i-1]) {
			return false
		}
	}
	return true
}