package main

import (
	"github.com/stretchr/testify/require"
	"testing"
)

func Test_TrimString(t *testing.T) {
	tests := []struct {
		in     string
		output string
		name   string
	}{
		{
			in:     "String + space",
			output: "S12e",
			name:   "Positive scenario + space",
		},
		{
			in:     "東東東東東東東東東東東",
			output: "東9東",
			name:   "Positive scenario use rune",
		},
		{
			in:     "localization",
			output: "l10n",
			name:   "Positive scenario",
		},
	}
	for _, test := range tests {
		require.Equal(t, test.output, TrimString(test.in), test.name)
	}

}
