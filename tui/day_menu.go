package tui

import (
	"fmt"
	dayone "github.com/FACorreiaa/aoc-2023/cmd/day-one"
	"github.com/FACorreiaa/aoc-2023/common"
	model_solution "github.com/FACorreiaa/aoc-2023/model"
	"github.com/charmbracelet/bubbles/key"
	"github.com/charmbracelet/bubbles/list"
	tea "github.com/charmbracelet/bubbletea"
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
	var cmd tea.Cmd

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
				mappedValues := map[string]func() tea.Msg{
					"Day 1": dayone.Start,
					// Add more entries as needed
				}

				selectedItem := m.list.SelectedItem()
				startFn, _ := mappedValues[selectedItem.FilterValue()]

				solutionModel := InitSolution(selectedItem.FilterValue(), common.P, startFn)
				return solutionModel.Update(common.WindowSize)
			} else {
				m.subMenu = true
			}
			return m, nil
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
		default:
			m.list, cmd = m.list.Update(msg)
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

//func Start() error {
//	rand.NewSource(time.Now().UTC().UnixNano())
//
//	m, _ := InitProject()
//
//	common.P = tea.NewProgram(m, tea.WithAltScreen())
//
//	if _, err := common.P.Run(); err != nil {
//		fmt.Println("Error running program:", err)
//		os.Exit(1)
//	}
//
//	return nil
//}
