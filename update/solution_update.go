package update

import (
	"github.com/FACorreiaa/aoc-2023/common"
	"github.com/FACorreiaa/aoc-2023/model"
	"github.com/charmbracelet/bubbles/key"
	"github.com/charmbracelet/bubbletea"
)

// SolutionUpdateFunc is a function type that represents the update function for a solution model.
type SolutionUpdateFunc func(model_solution.SolutionModel, tea.Msg) (model_solution.SolutionModel, tea.Cmd)

// SolutionUpdate is the update function for a solution model.
func SolutionUpdate(model model_solution.SolutionModel, msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch {
		//case messages.SolutionTransitionMsg:
		//	// Transition to the selected solution model
		//	solutionModel := model_solution.DayOneStart(msg.DayTitle, msg.StartFn)
		//	return solutionModel, nil
		case key.Matches(msg, common.Keymap.Enter):
			model.StartProcessing() // Initiates processing and returns the result
			return model, nil
		case key.Matches(msg, common.Keymap.Back):
			return model, nil // Dismiss the solution and return to the menu
		}
	}

	return model, nil
}
