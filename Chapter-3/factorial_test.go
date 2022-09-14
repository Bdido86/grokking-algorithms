package main

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestFactorial(t *testing.T) {
	t.Run("errors", func(t *testing.T) {
		t.Parallel()

		cases := []struct {
			number uint
			want   uint
		}{
			{
				0,
				0,
			},
			{
				1,
				2,
			},
		}

		for _, cs := range cases {
			t.Run(fmt.Sprintf("factorial(%d) <> %d", cs.number, cs.want), func(t *testing.T) {
				want := factorial(cs.number)
				assert.NotEqual(t, cs.want, want)
			})
		}
	})

	t.Run("success", func(t *testing.T) {
		t.Parallel()
		cases := []struct {
			number uint
			want   uint
		}{
			{0, 1},
			{1, 1},
			{2, 2},
			{3, 6},
			{4, 24},
			{5, 120},
		}

		for _, cs := range cases {
			t.Run(fmt.Sprintf("factorial(%d) = %d", cs.number, cs.want), func(t *testing.T) {
				want := factorial(cs.number)
				assert.Equal(t, cs.want, want)
			})
		}
	})
}
