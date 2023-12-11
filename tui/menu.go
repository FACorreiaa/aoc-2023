package tui

import (
	"fmt"
	common2 "github.com/FACorreiaa/aoc-2023/cmd/common"
	dayone "github.com/FACorreiaa/aoc-2023/cmd/day-one"
	"github.com/FACorreiaa/aoc-2023/common"
	model_solution "github.com/FACorreiaa/aoc-2023/model"
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

func (m model) Init() tea.Cmd {
	return tea.EnterAltScreen
	//return nil
}

// testing only for Day 1
func createSolutionModel(item list.Item) (model_solution.SolutionModel, error) {
	title := item.FilterValue()
	mappedValues := map[string]func() tea.Msg{
		"Day 1": dayone.Start,
	}
	result, ok := mappedValues[title]
	if !ok {
		return nil, fmt.Errorf("unsupported day title: %s", title)
	}

	solutionModel, _ := model_solution.DayOneStart(title, result)
	return solutionModel, nil
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
		case key.Matches(msg, m.keys.chooseItem):
			if m.subMenu {
				selectedItem := m.list.SelectedItem().(list.Item)
				solutionModel, err := createSolutionModel(selectedItem)
				if err != nil {
					common2.HandleError(err, "Error getting the model")
					return m, nil
				}
				fmt.Printf("%#v", solutionModel)

				initCmd := solutionModel.Init()
				return solutionModel, initCmd
			} else {
				// If in the main menu, switch to the sub-menu
				m.subMenu = true
			}
			return m, nil
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

	newListModel, cmd := m.list.Update(msg)
	m.list = newListModel
	commands = append(commands, cmd)

	return m, tea.Batch(commands...)
}

func (m model) View() string {
	if m.subMenu {
		return m.list.View()
	}

	return common.AppStyle.Render(m.list.View())
}

func Start() error {
	rand.NewSource(time.Now().UTC().UnixNano())

	m, _ := InitProject()

	common.P = tea.NewProgram(m, tea.WithAltScreen())

	if _, err := common.P.Run(); err != nil {
		fmt.Println("Error running program:", err)
		os.Exit(1)
	}

	return nil
}
