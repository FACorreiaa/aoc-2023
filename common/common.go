package common

import (
	"github.com/charmbracelet/bubbles/key"
	"github.com/charmbracelet/bubbles/list"
	tea "github.com/charmbracelet/bubbletea"
	"sync"
)

type Day struct {
	DayTitle       string
	DayDescription string
	Result         int
	SubMenu        bool
}

func (d Day) Title() string       { return d.DayTitle }
func (d Day) Description() string { return d.DayDescription }
func (d Day) FilterValue() string { return d.DayTitle }

type RandomItemGenerator struct {
	titles      []string
	description []string
	titleIndex  int
	descIndex   int
	mtx         *sync.Mutex
}

func (r *RandomItemGenerator) Reset() {
	r.mtx = &sync.Mutex{}

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
}

func (r *RandomItemGenerator) Next() Day {
	if r.mtx == nil {
		r.Reset()
	}

	r.mtx.Lock()
	defer r.mtx.Unlock()

	i := Day{
		DayTitle:       r.titles[r.titleIndex],
		DayDescription: r.description[r.descIndex],
		SubMenu:        false,
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

func NewItemDelegate(keys *DelegateKeyMap) list.DefaultDelegate {
	d := list.NewDefaultDelegate()

	d.UpdateFunc = func(msg tea.Msg, m *list.Model) tea.Cmd {
		var title string

		if i, ok := m.SelectedItem().(Day); ok {
			title = i.Title()
		} else {
			return nil
		}

		m.NewStatusMessage(StatusMessageStyle("Check out " + title))

		switch
		msg := ms(
		type) {
	case tea.KeyMsg:
		switch {

		case key.Matches(msg, keys.Choose):
		m.NewStatusMessage(StatusMessageStyle("You chose " + title))

		case key.Matches(msg, keys.Back):
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
