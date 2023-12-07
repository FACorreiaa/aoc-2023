package daysix

import (
	"fmt"
	"github.com/FACorreiaa/aoc-2023/cmd/common"
	"log"
	"regexp"
	"strconv"
	"strings"
	"time"
)

var spaceRegExp = regexp.MustCompile(`\s+`)

func partOne(s string) int {
	var score = 1

	spaceRegExp.ReplaceAllString(s, ` `)
	lines := strings.Split(s, "\n")
	times := strings.Split(strings.Split(spaceRegExp.ReplaceAllString(lines[0], ` `), ": ")[1], " ")
	distances := strings.Split(strings.Split(spaceRegExp.ReplaceAllString(lines[1], ` `), ": ")[1], " ")

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
	fmt.Printf("%#v", distances)
	fmt.Printf("%#v\n", times)

	return score
}

func partTwo(s string) int {
	var score = 1

	spaceRegExp.ReplaceAllString(s, ` `)
	lines := strings.Split(s, "\n")
	times := strings.Split(strings.Split(spaceRegExp.ReplaceAllString(lines[0], ` `), ": ")[1], " ")
	distances := strings.Split(strings.Split(spaceRegExp.ReplaceAllString(lines[1], ` `), ": ")[1], " ")

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
	fmt.Printf("%#v", distances)
	fmt.Printf("%#v\n", times)

	return score
}

func calculateScore() int {
	return 0
}

func Start() {
	lines := common.GetLines("./cmd/day-six/document.txt")
	for _, line := range lines {
		log.Print(line)
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
