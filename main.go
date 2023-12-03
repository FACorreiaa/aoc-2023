package main

import (
	"bufio"
	"fmt"
	"github.com/FACorreiaa/aoc-2023/day-one"
	"os"
)

func main() {
	sum := 0
	file, err := os.Open("./day-one/calibration.txt")
	if err != nil {
		fmt.Println("Error opening file:", err)
		return
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		line := scanner.Text()
		//for part two
		calibrationValue := day_one.ExtractCalibrationValue(line)
		sum += calibrationValue
	}

	fmt.Println("Sum of calibration values:", sum)

	if err := scanner.Err(); err != nil {
		fmt.Println("Error reading file:", err)
	}
}
