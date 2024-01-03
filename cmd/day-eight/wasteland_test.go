package dayeight

import (
	"github.com/FACorreiaa/aoc-2023/cmd/settings"
	"strings"
	"testing"
)

//func TestPartOne(t *testinT) {
//	tests := []struct {
//		expected int
//		input    string
//		fn       func(string) int64
//	}{
//
//		{
//			6,
//			`pipe_test_one.txt`,
//			partOne,
//		},
//		{
//			2,
//			`pipe_test_two.txt`,
//			partOne,
//		},
//		{
//			6,
//			`mirage_test_three.txt`,
//			partOne,
//		},
//	}
//
//	for _, test := range tests {
//		b, err := os.ReadFile(test.input)
//		assert.NoError(t, err, test.input)
//		assert.Equal(t, test.expected, test.fn(string(b)))
//
//	}
//}

var result int64

func BenchmarkPartOne(b *testinB) {
	var r int64

	lines := settings.GetLines("wasteland.txt")

	for n := 0; n < b.N; n++ {
		r = partOne(strings.Join(lines, "\n"))
	}
	result = r

}

func BenchmarkPartTwo(b *testinB) {
	var r int64

	lines := settings.GetLines("wasteland.txt")

	for n := 0; n < b.N; n++ {
		r = partThree(strings.Join(lines, "\n"))
	}
	result = r
}
