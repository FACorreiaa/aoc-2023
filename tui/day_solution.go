package tui

import (
	"fmt"
	dayone "github.com/FACorreiaa/aoc-2023/cmd/day-one"
	"github.com/FACorreiaa/aoc-2023/common"
	"github.com/FACorreiaa/aoc-2023/solution"
	"github.com/charmbracelet/bubbles/key"
	"github.com/charmbracelet/bubbles/paginator"
	"github.com/charmbracelet/bubbles/viewport"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/glamour"
	"github.com/charmbracelet/lipgloss"
)

type (
	// UpdatedSolution holds the new entries from DB
	UpdatedSolution common.Day
)

var cmd tea.Cmd

// Entry implements tea.Model
type Entry struct {
	viewport            viewport.Model
	activeSolutionTitle tea.Msg
	error               string
	list                tea.Cmd
	paginator           paginator.Model
	entry               common.Day
	quitting            bool
	Result              tea.Msg // Result from dayone package
	Title               string
	//list                list.Model

}

func (e Entry) FilterValue() tea.Msg { return e.activeSolutionTitle }

// Init run any intial IO on program start
func (e Entry) Init() tea.Cmd {
	return nil
}

func getResult(title string) int {
	mappedValues := map[string]func() common.Day{
		"Day 1": dayone.Start,
		// Add more entries as needed
	}
	return mappedValues[title]().Result
}

// InitSolution initialize the solution model  program
func InitSolution(title string) *Entry {
	e := Entry{Title: title}
	top, right, bottom, left := common.DocStyle.GetMargin()
	e.viewport = viewport.New(common.WindowSize.Width-left-right, common.WindowSize.Height-top-bottom-1)
	e.viewport.Style = lipgloss.NewStyle().Align(lipgloss.Bottom)

	result := getResult(title)

	// Set the result and title in the Entry struct WIP
	e.Result = UpdatedSolution{
		DayTitle: title,
		Result:   result,
	}
	e.entry = common.Day(e.Result.(UpdatedSolution))
	// init paginator
	e.paginator = paginator.New()
	e.paginator.Type = paginator.Dots
	e.paginator.ActiveDot = lipgloss.NewStyle().Foreground(lipgloss.AdaptiveColor{Light: "235", Dark: "252"}).Render("•")
	e.paginator.InactiveDot = lipgloss.NewStyle().Foreground(lipgloss.AdaptiveColor{Light: "250", Dark: "238"}).Render("•")

	// set content
	e.setViewportContent()
	return &e
}

func (e Entry) setViewportContent() {
	var content string
	content = solution.FormatSolution(e.entry)
	//fmt.Printf("%#v", content)
	str, err := glamour.Render(content, "dark")
	if err != nil {
		e.error = "could not render content with glamour"
	}
	e.viewport.SetContent(str)
}

// Update handle IO and commands
func (e Entry) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	var cmds []tea.Cmd

	switch
	msg := ms(
	type) {

case UpdatedSolution:
	fmt.Printf("%#v", msg)

	e.entry = common.Day(msg)
	e.paginator.SetTotalPages(1)
	e.setViewportContent()
	case tea.WindowSizeMsg:
	fmt.Printf("%#v", msg)

	top, right, bottom, left := common.DocStyle.GetMargin()
	e.viewport = viewport.New(common.WindowSize.Width-left-right, common.WindowSize.Height-top-bottom-6)
	case errMsg:
	e.error = msError()
	//case editorFinishedMsg:
	//	e.quitting = false
	//	if mserr != nil {
	//		return m, tea.Quit
	//	}
	//	cmds = append(cmds, e.createEntryCmd(msfile))

	case tea.KeyMsg:
	switch {
	//case key.Matches(msg, common.Keymap.Create):
	//	// TODO: remove e.quitting after bug in Bubble Tea (#431) is fixed
	//	e.quitting = true
	//	return m, openEditorCmd()
	case key.Matches(msg, common.Keymap.Back):
	return InitProject()
	case key.Matches(msg, common.Keymap.Quit):
	e.quitting = true
	return e, tea.Quit
	}
	}

	e.viewport, cmd = e.viewport.Update(msg)
	e.paginator, cmd = e.paginator.Update(msg)
	cmds = append(cmds, cmd)
	e.setViewportContent() // refresh the content on every Update call
	return e, tea.Batch(cmds...)
	}

func (e Entry) helpView() string {
	// TODO: use the keymaps to populate the help string
	return common.HelpStyle("\n ↑/↓: navigate  • esc: back • c: create entry • d: delete entry • q: quit\n")
}

// View return the text UI to be output to the terminal
func (e Entry) View() string {
	if e.quitting {
		return ""
	}

	formatted := lipgloss.JoinVertical(lipgloss.Left, "\n", e.viewport.View(), e.helpView(), common.ErrStyle(e.error), e.paginator.View())
	return common.DocStyle.Render(formatted)
}
