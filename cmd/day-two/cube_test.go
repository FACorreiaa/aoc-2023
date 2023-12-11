package daytwo

import (
	"github.com/FACorreiaa/aoc-2023/cmd/settings"
	"testing"
)

//func TestPartOne(t *testing.T) {
//	tests := []struct {
//		expected int
//		input    string
//	}{
//		{
//			8,
//			`cube_test_one.txt`,
//		},
//	}
//
//	for _, test := range tests {
//		b, err := os.ReadFile(test.input)
//		lines := strings.Split(string(b), "\n")
//		numbers := parseGames(lines)
//		assert.NoError(t, err, test.input)
//		assert.Equal(t, test.expected, partOne(numbers))
//	}
//}
//
//func TestPartTwo(t *testing.T) {
//	tests := []struct {
//		expected int
//		input    string
//	}{
//		{
//			2286,
//			`cube_test_two.txt`,
//		},
//	}
//
//	for _, test := range tests {
//		b, err := os.ReadFile(test.input)
//		lines := strings.Split(string(b), "\n")
//		games := parseGames(lines)
//		assert.NoError(t, err, test.input)
//		assert.Equal(t, test.expected, partTwo(games))
//	}
//}

var result int

func BenchmarkPartOne(b *testing.B) {
	var r int

	lines := settings.GetLines("cube.txt")

	for n := 0; n < b.N; n++ {
		numbers := parseGames(lines)
		r = partTwo(numbers)
	}

	result = r
}

func BenchmarkPartTwo(b *testing.B) {
	var r int

	lines := settings.GetLines("cube.txt")

	for n := 0; n < b.N; n++ {
		numbers := parseGames(lines)
		r = partOne(numbers)
	}

	result = r
}
