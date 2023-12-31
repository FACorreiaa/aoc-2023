package dayfive

import (
	"github.com/FACorreiaa/aoc-2023/cmd/settings"
	"log"
	"strconv"
	"strings"
	"time"
)

func partOne(s string) int64 {
	var score int64 = -1
	seeds, maps := buildMaps(s)

	for _, seed := range seeds {
		temp := calculateScore(seed, maps)
		if score == -1 {
			score = temp
			continue
		}

		if temp < score {
			score = temp
		}
	}
	return score
}

func partTwo(s string) int64 {
	var score int64 = -1
	seeds, maps := buildMaps(s)
	for i := 0; i < len(seeds); i += 2 {
		for j := seeds[i]; j < (seeds[i] + seeds[i+1]); j++ {
			temp := calculateScore(j, maps)
			if score == -1 {
				score = temp
				continue
			}

			if temp < score {
				score = temp
			}
		}
	}
	return score
}

func buildMaps(s string) ([]int64, map[string][][]int64) {
	seeds := make([]int64, 0)
	maps := map[string][][]int64{}

	for _, line := range strings.Split(s, "\n\n") {
		if strings.Contains(line, "seeds:") {
			seeds = getInts(strings.ReplaceAll(line, "seeds: ", ""))
			continue
		}

		mappings := strings.Split(line, "\n")
		name := strings.Split(mappings[0], " ")[0]

		for i := 1; i < len(mappings); i++ {
			maps[name] = append(maps[name], getInts(mappings[i]))
		}
	}

	return seeds, maps
}

func getInts(s string) []int64 {
	ints := make([]int64, 0)
	for _, v := range strings.Split(s, " ") {
		i, err := strconv.Atoi(v)
		if err != nil {
			settings.HandleError(err, "Error converting string")
		}
		ints = append(ints, int64(i))
	}
	return ints
}

func calculateScore(seed int64, maps map[string][][]int64) int64 {
	soil := getDest(seed, maps[`seed-to-soil`])
	fertilizer := getDest(soil, maps[`soil-to-fertilizer`])
	water := getDest(fertilizer, maps[`fertilizer-to-water`])
	light := getDest(water, maps[`water-to-light`])
	temperature := getDest(light, maps[`light-to-temperature`])
	humidity := getDest(temperature, maps[`temperature-to-humidity`])
	location := getDest(humidity, maps[`humidity-to-location`])
	return location
}

func getDest(source int64, mapping [][]int64) int64 {
	for _, m := range mapping {
		if m[1] <= source && source <= m[1]+m[2] {
			return m[0] + (source - m[1])
		}
	}
	return source
}

func Start() {
	lines := settings.GetLines("./cmd/day-five/seeds.txt")
	for _, line := range lines {
		println(line)
	}
	partOneStart := time.Now()
	partOneResult := partOne(strings.Join(lines, "\n"))
	loPrint("Result: ", partOneResult)
	loPrint("Day five part one took: ", time.Since(partOneStart))
	partTwoStart := time.Now()
	partTwoResult := partTwo(strings.Join(lines, "\n"))
	loPrint("Result: ", partTwoResult)
	loPrint("Day five part two took: ", time.Since(partTwoStart))

}
