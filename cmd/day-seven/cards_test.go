package dayseven

import (
	"github.com/stretchr/testify/assert"
	"os"
	"testing"
)

func TestPartOne(t *testing.T) {
	tests := []struct {
		expected int
		input    string
		jokers   bool
		fn       func(string, bool) int64
	}{

		{
			6440,
			`cards_test_one.txt`,
			false,
			partOne,
		},
		{
			71516,
			`cards_test_two.txt`,
			true,
			partOne,
		},
	}

	for _, test := range tests {
		b, err := os.ReadFile(test.input)
		assert.NoError(t, err, test.input)
		assert.Equal(t, test.expected, test.fn(string(b), false))
		assert.Equal(t, test.expected, test.fn(string(b), true))

	}
}
