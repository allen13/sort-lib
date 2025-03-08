package main

import (
	"fmt"

	"github.com/allen13/sort-lib/sort"
)

// Person represents a person with a name and age
type Person struct {
	Name string
	Age  int
}

func main() {
	// Example 1: Sorting integers
	numbers := []int{5, 2, 8, 1, 9, 3}
	fmt.Println("Original numbers:", numbers)
	
	sort.QuickSort(numbers, func(a, b int) bool {
		return a < b
	})
	fmt.Println("Sorted numbers:", numbers)

	// Example 2: Sorting strings
	words := []string{"banana", "apple", "cherry", "date"}
	fmt.Println("\nOriginal words:", words)
	
	sort.MergeSort(words, func(a, b string) bool {
		return a < b
	})
	fmt.Println("Sorted words:", words)

	// Example 3: Sorting custom types (Person) by age
	people := []Person{
		{"Alice", 30},
		{"Bob", 25},
		{"Charlie", 35},
		{"David", 28},
	}
	fmt.Println("\nOriginal people:", people)
	
	sort.HeapSort(people, func(a, b Person) bool {
		return a.Age < b.Age
	})
	fmt.Println("Sorted people by age:", people)

	// Example 4: Sorting custom types (Person) by name
	sort.InsertionSort(people, func(a, b Person) bool {
		return a.Name < b.Name
	})
	fmt.Println("Sorted people by name:", people)

	// Example 5: Sorting float64 numbers
	decimals := []float64{3.14, 1.41, 2.71, 1.73}
	fmt.Println("\nOriginal decimals:", decimals)
	
	sort.BubbleSort(decimals, func(a, b float64) bool {
		return a < b
	})
	fmt.Println("Sorted decimals:", decimals)
}