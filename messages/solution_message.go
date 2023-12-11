package messages

import tea "github.com/charmbracelet/bubbletea"

type SolutionTransitionMsg struct {
	DayTitle string
	StartFn  func() tea.Msg
}

// SolutionMsg represents the messages that can be sent to any SolutionModel.
type SolutionMsg string

const (
	Dismiss SolutionMsg = "dismiss"
)
