package tui

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
	"github.com/FACorreiaa/aoc-2023/lib"
	"github.com/FACorreiaa/aoc-2023/lib/constants"
	"github.com/charmbracelet/bubbles/key"
	"github.com/charmbracelet/bubbles/list"
	"github.com/charmbracelet/bubbles/paginator"
	"github.com/charmbracelet/bubbles/viewport"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
)

var cmd tea.Cmd

func (m model) UpdateSubmenu(msg tea.Msg) (tea.Model, tea.Cmd) {
	var cmds []tea.Cmd
	switch msg := msg.(type) {
	case tea.WindowSizeMsg:
		top, right, bottom, left := constants.DocStyle.GetMargin()
		m.viewport = viewport.New(constants.WindowSize.Width-left-right, constants.WindowSize.Height-top-bottom-6)
	case errMsg:
		m.error = msg.Error()
	case tea.KeyMsg:
		switch {
		//case key.Matches(msg, constants.Keymap.Create):
		//	// TODO: remove m.quitting after bug in Bubble Tea (#431) is fixed
		//	m.quitting = true
		//	return m, openEditorCmd()
		case key.Matches(msg, constants.Keymap.Back):
			return InitProject()
		case key.Matches(msg, constants.Keymap.Quit):
			m.quitting = true
			return m, tea.Quit
		}
	}

	m.viewport, cmd = m.viewport.Update(msg)
	m.paginator, cmd = m.paginator.Update(msg)
	cmds = append(cmds, cmd)
	return m, tea.Batch(cmds...)
}
func NewItemDelegate(keys *lib.DelegateKeyMap) list.DefaultDelegate {
	d := list.NewDefaultDelegate()

	d.UpdateFunc = func(msg tea.Msg, m *list.Model) tea.Cmd {
		var title string

		if i, ok := m.SelectedItem().(lib.Item); ok {
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
			case key.Matches(msg, keys.Back):
				//init, _ := InitProject()
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

func InitSubmenu(selectedProjectName string, p *tea.Program) *model {
	m := model{selectedProjectName: selectedProjectName}
	//mapFunction := map[string]func(){
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
	top, right, bottom, left := constants.DocStyle.GetMargin()
	m.viewport = viewport.New(constants.WindowSize.Width-left-right, constants.WindowSize.Height-top-bottom-1)
	m.viewport.Style = lipgloss.NewStyle().Align(lipgloss.Bottom)

	// init paginator
	m.paginator = paginator.New()
	m.paginator.Type = paginator.Dots
	m.paginator.ActiveDot = lipgloss.NewStyle().Foreground(lipgloss.AdaptiveColor{Light: "235", Dark: "252"}).Render("•")
	m.paginator.InactiveDot = lipgloss.NewStyle().Foreground(lipgloss.AdaptiveColor{Light: "250", Dark: "238"}).Render("•")

	//m.entries = m.setupEntries().(UpdatedEntries)
	//m.paginator.SetTotalPages(len(m.entries))
	// set content
	//m.setViewportContent()
	return &m
}
