package main

import (
	"github.com/stretchr/testify/require"
	"testing"
)

func Test_Finalist(t *testing.T) {
	tests := []struct {
		in     []int
		output int
		name   string
	}{
		{
			in:     []int{12, 9, 8, 8, 7, 7, 5, 5, 5},
			output: 6,
			name:   "Positive scenario",
		},
		{
			in:     []int{1, 1, 0, 0},
			output: 2,
			name:   "Positive scenario",
		},
	}
	for _, test := range tests {
		require.Equal(t, test.output, Finalist(test.in), test.name)
	}

}
