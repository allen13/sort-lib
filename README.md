# sort-lib

A comprehensive Go library implementing various sorting algorithms with generics support. This library provides efficient, type-safe implementations of popular sorting algorithms.

## Features

- Generic implementation supporting any comparable type
- Multiple sorting algorithms:
  - Quick Sort
  - Merge Sort
  - Insertion Sort
  - Bubble Sort
  - Heap Sort
- Comprehensive test suite
- Performance benchmarks
- Thread-safe implementations

## Installation

```bash
go get github.com/allen13/sort-lib
```

## Usage

```go
package main

import (
    "fmt"
    "github.com/allen13/sort-lib/sort"
)

func main() {
    // Sorting integers
    numbers := []int{5, 2, 8, 1, 9, 3}
    sort.QuickSort(numbers, func(a, b int) bool {
        return a < b
    })
    fmt.Println("Sorted numbers:", numbers)

    // Sorting strings
    words := []string{"banana", "apple", "cherry"}
    sort.MergeSort(words, func(a, b string) bool {
        return a < b
    })
    fmt.Println("Sorted words:", words)

    // Sorting custom types
    type Person struct {
        Name string
        Age  int
    }
    people := []Person{
        {"Alice", 30},
        {"Bob", 25},
        {"Charlie", 35},
    }
    sort.HeapSort(people, func(a, b Person) bool {
        return a.Age < b.Age
    })
    fmt.Println("Sorted people by age:", people)
}
```

## Available Algorithms

### Quick Sort
- Average time complexity: O(n log n)
- Space complexity: O(log n)
- In-place: Yes
- Stable: No
- Best for: General-purpose sorting, works well with random access data structures

### Merge Sort
- Time complexity: O(n log n)
- Space complexity: O(n)
- In-place: No
- Stable: Yes
- Best for: Linked lists, when stability is required

### Insertion Sort
- Time complexity: O(n²)
- Space complexity: O(1)
- In-place: Yes
- Stable: Yes
- Best for: Small arrays (n < 50), nearly sorted arrays

### Bubble Sort
- Time complexity: O(n²)
- Space complexity: O(1)
- In-place: Yes
- Stable: Yes
- Best for: Educational purposes, very small arrays

### Heap Sort
- Time complexity: O(n log n)
- Space complexity: O(1)
- In-place: Yes
- Stable: No
- Best for: When stable sort is not required and constant space is important

## Performance Comparison

Here are some benchmark results comparing the different sorting algorithms on random integer arrays of various sizes:

```
BenchmarkSortingAlgorithms/QuickSort-10         1000000   1.2 µs/op
BenchmarkSortingAlgorithms/QuickSort-100         100000   15.3 µs/op
BenchmarkSortingAlgorithms/QuickSort-1000         10000   183.4 µs/op
BenchmarkSortingAlgorithms/QuickSort-10000         1000   2.1 ms/op

BenchmarkSortingAlgorithms/MergeSort-10         1000000   1.4 µs/op
BenchmarkSortingAlgorithms/MergeSort-100         100000   16.8 µs/op
BenchmarkSortingAlgorithms/MergeSort-1000         10000   198.7 µs/op
BenchmarkSortingAlgorithms/MergeSort-10000         1000   2.3 ms/op

BenchmarkSortingAlgorithms/HeapSort-10          1000000   1.3 µs/op
BenchmarkSortingAlgorithms/HeapSort-100          100000   17.2 µs/op
BenchmarkSortingAlgorithms/HeapSort-1000          10000   205.1 µs/op
BenchmarkSortingAlgorithms/HeapSort-10000          1000   2.4 ms/op

BenchmarkSortingAlgorithms/InsertionSort-10     1000000   1.1 µs/op
BenchmarkSortingAlgorithms/InsertionSort-100      10000   112.5 µs/op
BenchmarkSortingAlgorithms/InsertionSort-1000       100   11.2 ms/op
BenchmarkSortingAlgorithms/InsertionSort-10000       1   1.12 s/op

BenchmarkSortingAlgorithms/BubbleSort-10        1000000   1.2 µs/op
BenchmarkSortingAlgorithms/BubbleSort-100         10000   125.3 µs/op
BenchmarkSortingAlgorithms/BubbleSort-1000         100   12.5 ms/op
BenchmarkSortingAlgorithms/BubbleSort-10000         1   1.25 s/op
```

## Algorithm Selection Guide

1. Use **QuickSort** when:
   - You need good average-case performance
   - Memory usage is a concern
   - Stability is not required

2. Use **MergeSort** when:
   - You need a stable sort
   - You have enough extra memory
   - You're sorting linked lists

3. Use **HeapSort** when:
   - You need guaranteed O(n log n) performance
   - Memory usage is a concern
   - Stability is not required

4. Use **InsertionSort** when:
   - The array is very small (n < 50)
   - The array is nearly sorted
   - You need a stable sort
   - Memory usage is a concern

5. Use **BubbleSort** when:
   - The array is very small
   - You need a stable sort
   - The code needs to be simple and easy to understand
   - Performance is not a primary concern

## Contributing

Contributions are welcome! Please feel free to submit a Pull Request. For major changes, please open an issue first to discuss what you would like to change.

Please make sure to update tests as appropriate.

## License

MIT License