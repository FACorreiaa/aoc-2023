package model_solution

import (
	"fmt"
	dayone "github.com/FACorreiaa/aoc-2023/cmd/day-one"
	"github.com/FACorreiaa/aoc-2023/common"
	"github.com/FACorreiaa/aoc-2023/messages"
	"github.com/charmbracelet/bubbles/key"
	tea "github.com/charmbracelet/bubbletea"
)

// SolutionModel is the interface that each solution model should implement.
type SolutionModel interface {
	tea.Model
	StartProcessing() tea.Cmd
}

// SolutionModelBase is a basic implementation of SolutionModel.
type SolutionModelBase struct {
	Title  string
	Result tea.Msg
}

// Mock processing, TODO

// StartProcessing implements the StartProcessing method for SolutionModelBase.
func (m *SolutionModelBase) StartProcessing() tea.Cmd {
	return func() tea.Msg {

		return messages.SolutionMsg(fmt.Sprintf("Day %s\n\nResult: %d\n\nPress 'q' to return to the menu", m.Title))
	}
}

func (s *SolutionModelBase) View() string {
	return common.DocStyle.Render(s.View() + "\n")
}

//func DayOneStart() SolutionModel {
//	return &SolutionModelBase{Title: "1", Result: dayone.Start()}
//}

// View returns the string representation of SolutionModelBase.
func (m *DayOneModel) View() string {
	return m.SolutionModelBase.View() // Reuse the base view function
}

type DayOneModel struct {
	SolutionModelBase
	StartFn func() tea.Msg
}

func DayOneStart(dayTitle string, startFn func() tea.Msg) (SolutionModel, error) {
	return &DayOneModel{
		SolutionModelBase: SolutionModelBase{
			Title:  dayTitle,
			Result: dayone.Start(),
		},
		StartFn: startFn,
	}, nil
}

func (m *DayOneModel) Init() tea.Cmd {
	return nil
}

func (m *DayOneModel) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch {
		case key.Matches(msg, common.Keymap.Enter):
			return m, m.StartProcessing() // Initiates processing and returns the result
		case key.Matches(msg, common.Keymap.Back):
			return m, nil // Dismiss the solution and return to the menu
		}
	}

	return m, nil
}
