package dayone

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

// TestPartOne tests the PartOne function
func TestPartOne(t *testing.T) {

	partOne := PartOne

	assert.Equal(t, 42, partOne)
}

// TestPartTwo tests the PartTwo function
func TestPartTwo(t *testing.T) {

	partTwo := PartTwo
	assert.Equal(t, 42, partTwo)
}

// BenchmarkPartOne benchmarks the PartOne function
func BenchmarkPartOne(b *testing.B) {
	for i := 0; i < b.N; i++ {

		PartOne()
	}
}

// BenchmarkPartTwo benchmarks the PartTwo function
func BenchmarkPartTwo(b *testing.B) {
	for i := 0; i < b.N; i++ {
		PartTwo()
	}
}
