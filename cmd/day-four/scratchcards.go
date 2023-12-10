package dayfour

import (
	"github.com/FACorreiaa/aoc-2023/cmd/common"
	"log"
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

func calculateScorePartTwo(s string) (int, int) {
	scratchCardID := strings.Split(s, ": ")

	id, err := strconv.Atoi(strings.Split(common.StringRegexMatch.ReplaceAllString(scratchCardID[0], " "), " ")[1])
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

func Start() {
	lines := common.GetLines("./cmd/day-four/scratchcards.txt")
	for _, line := range lines {
		println(line)
	}
	partOneStart := time.Now()
	partOneResult := partOne(strings.Join(lines, "\n"))
	log.Print("Result: ", partOneResult)
	log.Print("\nDay four part one took: ", time.Since(partOneStart))
	partTwoStart := time.Now()
	partTwoResult := partTwo(strings.Join(lines, "\n"))
	log.Print("Result: ", partTwoResult)
	log.Print("\nDay four part two took: ", time.Since(partTwoStart))

}
