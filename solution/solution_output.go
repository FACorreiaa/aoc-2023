package solution

import (
	"fmt"
	"github.com/FACorreiaa/aoc-2023/common"
)

const divider = "---"

func FormatSolution(day common.Day) string {
	return fmt.Sprintf("Title: %s\n Result: %d %s\n", day.Title, day.Result, divider)
}
