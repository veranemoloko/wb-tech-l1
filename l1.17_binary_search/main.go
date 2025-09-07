package main

import (
	"fmt"
	"sort"
)

// using sort package
// returns the index if found, otherwise -1
func BinarySearchUsingSort(n int, nums []int) int {
	i := sort.Search(len(nums), func(i int) bool {
		return nums[i] >= n
	})
	if i < len(nums) && nums[i] == n {
		return i
	}
	return -1
}

// handmade search
// returns the index if found, otherwise -1
func BinarySearchManual(n int, nums []int) int {
	left, right := 0, len(nums)-1

	for left <= right {
		mid := (left + right) / 2

		if nums[mid] < n {
			left = mid + 1
		} else if nums[mid] > n {
			right = mid - 1
		} else {
			return mid
		}
	}

	return -1
}

func main() {
	data := []int{-10000, -13, 0, 1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 1000000}

	fmt.Println(BinarySearchManual(-13, data))
	fmt.Println(BinarySearchUsingSort(-13, data))
}
