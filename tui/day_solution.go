package tui

import (
	dayone "github.com/FACorreiaa/aoc-2023/cmd/day-one"
	"github.com/FACorreiaa/aoc-2023/common"
	model_solution "github.com/FACorreiaa/aoc-2023/model"
	"github.com/charmbracelet/bubbles/key"
	"github.com/charmbracelet/bubbles/paginator"
	"github.com/charmbracelet/bubbles/viewport"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/glamour"
	"github.com/charmbracelet/lipgloss"
	"os"
)

var dayoneFunctions = map[string]func() tea.Msg{
	"Day 1": dayone.Start,
	// Add other entries as needed
}

type (
	// UpdatedSolution holds the new entries from DB
	UpdatedSolution model_solution.SolutionModelBase
)

//func GetSolution() (model_solution.SolutionModelBase, error) {
//	var s model_solution.SolutionModelBase
//	return &s{
//		"Day 1",
//		"ddd",
//	}, nil
//}

type editorFinishedMsg struct {
	err  error
	file *os.File
}

var cmd tea.Cmd

// Entry implements tea.Model
type Entry struct {
	viewport            viewport.Model
	activeSolutionTitle string
	error               string
	paginator           paginator.Model
	entries             model_solution.SolutionModelBase
	quitting            bool
	Result              tea.Msg        // Result from dayone package
	startFn             func() tea.Msg // Function to start the computation

}

func (e Entry) FilterValue() string { return e.activeSolutionTitle }

// Init run any intial IO on program start
func (e Entry) Init() tea.Cmd {
	return nil
}

// InitSolution initialize the entryui model for your program
func InitSolution(title string, p *tea.Program, startFn func() tea.Msg) *Entry {
	e := Entry{activeSolutionTitle: title, startFn: startFn}
	top, right, bottom, left := common.DocStyle.GetMargin()
	e.viewport = viewport.New(common.WindowSize.Width-left-right, common.WindowSize.Height-top-bottom-1)
	e.viewport.Style = lipgloss.NewStyle().Align(lipgloss.Bottom)

	// init paginator
	e.paginator = paginator.New()
	e.paginator.Type = paginator.Dots
	e.paginator.ActiveDot = lipgloss.NewStyle().Foreground(lipgloss.AdaptiveColor{Light: "235", Dark: "252"}).Render("•")
	e.paginator.InactiveDot = lipgloss.NewStyle().Foreground(lipgloss.AdaptiveColor{Light: "250", Dark: "238"}).Render("•")

	// Fetch the result from the dayone package
	startFn, ok := dayoneFunctions[title]
	if !ok {
		// Handle the case where there is no function for the given title
		return nil
	}
	result := startFn()

	// Set the result and title in the Entry struct
	e.Result = UpdatedSolution{
		Title:  title,
		Result: result,
	}
	e.entries = model_solution.SolutionModelBase(e.Result.(UpdatedSolution))

	// set content
	e.setViewportContent()
	return &e
}

func (e *Entry) setViewportContent() {
	var content string
	//if len(e.entries) == 0 {
	//	content = "There are no entries for this project :)"
	//} else {
	//	content = solution.FormatEntry(e.entries[e.paginator.Page])
	//}
	str, err := glamour.Render(content, "dark")
	if err != nil {
		e.error = "could not render content with glamour"
	}
	e.viewport.SetContent(str)
}

// Update handle IO and commands
func (e Entry) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	var cmds []tea.Cmd
	switch msg := msg.(type) {
	case tea.WindowSizeMsg:
		top, right, bottom, left := common.DocStyle.GetMargin()
		e.viewport = viewport.New(common.WindowSize.Width-left-right, common.WindowSize.Height-top-bottom-6)
	case errMsg:
		e.error = msg.Error()
	//case editorFinishedMsg:
	//	e.quitting = false
	//	if msg.err != nil {
	//		return m, tea.Quit
	//	}
	//	cmds = append(cmds, e.createEntryCmd(msg.file))
	case model_solution.SolutionModelBase:
		e.entries = msg
		//e.paginator.SetTotalPages(len(e.entries))
		e.setViewportContent()
		e.Result = e.startFn()

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
