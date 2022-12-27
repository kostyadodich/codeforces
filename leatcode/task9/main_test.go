package main

import (
	"github.com/stretchr/testify/require"
	"testing"
)

func Test_MySqrt(t *testing.T) {
	tests := []struct {
		in   int
		out  int
		name string
	}{
		{
			in:   256,
			out:  16,
			name: "Positive scenario",
		},
		{
			in:  4,
			out: 2,
		},
	}

	for _, test := range tests {
		require.Equal(t, test.out, MySqrt(test.in), test.name)
	}
}
