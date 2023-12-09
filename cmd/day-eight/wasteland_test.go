package dayeight

import (
	"github.com/stretchr/testify/assert"
	"os"
	"testing"
)

func TestPartOne(t *testing.T) {
	tests := []struct {
		expected int
		input    string
		fn       func(string) int64
	}{

		{
			6,
			`mirage_test_one.txt`,
			partOne,
		},
		{
			2,
			`mirage_test_two.txt`,
			partOne,
		},
		{
			6,
			`mirage_test_three.txt`,
			partOne,
		},
	}

	for _, test := range tests {
		b, err := os.ReadFile(test.input)
		assert.NoError(t, err, test.input)
		assert.Equal(t, test.expected, test.fn(string(b)))

	}
}
