package daynine

import (
	"github.com/FACorreiaa/aoc-2023/cmd/settings"
	"log"
	"regexp"
	"slices"
	"strconv"
	"strings"
	"time"
)

var items = regexp.MustCompile(`(\w+) = \((\w+), (\w+)\)`)

func parseFile(data []byte) (sequences [][]int64) {
	sequences = make([][]int64, 0)

	for i, line := range strings.Split(string(data), "\n") {
		sequences = append(sequences, make([]int64, 0))

		numbers := strings.Split(line, " ")

		for _, n := range numbers {
			x, _ := strconv.Atoi(n)
			sequences[i] = append(sequences[i], int64(x))
		}
	}
	return
}

func partOne(sequences [][]int64) int64 {
	var sum int64 = 0

	for _, s := range sequences {
		// get amount to add
		diffs := getSequence(s)
		slices.Reverse(diffs)

		var total int64 = 0
		for i := 0; i < len(diffs); i++ {
			total += diffs[i][len(diffs[i])-1]
		}
		sum += s[len(s)-1] + total

	}

	return sum

}

func getSequence(numbers []int64) (allDifferences [][]int64) {
	//degree
	var differences = numbers

	allDifferences = make([][]int64, 0)

	for {
		if areAllSame(differences) {
			break
		}

		differences = getDifference(differences)
		allDifferences = append(allDifferences, differences)
	}
	return
}

func areAllSame(numbers []int64) bool {
	x := numbers[0]
	for _, v := range numbers {
		if v != x {
			return false
		}
	}
	return true
}

func getDifference(numbers []int64) []int64 {
	result := make([]int64, 0)
	for i := 1; i < len(numbers); i++ {
		result = append(result, numbers[i]-numbers[i-1])
	}
	return result
}

func partTwo(sequences [][]int64) int64 {
	var sum int64 = 0

	for _, s := range sequences {
		diffs := getSequence(s)
		slices.Reverse(diffs)

		var total int64 = 0

		for i := 0; i < len(diffs); i++ {
			total = diffs[i][0] - total
		}

		sum += s[0] - total
	}

	return sum
}

func Start() {
	lines := settings.GetLines("./cmd/day-nine/pipe.txt")
	for _, line := range lines {
		println(line)
	}
	partOneStart := time.Now()
	partOneResult := partOne(parseFile([]byte(strings.Join(lines, "\n"))))
	log.Print("Result: ", partOneResult)
	log.Print("Day seven part one took: ", time.Since(partOneStart))

	partTwoStart := time.Now()
	partTwoResult := partTwo(parseFile([]byte(strings.Join(lines, "\n"))))
	log.Print("Result: ", partTwoResult)
	log.Print("Day seven part two took: ", time.Since(partTwoStart))
}
