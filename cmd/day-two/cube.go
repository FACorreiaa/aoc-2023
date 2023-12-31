package daytwo

import (
	"fmt"
	"github.com/FACorreiaa/aoc-2023/cmd/settings"
	"log"
	"strings"
	"time"
)

type Set map[string]int

type Game struct {
	ID   int
	Sets []Set
}

func (g Game) CalculateMax() Set {
	maxValue := make(Set)

	for _, set := range Sets {
		for colour, num := range set {
			if num > maxValue[colour] {
				maxValue[colour] = num
			}
		}
	}

	return maxValue
}
func (g Game) IsPossible(red, green, blue int) bool {
	maxValue := CalculateMax()

	if maxValue["red"] > red || maxValue["green"] > green || maxValue["blue"] > blue {
		return false
	}
	return true
}

func partOne(games []Game) int {
	var total int

	for _, game := range games {
		if game.IsPossible(12, 13, 14) {
			total += game.ID
		}
	}
	loPrint("Day two part one: ", total)
	return total
}

func parseGames(lines []string) []Game {
	var games []Game
	for _, line := range lines {
		games = append(games, parseGame(line))
	}
	return games
}

func parseGame(line string) Game {
	game, sets := settings.Split(line, ":")
	var id int
	if _, err := fmt.Sscanf(game, "Game %d", &id); err != nil {
		settings.HandleError(err, "Error handling game parsing")
		panic(err)
	}
	return Game{
		ID:   id,
		Sets: parseSets(sets),
	}
}

func parseSets(line string) []Set {
	var sets []Set
	for _, set := range strings.Split(line, ";") {
		sets = append(sets, parseSet(set))
	}
	return sets
}

func parseSet(input string) Set {
	set := make(Set)
	for _, part := range strings.Split(input, ",") {
		var num int
		var colour string
		if _, err := fmt.Sscanf(part, "%d %s", &num, &colour); err != nil {
			settings.HandleError(err, "Error parsing set")
			panic(err)
		}
		set[colour] = num
	}
	return set
}

func partTwo(games []Game) int {
	var total int
	for _, game := range games {
		total += game.Power()
	}

	loPrint("Day two part one: ", total)
	return total
}

func (g Game) Power() int {
	maxValue := CalculateMax()
	return maxValue["red"] * maxValue["blue"] * maxValue["green"]
}

func Start() {

	lines := settings.GetLines("./cmd/day-two/cube.txt")
	for _, line := range lines {
		println(line)
	}
	games := parseGames(lines)
	partOneStart := time.Now()
	partOneResult := partOne(games)
	loPrint("Result: ", partOneResult)
	loPrint("\nDay two part one took: ", time.Since(partOneStart))
	partTwoStart := time.Now()
	partTwoResult := partTwo(games)
	loPrint("Result: ", partTwoResult)
	loPrint("\nDay two part two took: ", time.Since(partTwoStart))
}
