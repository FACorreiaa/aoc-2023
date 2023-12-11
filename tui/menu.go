package tui

import (
	"fmt"
	"github.com/FACorreiaa/aoc-2023/common"
	"github.com/charmbracelet/bubbles/key"
	"github.com/charmbracelet/bubbles/list"
	tea "github.com/charmbracelet/bubbletea"
	"math/rand"
	"os"
	"time"
)

type Model interface {
	Init() tea.Cmd
	Update(msg tea.Msg) (Model, tea.Cmd)
	View() string
}

//func (m menu) Init() tea.Cmd { return nil }

func (m model) Init() tea.Cmd {
	return tea.EnterAltScreen
}

func (m model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	var commands []tea.Cmd

	switch msg := msg.(type) {

	case tea.WindowSizeMsg:
		h, v := common.AppStyle.GetFrameSize()
		m.list.SetSize(msg.Width-h, msg.Height-v)

	case tea.KeyMsg:
		// Don't match any of the keys below if we're actively filtering.
		if m.list.FilterState() == list.Filtering {
			break
		}

		switch {
		//case key.Matches(msg, constants.Keymap.Enter):
		//	return m,
		case key.Matches(msg, common.Keymap.Quit):
			m.quitting = true
			return m, tea.Quit

		case key.Matches(msg, m.keys.toggleSpinner):
			cmd := m.list.ToggleSpinner()
			return m, cmd

		case key.Matches(msg, m.keys.toggleTitleBar):
			v := !m.list.ShowTitle()
			m.list.SetShowTitle(v)
			m.list.SetShowFilter(v)
			m.list.SetFilteringEnabled(v)
			return m, nil

		case key.Matches(msg, m.keys.toggleStatusBar):
			m.list.SetShowStatusBar(!m.list.ShowStatusBar())
			return m, nil

		case key.Matches(msg, m.keys.togglePagination):
			m.list.SetShowPagination(!m.list.ShowPagination())
			return m, nil

		case key.Matches(msg, m.keys.toggleHelpMenu):
			m.list.SetShowHelp(!m.list.ShowHelp())
			return m, nil

		case key.Matches(msg, m.keys.insertItem):
			m.delegateKeys.Back.SetEnabled(true)
			newItem := m.itemGenerator.Next()
			insCmd := m.list.InsertItem(0, newItem)
			statusCmd := m.list.NewStatusMessage(common.StatusMessageStyle("Added " + newItem.Title()))
			return m, tea.Batch(insCmd, statusCmd)
		}
	}

	// This will also call our delegate's update function.
	newListModel, cmd := m.list.Update(msg)
	m.list = newListModel
	commands = append(commands, cmd)

	return m, tea.Batch(commands...)
}

func (m model) View() string {
	if m.subMenu {
		// Render sub-menu UI
		return m.list.View()
	}

	return common.AppStyle.Render(m.list.View())
}

func Start() error {
	rand.NewSource(time.Now().UTC().UnixNano())
	//f, err := tea.LogToFile("debug.log", "debug")
	//if err != nil {
	//	fmt.Println(err)
	//	os.Exit(1)
	//}
	//defer f.Close()
	//m, _ := InitProject()
	m, _ := InitProject()

	common.P = tea.NewProgram(m, tea.WithAltScreen())

	if _, err := common.P.Run(); err != nil {
		fmt.Println("Error running program:", err)
		os.Exit(1)
	}

	return nil
}
