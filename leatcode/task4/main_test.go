package main

import (
	"github.com/stretchr/testify/require"
	"testing"
)

func Test_SearchInsert(t *testing.T) {
	tests := []struct {
		in     []int
		out    int
		target int
		name   string
	}{
		{
			in:     []int{1, 3, 4, 5},
			out:    1,
			target: 2,
			name:   "Positive scenario",
		},
	}

	for _, test := range tests {
		require.Equal(t, test.out, SearchInsert(test.in, test.target), test.name)
	}
}
