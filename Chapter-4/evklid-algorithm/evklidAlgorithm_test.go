package main

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_findNOD(t *testing.T) {
	cases := []struct {
		arg  []int
		want int
	}{
		{
			[]int{10},
			10,
		},
		{
			[]int{7, -7},
			7,
		},
		{
			[]int{64, 48},
			16,
		},
		{
			[]int{111, -432},
			3,
		},
		{
			[]int{661, 113},
			1,
		},
		{
			[]int{48, 72, -288},
			24,
		},
		{
			[]int{78, -294, 570, 36},
			6,
		},
		{
			[]int{345, 775, 570, 95},
			5,
		},
		{
			[]int{345, 775, 1},
			1,
		},
	}

	for _, cs := range cases {
		t.Run(fmt.Sprintf("NOD(%+v) = %d", cs.arg, cs.want), func(t *testing.T) {

			NOD := findNOD(cs.arg...)

			assert.Equal(t, cs.want, NOD)
		})
	}
}
