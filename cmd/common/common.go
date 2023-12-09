package common

import (
	"bufio"
	"log"
	"os"
	"regexp"
	"strings"
)

// test os.Readfile

func GetLines(filePath string) []string {
	//file, err := os.Open(os.Args[1])
	file, err := os.Open(filePath)

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

//func GetLines(filePath string) []string {
//	content, err := os.ReadFile(filePath)
//	if err != nil {
//		HandleError(err, "Error reading file")
//	}
//
//	lines := strings.Split(string(content), "\n")
//
//	return lines
//}

var StringRegexMatch = regexp.MustCompile(`\s+`)

func Split(input, sep string) (left, right string) {
	split := strings.Split(input, sep)

	return split[0], split[1]
}

func HandleError(err error, message string) {
	if err != nil {
		log.Printf("%s: %v", message, err)
	}
	panic(err)
}

func Sum(nums []int) int {
	var result int
	for _, i := range nums {
		result += i
	}

	return result
}
