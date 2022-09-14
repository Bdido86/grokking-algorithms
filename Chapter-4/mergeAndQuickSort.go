package main

import (
	"fmt"
)

func main() {
	list := []int{4, 8, 6, 2, 1, 0, 7, 5, 3, -2, 1, 9}
	fmt.Printf("%+v\n\n", mergeSort(list))

	list1 := []int{4, 8, 6, 2, 1, 0, 7, 5, 3, -2, 1, 9}
	fmt.Printf("%+v", quickSort(list1))
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

func quickSort(items []int) []int {
	length := len(items)
	if length < 2 {
		return items
	}

	pivot := items[0]
	leftSorted := make([]int, 0, 2)
	rightSorted := make([]int, 0, 2)

	for _, value := range items[1:] {
		if value <= pivot {
			leftSorted = append(leftSorted, value)
		} else {
			rightSorted = append(rightSorted, value)
		}
	}

	return append(
		append(
			append([]int(nil), quickSort(leftSorted)...),
			pivot), quickSort(rightSorted)...,
	)
}
