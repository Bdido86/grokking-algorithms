package main

import (
	"github.com/stretchr/testify/assert"
	"golang.org/x/exp/slices"
	"testing"
)

func Test_SelectionSort(t *testing.T) {
	t.Run("errors", func(t *testing.T) {
		t.Parallel()
		t.Run("slice is nil", func(t *testing.T) {
			var items []int

			sortedItems, err := selectionSort(items)

			assert.EqualError(t, errorSliceIsNil, err.Error())
			assert.Nil(t, sortedItems)
		})
	})

	t.Run("success", func(t *testing.T) {
		t.Parallel()
		cases := []struct {
			name string
			arg  []int
			want []int
		}{
			{
				"slice is empty",
				[]int{},
				[]int{},
			},
			{
				"one element",
				[]int{1},
				[]int{1},
			},
			{
				"two element",
				[]int{2, 1},
				[]int{1, 2},
			},
			{
				"positive elements",
				[]int{2, 0, 5, 1, 7, 9, 6, 3, 4, 8},
				[]int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9},
			},
			{
				"negative elements",
				[]int{-1, -2, -3, -4, -8},
				[]int{-8, -4, -3, -2, -1},
			},
			{
				"mix elements",
				[]int{-50, -70, 0, -5, 1, 20, -2, -2, 1, 0},
				[]int{-70, -50, -5, -2, -2, 0, 0, 1, 1, 20},
			},
			{
				"identical elements",
				[]int{2, 2, 2, 2, 2},
				[]int{2, 2, 2, 2, 2},
			},
		}

		for _, cs := range cases {
			t.Run(cs.name, func(t *testing.T) {

				sortedItems, err := selectionSort(cs.arg)

				assert.Nil(t, err)
				assert.Equal(t, true, slices.Equal(cs.want, sortedItems))
			})
		}
	})
}
