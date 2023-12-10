package lib

import "sync"

type Item struct {
	title       string
	description string
	subMenu     bool
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

func (r *RandomItemGenerator) Next() Item {
	if r.mtx == nil {
		r.Reset()
	}

	r.mtx.Lock()
	defer r.mtx.Unlock()

	i := Item{
		title:       r.titles[r.titleIndex],
		description: r.description[r.descIndex],
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
