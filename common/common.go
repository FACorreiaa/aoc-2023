package common

import (
	"bufio"
	"log"
	"os"
	"strings"
)

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

func Split(input, sep string) (left, right string) {
	split := strings.SplitN(input, sep, 2)
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