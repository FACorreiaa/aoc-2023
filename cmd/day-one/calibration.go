package dayone

import (
	"fmt"
	"github.com/FACorreiaa/aoc-2023/cmd/common"
	tea "github.com/charmbracelet/bubbletea"
	"strconv"
	"strings"
	"unicode"
)

type MsgDayOne struct {
	Title  string
	Result int
}

var (
	lookup = []string{"one", "two", "three", "four", "five", "six", "seven", "eight", "nine"}
)

func partOne(lines []string) int {
	numbers := extractNumbers(lines)
	return common.Sum(numbers)
}

func partTwo(lines []string) int {
	lines = replaceAlphaNumbers(lines)
	numbers := extractNumbers(lines)
	return common.Sum(numbers)
}

func replaceAlphaNumbers(lines []string) []string {
	var output []string
	for _, line := range lines {
		output = append(output, replaceAlphaNumber(line))
	}

	return output
}

func replaceAlphaNumber(line string) string {
	for index, word := range lookup {
		line = strings.ReplaceAll(line, word, fmt.Sprintf("%s%d%s", word, index+1, word))
	}

	return line
}

func extractNumbers(lines []string) []int {
	var numbers []int
	for _, line := range lines {
		numbers = append(numbers, extractNumber(line))
	}

	return numbers
}

func extractNumber(line string) int {
	var digits []rune
	for _, digit := range line {
		if unicode.IsNumber(digit) {
			digits = append(digits, digit)
		}
	}

	num, err := strconv.Atoi(string([]rune{digits[0], digits[len(digits)-1]}))
	if err != nil {
		common.HandleError(err, "Error getting number")
	}

	return num
}

func Start() tea.Msg {
	lines := common.GetLines("./cmd/day-one/calibration.txt")
	//for _, line := range lines {
	//	println(line)
	//}
	//partOneStart := time.Now()
	partOneResult := partOne(lines)
	//log.Print("Result: ", partOneResult)
	//log.Print("\nDay one part one took: ", time.Since(partOneStart))
	//partTwoStart := time.Now()
	//partTwoResult := partTwo(lines)
	//log.Print("Result: ", partTwoResult)
	//log.Print("\nDay one part two took: ", time.Since(partTwoStart))
	return MsgDayOne{Title: "Day 1", Result: partOneResult}

}
