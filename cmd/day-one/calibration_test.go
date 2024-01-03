package dayone

import (
	"github.com/FACorreiaa/aoc-2023/cmd/settings"
	"github.com/stretchr/testify/assert"
	"os"
	"strings"
	"testing"
)

func TestPartOne(t *testinT) {
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
		expectedSum := settings.Sum(numbers)
		assert.NoError(t, err, test.input)
		assert.Equal(t, test.expected, expectedSum)
	}
}

//func TestPartTwo(t *testinT) {
//	tests := []struct {
//		expected int
//		input    string
//	}{
//		{
//			281,
//			`calibration_test_two.txt`,
//		},
//	}
//
//	for _, test := range tests {
//		b, err := os.ReadFile(test.input)
//		lines := strings.Split(string(b), "\n")
//		numbers := extractNumbers(lines)
//		expectedSum := settings.Sum(numbers)
//		assert.NoError(t, err, test.input)
//		assert.Equal(t, test.expected, expectedSum)
//	}
//}

var result int

func BenchmarkPartOne(b *testinB) {
	var r int

	lines := settings.GetLines("calibration.txt")

	for n := 0; n < b.N; n++ {
		r = partOne(lines)
	}
	result = r

}

func BenchmarkPartTwo(b *testinB) {
	var r int

	lines := settings.GetLines("calibration.txt")

	for n := 0; n < b.N; n++ {
		r = partTwo(lines)
	}
	result = r
}
