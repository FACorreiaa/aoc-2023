package main

import (
	"github.com/FACorreiaa/aoc-2023/tui"
	tea "github.com/charmbracelet/bubbletea"
	"log"
	"math/rand"
	"os"
	"time"
)

func main() {
	rand.NewSource(time.Now().UTC().UnixNano())

	if _, err := tea.NewProgram(tui.NewModel()).Run(); err != nil {
		log.Print("Error running program:", err)
		os.Exit(1)
	}
	//dayseven.Start()
	//if err := tui.StartTea(); err != nil {
	//	log.Print("Error running program:", err)
	//	os.Exit(1)
	//}
}
