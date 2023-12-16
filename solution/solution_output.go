package solution

import (
	"fmt"
	"github.com/FACorreiaa/aoc-2023/cmd/settings"
)

const divider = "---"

func FormatSolution(day settings.Day) string {
	return fmt.Sprintf("Title: %s\n Result: %d %s\n", day.Title, day.Result, divider)
}
