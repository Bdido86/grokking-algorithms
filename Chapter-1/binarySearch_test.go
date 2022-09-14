package main

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_BinarySearch(t *testing.T) {
	t.Run("errors", func(t *testing.T) {
		t.Parallel()
		t.Run("slice is nil", func(t *testing.T) {
			var items []int

			idx, err := binarySearch(items, 5)

			assert.EqualError(t, errorSliceIsNil, err.Error())
			assert.Equal(t, idx, uint(0))
		})
		t.Run("slice is empty", func(t *testing.T) {
			items := make([]int, 0)

			idx, err := binarySearch(items, 5)

			assert.EqualError(t, errorSliceIsEmpty, err.Error())
			assert.Equal(t, idx, uint(0))
		})
		t.Run("not found 1", func(t *testing.T) {
			items := []int{1, 3, 4, 5, 7}

			idx, err := binarySearch(items, 10)

			assert.EqualError(t, errorNotFound, err.Error())
			assert.Equal(t, uint(0), idx)
		})
		t.Run("not found 2", func(t *testing.T) {
			items := []int{-8, -5, 0, 1, 230}

			idx, err := binarySearch(items, 231)

			assert.EqualError(t, errorNotFound, err.Error())
			assert.Equal(t, uint(0), idx)
		})
		t.Run("not found 3", func(t *testing.T) {
			items := []int{1}

			idx, err := binarySearch(items, 2)

			assert.EqualError(t, errorNotFound, err.Error())
			assert.Equal(t, uint(0), idx)
		})
		t.Run("not found 4", func(t *testing.T) {
			items := []int{10, 20, 30, 40, 50, 60, 70}

			idx, err := binarySearch(items, 11)

			assert.EqualError(t, errorNotFound, err.Error())
			assert.Equal(t, uint(0), idx)
		})
	})

	t.Run("success", func(t *testing.T) {
		t.Parallel()
		t.Run("negative slice", func(t *testing.T) {
			items := []int{-100, -45, -34, -33, -29, -20, -10, -5, -1}

			idx, err := binarySearch(items, -5)

			assert.Nil(t, err)
			assert.Equal(t, uint(7), idx)
		})
		t.Run("positive slice", func(t *testing.T) {
			items := []int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}

			idx, err := binarySearch(items, 2)

			assert.Nil(t, err)
			assert.Equal(t, uint(2), idx)
		})
		t.Run("mixed slice", func(t *testing.T) {
			items := []int{-5, -4, -1, 1, 2, 5, 6, 7, 9, 11}

			idx, err := binarySearch(items, 5)

			assert.Nil(t, err)
			assert.Equal(t, uint(5), idx)
		})
		t.Run("right border", func(t *testing.T) {
			items := []int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}

			idx, err := binarySearch(items, 9)

			assert.Nil(t, err)
			assert.Equal(t, uint(9), idx)
		})
		t.Run("left border", func(t *testing.T) {
			items := []int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}

			idx, err := binarySearch(items, 0)

			assert.Nil(t, err)
			assert.Equal(t, uint(0), idx)
		})
		t.Run("one element", func(t *testing.T) {
			items := []int{1}

			idx, err := binarySearch(items, 1)

			assert.Nil(t, err)
			assert.Equal(t, uint(0), idx)
		})
		t.Run("two elements", func(t *testing.T) {
			items := []int{1, 2}

			idx, err := binarySearch(items, 2)

			assert.Nil(t, err)
			assert.Equal(t, uint(1), idx)
		})
	})
}
