package main

import (
	"fmt"
	"github.com/FACorreiaa/aoc-2023/common"
	"github.com/FACorreiaa/aoc-2023/tui"
	tea "github.com/charmbracelet/bubbletea"
	"math/rand"
	"os"
	"time"
)

func main() {
	rand.NewSource(time.Now().UTC().UnixNano())

	m, _ := tui.InitProject()

	common.P = tea.NewProgram(m, tea.WithAltScreen())

	if _, err := common.P.Run(); err != nil {
		fmt.Println("Error running program:", err)
		os.Exit(1)
	}
}
