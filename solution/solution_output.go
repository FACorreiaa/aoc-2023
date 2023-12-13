package solution

import (
	"fmt"
	modelSolution "github.com/FACorreiaa/aoc-2023/model"
)

const divider = "---"

func FormatSolution(solution modelSolution.SolutionModelBase) string {
	return fmt.Sprintf("Title: %s\n Result: %d %s\n", solution.Title, solution.Result, divider)
}
