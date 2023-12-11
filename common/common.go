package common

import (
	"github.com/FACorreiaa/aoc-2023/messages"
	"github.com/charmbracelet/bubbles/key"
	"github.com/charmbracelet/bubbles/list"
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

//	func (r *RandomItemGenerator) Choice(title string) func() {
//		mapFunction := map[string]func() tea.Msg{
//			"Day 1": dayone.Start,
//			//"Day 2": daytwo.Start,
//			//"Day 3": daythree.Start,
//			//"Day 4": dayfour.Start,
//			//"Day 5": dayfive.Start,
//			//"Day 6": daysix.Start,
//			//"Day 7": dayseven.Start,
//			//"Day 8": dayeight.Start,
//			//"Day 9": daynine.Start,
//		}
//		return mapFunction[title]
//	}
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

func NewItemDelegate(keys *DelegateKeyMap, days map[string]func() tea.Msg) list.DefaultDelegate {
	d := list.NewDefaultDelegate()

	d.UpdateFunc = func(msg tea.Msg, m *list.Model) tea.Cmd {
		var title string

		if i, ok := m.SelectedItem().(Item); ok {
			title = i.Title()
		} else {
			return nil
		}

		//
		//mapFunction := map[string]func() tea.Msg{
		//	"Day 1": dayone.Start,
		//	"Day 2": daytwo.Start,
		//	"Day 3": daythree.Start,
		//	"Day 4": dayfour.Start,
		//	"Day 5": dayfive.Start,
		//	"Day 6": daysix.Start,
		//	"Day 7": dayseven.Start,
		//	"Day 8": dayeight.Start,
		//	"Day 9": daynine.Start,
		//}
		switch msg := msg.(type) {
		case tea.KeyMsg:
			switch {
			case key.Matches(msg, keys.Choose):
				m.NewStatusMessage(StatusMessageStyle("You chose " + title))
				return func() tea.Msg {
					return messages.SolutionTransitionMsg{title, days[title]} // Signal the transition to the selected solution
				}
				//index := m.Index()
				//m.RemoveItem(index)
				//if len(m.Items()) == 0 {
				//    keys.Back.SetEnabled(false)
				//}
				//return m.NewStatusMessage(constants.StatusMessageStyle("Deleted " + title))
			//case key.Matches(msg, keys.Back):
			case key.Matches(msg, keys.Quit):
				//m.quitting = true
				return tea.Quit
			}

		}

		return nil
	}

	help := []key.Binding{keys.Choose, keys.Back}

	d.ShortHelpFunc = func() []key.Binding {
		return help
	}

	d.FullHelpFunc = func() [][]key.Binding {
		return [][]key.Binding{help}
	}

	return d
}
