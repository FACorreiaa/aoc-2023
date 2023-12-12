package solution

import (
	"fmt"
	model_solution "github.com/FACorreiaa/aoc-2023/model"
)

const divider = "---"

func FormatSolution(solution model_solution.SolutionModelBase) string {
	return fmt.Sprintf("Title: %s\n Result: %d %s\n", solution.Title, solution.Result, divider)
}
