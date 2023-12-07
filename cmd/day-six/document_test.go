package daysix

import (
	"github.com/stretchr/testify/assert"
	"os"
	"testing"
)

func TestPartOne(t *testing.T) {
	tests := []struct {
		expected int
		input    string
		fn       func(string) int
	}{

		{
			288,
			`document_test_one.txt`,
			partOne,
		},
		{
			71516,
			`document_test_two.txt`,
			partTwo,
		},
	}

	for _, test := range tests {
		b, err := os.ReadFile(test.input)
		assert.NoError(t, err, test.input)
		assert.Equal(t, test.expected, test.fn(string(b)))
	}
}
