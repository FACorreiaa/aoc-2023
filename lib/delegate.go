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
	"github.com/FACorreiaa/aoc-2023/lib/constants"
	"github.com/charmbracelet/bubbles/key"
	"github.com/charmbracelet/bubbles/list"
	tea "github.com/charmbracelet/bubbletea"
)

func NewItemDelegate(keys *DelegateKeyMap) list.DefaultDelegate {
	d := list.NewDefaultDelegate()

	d.UpdateFunc = func(msg tea.Msg, m *list.Model) tea.Cmd {
		var title string

		if i, ok := m.SelectedItem().(Item); ok {
			title = i.Title()
		} else {
			return nil
		}

		//
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
		switch msg := msg.(type) {
		case tea.KeyMsg:
			switch {
			case key.Matches(msg, keys.Choose):
				m.NewStatusMessage(constants.StatusMessageStyle("You chose " + title))
				return func() tea.Msg {
					mapFunction[title]() // Call the selected function
					return nil
				}
			case key.Matches(msg, keys.Remove):
				index := m.Index()
				m.RemoveItem(index)
				if len(m.Items()) == 0 {
					keys.Remove.SetEnabled(false)
				}
				return m.NewStatusMessage(constants.StatusMessageStyle("Deleted " + title))
			}
		}

		return nil
	}

	help := []key.Binding{keys.Choose, keys.Remove}

	d.ShortHelpFunc = func() []key.Binding {
		return help
	}

	d.FullHelpFunc = func() [][]key.Binding {
		return [][]key.Binding{help}
	}

	return d
}

type DelegateKeyMap struct {
	Choose key.Binding
	Remove key.Binding
}

func (d DelegateKeyMap) ShortHelp() []key.Binding {
	return []key.Binding{
		d.Choose,
		d.Remove,
	}
}

func (d DelegateKeyMap) FullHelp() [][]key.Binding {
	return [][]key.Binding{
		{
			d.Choose,
			d.Remove,
		},
	}
}

func NewDelegateKeyMap() *DelegateKeyMap {
	return &DelegateKeyMap{
		Choose: key.NewBinding(
			key.WithKeys("enter"),
			key.WithHelp("enter", "Choose"),
		),
		Remove: key.NewBinding(
			key.WithKeys("x", "backspace"),
			key.WithHelp("x", "delete"),
		),
	}
}
