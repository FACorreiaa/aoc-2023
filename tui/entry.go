package tui

import (
	"github.com/FACorreiaa/aoc-2023/entry"
	"github.com/FACorreiaa/aoc-2023/tui/constants"
	"github.com/charmbracelet/bubbles/key"
	"github.com/charmbracelet/bubbles/paginator"
	"github.com/charmbracelet/bubbles/viewport"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
	"os"
)

type (
	// UpdatedEntries holds the new entries from DB
	UpdatedEntries []entry.Entry
)

type editorFinishedMsg struct {
	err  error
	file *os.File
}

var cmd tea.Cmd

// Entry implements tea.Model
type Entry struct {
	viewport          viewport.Model
	activeProjectName string
	error             string
	paginator         paginator.Model
	entries           []entry.Entry
	quitting          bool
}

// Init run any intial IO on program start
func (m Entry) Init() tea.Cmd {
	return nil
}

// Update handle IO and commands
func (m Entry) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	var cmds []tea.Cmd
	switch msg := msg.(type) {
	case tea.WindowSizeMsg:
		top, right, bottom, left := constants.DocStyle.GetMargin()
		m.viewport = viewport.New(constants.WindowSize.Width-left-right, constants.WindowSize.Height-top-bottom-6)
	case errMsg:
		m.error = msg.Error()
	//case editorFinishedMsg:
	//	m.quitting = false
	//	if msg.err != nil {
	//		return m, tea.Quit
	//	}
	//	cmds = append(cmds, m.createEntryCmd(msg.file))
	case UpdatedEntries:
		m.entries = msg
		m.paginator.SetTotalPages(len(m.entries))
	case tea.KeyMsg:
		switch {
		//case key.Matches(msg, constants.Keymap.Create):
		//	// TODO: remove m.quitting after bug in Bubble Tea (#431) is fixed
		//	m.quitting = true
		//	return m, openEditorCmd()
		case key.Matches(msg, constants.Keymap.Back):
			return NewModel(), nil
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

func (m Entry) helpView() string {
	// TODO: use the keymaps to populate the help string
	return constants.HelpStyle("\n ↑/↓: navigate  • esc: back • q: quit\n")
}

// InitEntry initialize the entryui model for your program
func InitEntry(activeProject string, p *tea.Program) *Entry {
	m := Entry{activeProjectName: activeProject}
	top, right, bottom, left := constants.DocStyle.GetMargin()
	m.viewport = viewport.New(constants.WindowSize.Width-left-right, constants.WindowSize.Height-top-bottom-1)
	m.viewport.Style = lipgloss.NewStyle().Align(lipgloss.Bottom)

	// init paginator
	m.paginator = paginator.New()
	m.paginator.Type = paginator.Dots
	m.paginator.ActiveDot = lipgloss.NewStyle().Foreground(lipgloss.AdaptiveColor{Light: "235", Dark: "252"}).Render("•")
	m.paginator.InactiveDot = lipgloss.NewStyle().Foreground(lipgloss.AdaptiveColor{Light: "250", Dark: "238"}).Render("•")

	m.paginator.SetTotalPages(len(m.entries))
	// set content
	return &m
}

func (m Entry) View() string {
	if m.quitting {
		return ""
	}

	formatted := lipgloss.JoinVertical(lipgloss.Left, "\n", m.viewport.View(), m.helpView(), constants.ErrStyle(m.error), m.paginator.View())
	return constants.DocStyle.Render(formatted)
}
