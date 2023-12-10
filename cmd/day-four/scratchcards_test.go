package dayfour

import (
	"github.com/FACorreiaa/aoc-2023/cmd/common"
	"strings"
	"testing"
)

//func TestPartOne(t *testing.T) {
//	tests := []struct {
//		expected int
//		input    string
//		fn       func(string) int
//	}{
//
//		{
//			13,
//			`scratchcards_test_one.txt`,
//			partOne,
//		},
//		{
//			30,
//			`scratchcards_test_two.txt`,
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

func BenchmarkPartOne(b *testing.B) {
	var r int

	lines := common.GetLines("scratchcards.txt")

	for n := 0; n < b.N; n++ {
		r = partOne(strings.Join(lines, "\n"))
	}
	result = r

}

func BenchmarkPartTwo(b *testing.B) {
	var r int

	lines := common.GetLines("scratchcards.txt")

	for n := 0; n < b.N; n++ {
		r = partTwo(strings.Join(lines, "\n"))
	}
	result = r
}
