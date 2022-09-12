package main

import (
	"fmt"
	"github.com/pkg/errors"
)

var (
	errorSliceIsNil = errors.New("slice is nil")
)

func main() {
	mas := []int{4, 4, 4}
	items, _ := selectionSort(mas)

	fmt.Printf("%+v", items)
}

// O(n*n) = O(n * 1/2 * n)
func selectionSort(items []int) ([]int, error) {
	if items == nil {
		return nil, errorSliceIsNil
	}

	if len(items) == 0 || len(items) == 1 {
		return items, nil
	}

	minKey := func(key uint, items []int) uint {
		minValue := items[key]
		lenItems := uint(len(items))
		for i := key + 1; i < lenItems; i++ {
			if items[i] < minValue {
				minValue = items[i]
				key = i
			}
		}
		return key
	}

	for j := 0; j < len(items); j++ {
		key := minKey(uint(j), items)
		keyValue := items[key]

		if key != uint(j) {
			items[key] = items[j]
			items[j] = keyValue
		}
	}

	return items, nil
}
