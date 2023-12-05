package daythree

import (
	"github.com/stretchr/testify/assert"
	"os"
	"testing"
)

func TestPartOne(t *testing.T) {
	tests := []struct {
		expected int
		input    string
		//fn       func(string) int
	}{
		{
			4361,
			`gear_test.txt`,
		},
	}

	for _, test := range tests {
		b, err := os.ReadFile(test.input)
		assert.NoError(t, err, test.input)
		assert.Equal(t, test.expected, partOne(string(b)))
	}
}