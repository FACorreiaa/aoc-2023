package day_one

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
	"unicode"
)

var (
	lookup = []string{"one", "two", "three", "four", "five", "six", "seven", "eight", "nine"}
)

func PartOne() {
	lines := getLines(os.Args[1])
	numbers := extractNumbers(lines)
	fmt.Println(sum(numbers))
}

func PartTwo() {
	lines := getLines(os.Args[1])
	lines = replaceAlphaNumbers(lines)
	numbers := extractNumbers(lines)
	fmt.Println(sum(numbers))
}

func sum(nums []int) int {
	var result int
	for _, i := range nums {
		result += i
	}

	return result
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
		panic(err)
	}

	return num
}

func getLines(filename string) []string {
	file, err := os.Open(filename)
	if err != nil {
		panic(err)
	}

	defer file.Close()

	var lines []string
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}

	if scanner.Err() != nil {
		panic(err)
	}

	return lines
}
