package daythree

import (
	"fmt"
	"github.com/FACorreiaa/aoc-2023/cmd/common"
	"strconv"
	"strings"
	"time"
)

func partOne(s string) int {
	score := 0
	lines := strings.Split(s, "\n")

	numStart := -1
	numEnd := -1
	for x, line := range lines {
		for y, r := range line {
			if r >= '0' && r <= '9' {
				if numStart == -1 {
					numStart = y
					numEnd = y
				}
				numEnd = y
			}

			if (r == '.' || isSymbol(byte(r))) && numStart != -1 {
				//number check
				i, err := strconv.Atoi(lines[x][numStart : numEnd+1])
				if err != nil {
					common.HandleError(err, "Error finding symbols around number")
				}

				if found, _, _ := hasNeighbour(lines, x, numStart, numEnd, isSymbol); found {
					score += i
				}

				numStart = -1
				numEnd = -1
			}
		}
		if numStart != -1 {
			//number check
			i, err := strconv.Atoi(lines[x][numStart : numEnd+1])
			if err != nil {
				common.HandleError(err, "Error finding symbols around number")
			}

			if found, _, _ := hasNeighbour(lines, x, numStart, numEnd, isSymbol); found {
				score += i
			}

			numStart = -1
			numEnd = -1
		}
	}
	return score
}

func hasNeighbour(input []string, x, ys, ye int, check func(byte) bool) (bool, int, int) {
	for i := ys; i <= ye; i++ {
		//checks

		if x > 0 {
			if check(input[x-1][i]) {
				return true, x - 1, i
			}
		}

		if x < len(input)-1 {
			if check(input[x+1][i]) {
				return true, x + 1, i
			}
		}
	}

	if ys > 0 {
		if check(input[x][ys-1]) {
			return true, x, ys - 1
		}

		if x > 0 {
			if check(input[x-1][ys-1]) {
				return true, x - 1, ys - 1
			}
		}

		if x < len(input)-1 {
			if check(input[x+1][ys-1]) {
				return true, x + 1, ys - 1
			}
		}
	}

	if ye < len(input[x])-1 {
		if check(input[x][ye+1]) {
			return true, x, ye + 1
		}

		if x > 0 {
			if check(input[x-1][ye+1]) {
				return true, x - 1, ye + 1
			}
		}

		if x < len(input)-1 {
			if check(input[x+1][ye+1]) {
				return true, x + 1, ye + 1
			}
		}
	}

	return false, -1, -1
}

func isSymbol(b byte) bool {
	if b >= '0' && b <= '9' || b == '.' {
		return false
	}
	return true
}

func extractPartOneNumbers(lines []string) []int {
	var numbers []int
	for _, _ = range lines {
		numbers = append(numbers, partOne(strings.Join(lines, "\n")))
	}

	return numbers
}

func isAsterisk(b byte) bool {
	return b == '*'
}

func partTwo(s string) int {
	score := 0
	validScore := map[string][]int{}
	lines := strings.Split(s, "\n")

	numStart := -1
	numEnd := -1
	for x, line := range lines {
		for y, r := range line {
			if r >= '0' && r <= '9' {
				if numStart == -1 {
					numStart = y
					numEnd = y
				}
				numEnd = y
			}

			if (r == '.' || isSymbol(byte(r))) && numStart != -1 {
				//number check
				i, err := strconv.Atoi(lines[x][numStart : numEnd+1])
				if err != nil {
					common.HandleError(err, "Error finding symbols around number")
				}

				if found, fx, fy := hasNeighbour(lines, x, numStart, numEnd, isAsterisk); found {
					validScore[fmt.Sprintf("%d,%d", fx, fy)] = append(validScore[fmt.Sprintf("%d,%d", fx, fy)], i)
				}

				numStart = -1
				numEnd = -1
			}
		}
		if numStart != -1 {
			//number check
			i, err := strconv.Atoi(lines[x][numStart : numEnd+1])
			if err != nil {
				common.HandleError(err, "Error finding symbols around number")
			}

			if found, fx, fy := hasNeighbour(lines, x, numStart, numEnd, isAsterisk); found {
				validScore[fmt.Sprintf("%d,%d", fx, fy)] = append(validScore[fmt.Sprintf("%d,%d", fx, fy)], i)
			}

			numStart = -1
			numEnd = -1
		}
	}
	for _, value := range validScore {
		if len(value) != 2 {
			continue
		}
		score += value[0] * value[1]
	}
	return score
}

func extractPartTwoNumbers(lines []string) []int {
	var numbers []int
	for _, _ = range lines {
		numbers = append(numbers, partTwo(strings.Join(lines, "\n")))
	}

	return numbers
}

func StartDayThree() {
	lines := common.GetLines("./cmd/day-three/gear.txt")
	for _, line := range lines {
		fmt.Println(line)
	}
	extractPartOneNumbers(lines)
	partOneStart := time.Now()
	fmt.Println("\nDay three part one took: ", time.Since(partOneStart))
	fmt.Println("Result: ", extractPartOneNumbers(lines))

	partTwoStart := time.Now()
	extractPartTwoNumbers(lines)
	fmt.Println("\nDay three part two took: ", time.Since(partTwoStart))
	fmt.Println("Result: ", extractPartTwoNumbers(lines))

}
