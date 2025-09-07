package main

import (
	"cmp"
	"fmt"
)

func swap[T any](arr []T, i, j int) {
	arr[i], arr[j] = arr[j], arr[i]
}

func quickSort[T cmp.Ordered](arr []T) {
	head, tail := 0, len(arr)-1

	if len(arr) <= 1 {
		return
	}

	// This avoids creating deeper recursive calls unnecessarily
	// and helps reduce recursion depth overall.
	if len(arr) == 2 {
		if arr[head] > arr[tail] {
			swap(arr, head, tail)
		}
		return
	}

	// choose the pivot as the rightmost element
	pivot := arr[tail]

	// maxBeforePivot points to the last element smaller than the pivot.
	// It starts at (head-1), meaning "no elements are smaller yet".
	maxBeforePivot := head - 1

	// move all elements smaller than pivot to the left side
	for j := head; j < tail; j++ {
		if arr[j] < pivot {
			maxBeforePivot++
			swap(arr, maxBeforePivot, j)
		}
	}

	// place pivot into its correct sorted position
	swap(arr, maxBeforePivot+1, tail)
	pivotIndex := maxBeforePivot + 1

	// recursively sort (all < pivot)
	quickSort(arr[head:pivotIndex])

	// recursively sort (all > pivot)
	quickSort(arr[pivotIndex+1 : tail+1])
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
