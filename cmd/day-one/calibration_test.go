package dayone

import (
	"github.com/FACorreiaa/aoc-2023/cmd/common"
	"github.com/stretchr/testify/assert"
	"os"
	"strings"
	"testing"
)

func TestPartOne(t *testing.T) {
	tests := []struct {
		expected int
		input    string
	}{
		{
			142,
			`calibration_test_one.txt`,
		},
	}

	for _, test := range tests {
		b, err := os.ReadFile(test.input)
		lines := strings.Split(string(b), "\n")
		numbers := extractNumbers(lines)
		expectedSum := common.Sum(numbers)
		assert.NoError(t, err, test.input)
		assert.Equal(t, test.expected, expectedSum)
	}
}

func TestPartTwo(t *testing.T) {
	tests := []struct {
		expected int
		input    string
	}{
		{
			281,
			`calibration_test_two.txt`,
		},
	}

	for _, test := range tests {
		b, err := os.ReadFile(test.input)
		lines := strings.Split(string(b), "\n")
		numbers := extractNumbers(lines)
		expectedSum := common.Sum(numbers)
		assert.NoError(t, err, test.input)
		assert.Equal(t, test.expected, expectedSum)
	}
}

func BenchmarkPartOne(b *testing.B) {
	lines := common.GetLines("calibration.txt")

	for i := 0; i < b.N; i++ {
		partOne(lines)
	}
}

func BenchmarkPartTwo(b *testing.B) {
	lines := common.GetLines("calibration.txt")

	for i := 0; i < b.N; i++ {
		partTwo(lines)
	}
}
