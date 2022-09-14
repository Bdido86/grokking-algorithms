package main

import (
	"fmt"
)

func main() {
	list := []int{4, 8, 6, 2, 1, 0, 7, 5, 3, -2, 1, 9}
	fmt.Printf("%+v", mergeSort(list))
}

func mergeSort(items []int) []int {
	length := len(items)
	if length < 2 {
		return items
	}

	middle := length / 2
	leftSorted := mergeSort(items[:middle])
	rightSorted := mergeSort(items[middle:])

	leftLength := len(leftSorted)
	rightLength := len(rightSorted)
	leftI := 0
	rightI := 0
	sortedItems := make([]int, 0, length)
	for i := 0; i <= length; i++ {
		if leftLength > 0 && rightLength > 0 {
			if leftSorted[leftI] <= rightSorted[rightI] {
				sortedItems = append(sortedItems, leftSorted[leftI])
				leftLength--
				leftI++
				continue
			}

			sortedItems = append(sortedItems, rightSorted[rightI])
			rightLength--
			rightI++
			continue
		}

		if leftLength > 0 {
			sortedItems = append(sortedItems, leftSorted[leftI])
			leftLength--
			leftI++
			continue
		}

		if rightLength > 0 {
			sortedItems = append(sortedItems, rightSorted[rightI])
			rightLength--
			rightI++
			continue
		}
	}

	return sortedItems
}
