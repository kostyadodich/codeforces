package main

import (
	"github.com/stretchr/testify/require"
	"testing"
)

func Test_PlusOne(t *testing.T) {
	tests := []struct {
		in  []int
		out []int
	}{
		{
			in:  []int{9, 9},
			out: []int{1, 0, 0},
		},
		{
			in:  []int{1, 9, 9},
			out: []int{2, 0, 0},
		},
	}

	for _, test := range tests {
		require.Equal(t, test.out, PlusOne(test.in))
	}

}
