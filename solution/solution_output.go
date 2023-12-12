package solution

import (
	"fmt"
	model_solution "github.com/FACorreiaa/aoc-2023/model"
)

const divider = "---"

func FormatEntry(entry model_solution.SolutionModelBase) string {
	return fmt.Sprintf("ID: %d\nCreated: %s\nMessage:\n\n %s\n %s\n", entry.Title, entry.Result, divider)
}
