package dayfour

import (
	"fmt"
	"github.com/FACorreiaa/aoc-2023/cmd/common"
	"regexp"
	"strconv"
	"strings"
	"time"
)

func partOne(s string) int {
	score := 0
	for _, line := range strings.Split(s, "\n") {
		score += calculateScore(line)
	}

	return score
}

func partTwo(s string) int {
	score := 0
	cardScores := map[int]int{1: 1}
	for _, line := range strings.Split(s, "\n") {
		id, score := calculateScorePartTwo(line)

		if _, ok := cardScores[id]; !ok {
			cardScores[id] = 1
		}

		if score == 0 {
			continue
		}

		for i := id + 1; i <= id+score; i++ {
			if _, ok := cardScores[i]; !ok {
				cardScores[i] = 1
			}
			cardScores[i] += cardScores[id]
		}
	}

	for _, v := range cardScores {
		score += v
	}

	return score
}

func calculateScore(s string) int {
	scratchCards := strings.Split(s, ": ")[1]

	wp := strings.Split(scratchCards, " | ")

	wins := map[string]bool{}
	plays := map[string]bool{}

	//wins
	for _, win := range strings.Split(wp[0], " ") {
		win = strings.Trim(win, " ")

		if win == "" {
			continue
		}

		wins[win] = true
	}

	//plays
	for _, play := range strings.Split(wp[1], " ") {
		play = strings.Trim(play, " ")

		if play == "" {
			continue
		}

		plays[play] = true
	}

	score := 0
	for key, _ := range plays {
		if ok := wins[key]; ok {
			if score == 0 {
				score = 1
				continue
			}
			score = score * 2
		}
	}
	return score
}

var stringRegexMatch = regexp.MustCompile(`\s+`)

func calculateScorePartTwo(s string) (int, int) {
	scratchCardID := strings.Split(s, ": ")

	id, err := strconv.Atoi(strings.Split(stringRegexMatch.ReplaceAllString(scratchCardID[0], " "), " ")[1])
	if err != nil {
		common.HandleError(err, "error converting string")
	}
	scratchCards := scratchCardID[1]
	wp := strings.Split(scratchCards, " | ")

	wins := map[string]bool{}
	plays := map[string]bool{}

	//wins
	for _, win := range strings.Split(wp[0], " ") {
		win = strings.Trim(win, " ")

		if win == "" {
			continue
		}

		wins[win] = true
	}

	//plays
	for _, play := range strings.Split(wp[1], " ") {
		play = strings.Trim(play, " ")

		if play == "" {
			continue
		}

		plays[play] = true
	}

	score := 0
	for key, _ := range plays {
		if ok := wins[key]; ok {
			score += 1
		}
	}
	return id, score
}

func extractPartOneNumbers(lines []string) []int {
	var numbers []int
	for _, _ = range lines {
		numbers = append(numbers, partOne(strings.Join(lines, "\n")))
	}

	return numbers
}

func extractPartTwoNumbers(lines []string) []int {
	var numbers []int
	for _, _ = range lines {
		numbers = append(numbers, partTwo(strings.Join(lines, "\n")))
	}

	return numbers
}

func StartDayFour() {
	lines := common.GetLines("./cmd/day-four/scratchcards.txt")
	for _, line := range lines {
		fmt.Println(line)
	}
	partOneStart := time.Now()
	extractPartOneNumbers(lines)
	fmt.Println("\nDay four part one took: ", time.Since(partOneStart))
	fmt.Println("Result: ", extractPartOneNumbers(lines))

	partTwoStart := time.Now()
	fmt.Println("\nDay four part two took: ", time.Since(partTwoStart))
	extractPartTwoNumbers(lines)
	fmt.Println("Result: ", extractPartTwoNumbers(lines))

}
