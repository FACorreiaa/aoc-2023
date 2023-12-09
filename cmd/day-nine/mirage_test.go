package daynine

import (
	"github.com/stretchr/testify/assert"
	"os"
	"strings"
	"testing"
)

func TestPartOne(t *testing.T) {
	tests := []struct {
		expected int
		input    string
		fn       func([][]int64) int64
	}{

		{
			28,
			`mirage_test_one.txt`,
			partOne,
		},
		{
			68,
			`mirage_test_two.txt`,
			partOne,
		},
		{
			-2,
			`mirage_test_three.txt`,
			partTwo,
		},
	}

	for _, test := range tests {
		b, err := os.ReadFile(test.input)
		assert.NoError(t, err, test.input)
		sequences := parseFile([]byte(strings.Join(strings.Split(string(b), "\n"), " ")))
		assert.Equal(t, test.expected, test.fn(sequences))

	}
}
