package dayfive

import (
	"github.com/FACorreiaa/aoc-2023/cmd/settings"
	"strings"
	"testing"
)

//func TestPartOne(t *testing.T) {
//	tests := []struct {
//		expected int64
//		input    string
//		fn       func(string) int64
//	}{
//
//		{
//			35,
//			`mirage_test_one.txt`,
//			partOne,
//		},
//		{
//			46,
//			`mirage_test_two.txt`,
//			partTwo,
//		},
//	}
//
//	for _, test := range tests {
//		b, err := os.ReadFile(test.input)
//		assert.NoError(t, err, test.input)
//		assert.Equal(t, test.expected, test.fn(string(b)))
//	}
//}

var result int64

func BenchmarkPartOne(b *testing.B) {
	var r int64

	lines := settings.GetLines("seeds.txt")

	for n := 0; n < b.N; n++ {
		r = partOne(strings.Join(lines, "\n"))
	}
	result = r

}

func BenchmarkPartTwo(b *testing.B) {
	var r int64

	lines := settings.GetLines("seeds.txt")

	for n := 0; n < b.N; n++ {
		r = partTwo(strings.Join(lines, "\n"))
	}
	result = r
}
