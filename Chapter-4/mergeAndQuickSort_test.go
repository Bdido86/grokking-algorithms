package main

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"golang.org/x/exp/slices"
	"testing"
)

func Test_mergeSort(t *testing.T) {
	cases := []struct {
		arg  []int
		want []int
	}{
		{
			[]int{},
			[]int{},
		},
		{
			[]int{2},
			[]int{2},
		},
		{
			[]int{7, 2},
			[]int{2, 7},
		},
		{
			[]int{9, 8, 7, 6, 5, 4, 3, 2, 1},
			[]int{1, 2, 3, 4, 5, 6, 7, 8, 9},
		},
		{
			[]int{300, -4, 4, 0, 55},
			[]int{-4, 0, 4, 55, 300},
		},
		{
			[]int{-3, -6, -8},
			[]int{-8, -6, -3},
		},
	}

	for _, cs := range cases {
		t.Run(fmt.Sprintf("%+v = %+v", cs.arg, cs.want), func(t *testing.T) {
			mergedSort := mergeSort(cs.arg)
			assert.Equal(t, true, slices.Equal(cs.want, mergedSort))
		})
	}
}

func Test_quickSort(t *testing.T) {
	cases := []struct {
		arg  []int
		want []int
	}{
		{
			[]int{},
			[]int{},
		},
		{
			[]int{2},
			[]int{2},
		},
		{
			[]int{7, 2},
			[]int{2, 7},
		},
		{
			[]int{9, 8, 7, 6, 5, 4, 3, 2, 1},
			[]int{1, 2, 3, 4, 5, 6, 7, 8, 9},
		},
		{
			[]int{300, -4, 4, 0, 55},
			[]int{-4, 0, 4, 55, 300},
		},
		{
			[]int{-3, -6, -8},
			[]int{-8, -6, -3},
		},
	}

	for _, cs := range cases {
		t.Run(fmt.Sprintf("%+v = %+v", cs.arg, cs.want), func(t *testing.T) {
			quickedSort := quickSort(cs.arg)
			assert.Equal(t, true, slices.Equal(cs.want, quickedSort))
		})
	}
}
