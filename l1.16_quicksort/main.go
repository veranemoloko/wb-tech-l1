package main

import (
	"cmp"
	"fmt"
)

// swap exchanges elements at positions i and j
func swap[T any](arr []T, i, j int) {
	arr[i], arr[j] = arr[j], arr[i]
}

// quickSort sorts a slice in-place using QuickSort algorithm
func quickSort[T cmp.Ordered](arr []T) {
	if len(arr) <= 1 {
		return // already sorted
	}

	head, tail := 0, len(arr)-1

	// choose the first element as pivot
	pivot := arr[head]

	// partition: maxBeforePivot points to last element < pivot
	maxBeforePivot := head

	for j := head + 1; j <= tail; j++ {
		if arr[j] < pivot {
			maxBeforePivot++
			swap(arr, maxBeforePivot, j)
		}
	}

	// put pivot in correct position
	swap(arr, head, maxBeforePivot)
	pivotIndex := maxBeforePivot

	// recursively sort left and right parts
	quickSort(arr[:pivotIndex])
	quickSort(arr[pivotIndex+1:])
}

func main() {
	dataInt := []int{10, 7, 8, 9, 1, 5, 3, -2, 11, 4, -2, 0, 0, 0, 0, 0, 1, 1, 1, 0, 10}
	quickSort(dataInt)
	fmt.Println("Sorted ints:", dataInt)

	dataFloats := []float64{3.14, 2.71, 1.41, -0.5, 10.0}
	quickSort(dataFloats)
	fmt.Println("Sorted floats:", dataFloats)

	dataStrs := []string{"bear", "rabbit", "wombat", "i"}
	quickSort(dataStrs)
	fmt.Println("Sorted strings:", dataStrs)
}
