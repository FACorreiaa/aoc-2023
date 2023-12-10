package daysix

import (
	"github.com/FACorreiaa/aoc-2023/cmd/common"
	"log"
	"strconv"
	"strings"
	"time"
)

func partOne(s string) int {
	var score = 1

	common.StringRegexMatch.ReplaceAllString(s, ` `)
	lines := strings.Split(s, "\n")
	times := strings.Split(strings.Split(common.StringRegexMatch.ReplaceAllString(lines[0], ` `), ": ")[1], " ")
	distances := strings.Split(strings.Split(common.StringRegexMatch.ReplaceAllString(lines[1], ` `), ": ")[1], " ")

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

	common.StringRegexMatch.ReplaceAllString(s, ` `)
	lines := strings.Split(s, "\n")
	times := strings.Split(strings.Split(common.StringRegexMatch.ReplaceAllString(lines[0], ` `), ": ")[1], " ")
	distances := strings.Split(strings.Split(common.StringRegexMatch.ReplaceAllString(lines[1], ` `), ": ")[1], " ")

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
	lines := common.GetLines("./cmd/day-six/document.txt")
	for _, line := range lines {
		println(line)
	}
	partOneStart := time.Now()
	partOneResult := partOne(strings.Join(lines, "\n"))
	log.Print("Result: ", partOneResult)
	log.Print("Day six part one took: ", time.Since(partOneStart))
	partTwoStart := time.Now()
	partTwoResult := partTwo(strings.Join(lines, "\n"))
	log.Print("Result: ", partTwoResult)
	log.Print("Day six part two took: ", time.Since(partTwoStart))

}
