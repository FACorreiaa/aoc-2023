package dayten

import (
	"github.com/FACorreiaa/aoc-2023/cmd/settings"
	"testing"
)

//func TestPartOne(t *testinT) {
//	tests := []struct {
//		expected int
//		input    string
//		fn       func([][]int64) int64
//	}{
//
//		{
//			8,
//			`pipe_test_one.txt`,
//			partOne,
//		},
//		{
//			68,
//			`pipe_test_two.txt`,
//			partOne,
//		},
//	}
//
//	for _, test := range tests {
//		b, err := os.ReadFile(test.input)
//		assert.NoError(t, err, test.input)
//		sequences := parseFile([]byte(strings.Join(strings.Split(string(b), "\n"), " ")))
//		assert.Equal(t, test.expected, test.fn(sequences))
//
//	}
//}

var result int64

func BenchmarkPartOne(b *testing.B) {
	var r int64

	lines := settings.GetLines("pipe.txt")

	for n := 0; n < b.N; n++ {
		r = int64(partOne(lines))
	}
	result = r

}

func BenchmarkPartTwo(b *testing.B) {
	var r int64

	lines := settings.GetLines("pipe.txt")

	for n := 0; n < b.N; n++ {
		r = int64(partOne(lines))
	}
	result = r
}
