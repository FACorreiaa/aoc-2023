package lib

import (
	dayeight "github.com/FACorreiaa/aoc-2023/cmd/day-eight"
	dayfive "github.com/FACorreiaa/aoc-2023/cmd/day-five"
	dayfour "github.com/FACorreiaa/aoc-2023/cmd/day-four"
	daynine "github.com/FACorreiaa/aoc-2023/cmd/day-nine"
	dayone "github.com/FACorreiaa/aoc-2023/cmd/day-one"
	dayseven "github.com/FACorreiaa/aoc-2023/cmd/day-seven"
	daysix "github.com/FACorreiaa/aoc-2023/cmd/day-six"
	daythree "github.com/FACorreiaa/aoc-2023/cmd/day-three"
	daytwo "github.com/FACorreiaa/aoc-2023/cmd/day-two"
	tea "github.com/charmbracelet/bubbletea"
	"sync"
)

type Item struct {
	title       string
	description string
	subMenu     bool
	onPress     func(title string) tea.Msg
}

func (i Item) Title() string       { return i.title }
func (i Item) Description() string { return i.description }
func (i Item) FilterValue() string { return i.title }

type RandomItemGenerator struct {
	titles      []string
	description []string
	titleIndex  int
	descIndex   int
	mtx         *sync.Mutex
	onPress     func(title string) tea.Msg
	//shuffle     *sync.Once
}

func (r *RandomItemGenerator) Reset() {
	r.mtx = &sync.Mutex{}
	//r.shuffle = &sync.Once{}

	r.titles = []string{
		"Day 1",
		"Day 2",
		"Day 3",
		"Day 4",
		"Day 5",
		"Day 6",
		"Day 7",
		"Day 8",
		"Day 9",
	}

	r.description = []string{
		"Trebuchet?!",
		"Cube Conundrum",
		"Gear Ratios",
		"Scratchcards",
		"If You Give A Seed A Fertilizer",
		"Wait for it",
		"Camel Cards",
		"Haunted Wasteland",
		"Mirage Maintenance",
	}

	//r.shuffle.Do(func() {
	//	shuffle := func(x []string) {
	//		rand.Shuffle(len(x), func(i, j int) { x[i], x[j] = x[j], x[i] })
	//	}
	//	shuffle(r.titles)
	//	shuffle(r.description)
	//})
}

func (r *RandomItemGenerator) Choice(title string) func() {
	mapFunction := map[string]func(){
		"Day 1": dayone.Start,
		"Day 2": daytwo.Start,
		"Day 3": daythree.Start,
		"Day 4": dayfour.Start,
		"Day 5": dayfive.Start,
		"Day 6": daysix.Start,
		"Day 7": dayseven.Start,
		"Day 8": dayeight.Start,
		"Day 9": daynine.Start,
	}
	return mapFunction[title]
}
func (r *RandomItemGenerator) Next() Item {
	if r.mtx == nil {
		r.Reset()
	}

	r.mtx.Lock()
	defer r.mtx.Unlock()

	i := Item{
		title:       r.titles[r.titleIndex],
		description: r.description[r.descIndex],
		subMenu:     false,
	}

	r.titleIndex++
	if r.titleIndex >= len(r.titles) {
		r.titleIndex = 0
	}

	r.descIndex++
	if r.descIndex >= len(r.description) {
		r.descIndex = 0
	}

	return i
}
