package daysix

import (
	"github.com/FACorreiaa/aoc-2023/cmd/settings"
	"log"
	"strconv"
	"strings"
	"time"
)

func partOne(s string) int {
	var score = 1

	settings.StringRegexMatch.ReplaceAllString(s, ` `)
	lines := strings.Split(s, "\n")
	times := strings.Split(strings.Split(settings.StringRegexMatch.ReplaceAllString(lines[0], ` `), ": ")[1], " ")
	distances := strings.Split(strings.Split(settings.StringRegexMatch.ReplaceAllString(lines[1], ` `), ": ")[1], " ")

	for i, ts := range times {
		t, _ := strconv.Atoi(ts)
		d, _ := strconv.Atoi(distances[i])
		count := 0
		for hold := 1; hold < t; hold++ {
			if hold*(t-hold) > d {
				count++
			}
		}
		score = score * count
	}

	return score
}

func partTwo(s string) int {
	var score = 1

	settings.StringRegexMatch.ReplaceAllString(s, ` `)
	lines := strings.Split(s, "\n")
	times := strings.Split(strings.Split(settings.StringRegexMatch.ReplaceAllString(lines[0], ` `), ": ")[1], " ")
	distances := strings.Split(strings.Split(settings.StringRegexMatch.ReplaceAllString(lines[1], ` `), ": ")[1], " ")

	ts := strings.Join(times, "")
	ds := strings.Join(distances, "")
	t, _ := strconv.Atoi(ts)
	d, _ := strconv.Atoi(ds)
	count := 0
	for hold := 1; hold < t; hold++ {
		if hold*(t-hold) > d {
			count++
		}
		score = count
	}

	return score
}

func Start() {
	lines := settings.GetLines("./cmd/day-six/document.txt")
	for _, line := range lines {
		println(line)
	}
	partOneStart := time.Now()
	partOneResult := partOne(strings.Join(lines, "\n"))
	loPrint("Result: ", partOneResult)
	loPrint("Day six part one took: ", time.Since(partOneStart))
	partTwoStart := time.Now()
	partTwoResult := partTwo(strings.Join(lines, "\n"))
	loPrint("Result: ", partTwoResult)
	loPrint("Day six part two took: ", time.Since(partTwoStart))

}
