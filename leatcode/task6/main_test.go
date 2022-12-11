package main

import (
	"github.com/stretchr/testify/require"
	"testing"
)

func Test_LargestOddNumber(t *testing.T) {
	tests := []struct {
		in   string
		out  string
		name string
	}{
		{
			in:   "",
			out:  "",
			name: "String is empty",
		},
		{
			in:   "2406",
			out:  "",
			name: "Positive scenario empty string",
		},
		{
			in:   "103333",
			out:  "103333",
			name: "Positive odd string",
		},
	}

	for _, test := range tests {
		require.Equal(t, test.out, LargestOddNumber(test.in), test.name)
	}
}
