package main

import (
	"github.com/pkg/errors"
)

var (
	errorSliceIsNil   = errors.New("slice is nil")
	errorSliceIsEmpty = errors.New("slice is empty")
	errorNotFound     = errors.New("not found")
)

// O(log n), n - len(items)
func binarySearch(items []int, item int) (uint, error) {
	if items == nil {
		return 0, errorSliceIsNil
	}

	if len(items) == 0 {
		return 0, errorSliceIsEmpty
	}

	low := 0
	high := len(items) - 1
	for low <= high {
		mid := (low + high) / 2
		found := items[mid]

		if found == item {
			return uint(mid), nil
		}

		if found > item {
			high = mid - 1
		} else {
			low = mid + 1
		}
	}

	return 0, errorNotFound
}
