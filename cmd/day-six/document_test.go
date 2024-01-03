package daysix

import (
	"github.com/FACorreiaa/aoc-2023/cmd/settings"
	"strings"
	"testing"
)

//func TestPartOne(t *testinT) {
//	tests := []struct {
//		expected int
//		input    string
//		fn       func(string) int
//	}{
//
//		{
//			288,
//			`pipe_test_one.txt`,
//			partOne,
//		},
//		{
//			71516,
//			`pipe_test_two.txt`,
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

var result int

func BenchmarkPartOne(b *testinB) {
	var r int

	lines := settings.GetLines("document.txt")

	for n := 0; n < b.N; n++ {
		r = partOne(strings.Join(lines, "\n"))
	}
	result = r

}

func BenchmarkPartTwo(b *testinB) {
	var r int

	lines := settings.GetLines("document.txt")

	for n := 0; n < b.N; n++ {
		r = partTwo(strings.Join(lines, "\n"))
	}
	result = r
}
