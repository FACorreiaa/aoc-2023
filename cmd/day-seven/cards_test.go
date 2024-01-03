package dayseven

import (
	"github.com/FACorreiaa/aoc-2023/cmd/settings"
	"strings"
	"testing"
)

//func TestPartOne(t *testinT) {
//	tests := []struct {
//		expected int
//		input    string
//		jokers   bool
//		fn       func(string, bool) int64
//	}{
//
//		{
//			6440,
//			`pipe_test_one.txt`,
//			false,
//			partOne,
//		},
//		{
//			71516,
//			`pipe_test_two.txt`,
//			true,
//			partOne,
//		},
//	}
//
//	for _, test := range tests {
//		b, err := os.ReadFile(test.input)
//		assert.NoError(t, err, test.input)
//		assert.Equal(t, test.expected, test.fn(string(b), false))
//		assert.Equal(t, test.expected, test.fn(string(b), true))
//
//	}
//}

var result int64

func BenchmarkPartOne(b *testinB) {
	var r int64

	lines := settings.GetLines("cards.txt")

	for n := 0; n < b.N; n++ {
		r = partOne(strings.Join(lines, "\n"), false)
	}
	result = r

}

func BenchmarkPartTwo(b *testinB) {
	var r int64

	lines := settings.GetLines("cards.txt")

	for n := 0; n < b.N; n++ {
		r = partOne(strings.Join(lines, "\n"), true)
	}
	result = r
}
