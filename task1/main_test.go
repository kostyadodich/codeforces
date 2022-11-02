package main

import (
	"github.com/stretchr/testify/require"
	"testing"
)

func Test_Remainder(t *testing.T) {
	tests := []struct {
		input  int
		output string
		name   string
	}{
		{
			input:  8,
			output: "YES",
			name:   "Positive scenario",
		},
		{
			input:  5,
			output: "NO",
			name:   "Positive scenario fail",
		},
		{
			input:  99999,
			output: "NO",
			name:   "Positive scenario over integer",
		},
	}
	for _, test := range tests {
		require.Equal(t, test.output, Remainder(test.input), test.name)
	}
}
