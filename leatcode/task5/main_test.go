package main

import (
	"github.com/stretchr/testify/require"
	"testing"
)

func TestLengthOfLastWord(t *testing.T) {
	tests := []struct {
		in   string
		out  int
		name string
	}{
		{
			in:   "Hello World ",
			out:  5,
			name: "One space end at the string",
		},
	}
	for _, test := range tests {
		require.Equal(t, test.out, LengthOfLastWord(test.in), test.name)
	}
}
