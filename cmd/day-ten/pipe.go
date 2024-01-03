package dayten

import (
	"fmt"
	"github.com/FACorreiaa/aoc-2023/cmd/settings"
	"log"
	"strings"
	"time"
)

const (
	vertical   = '|'
	horizontal = '-'
	NE         = 'L'
	NW         = 'J'
	SW         = '7'
	SE         = 'F'
	ground     = '.'
	start      = 'S'
)

func isHairpin(a, b rune) bool {
	switch {
	case a == NE && b == NW:
		fallthrough
	case a == SE && b == SW:
		return true
	default:
		return false
	}
}

var directionChange = map[rune]map[rune]rune{
	'E': {
		horizontal: 'E',
		SW:         'S',
		NW:         'N',
		'S':        'D', // end marker
	},
	'N': {
		vertical: 'N',
		SW:       'W',
		SE:       'E',
		'S':      'D', //end marker
	},
	'S': {
		vertical: 'S',
		NW:       'W',
		NE:       'E',
		'S':      'D', // end marker
	},
	'W': {
		horizontal: 'W',
		SE:         'S',
		NE:         'N',
		'S':        'D', // end marker
	},
}

//var grid [][]byte
//var visited = make([]bool, 20000)
//var distances = make([]int, 20000)

//func (g *Grid) parseFile(data []byte) {
//	grid = grid[:0]
//	visited = make([]bool, 20000)
//	distances = make([]int, 20000)
//
//	for y, line := range strings.Split(string(data), "\n") {
//		grid = append(grid, make([]byte, 0))
//
//		for x, c := range line {
//			if c == 'S' {
//				startx = x
//				starty = y
//				c = '-'
//			}
//			grid[y] = append(grid[y], byte(c))
//		}
//	}
//}

//func (g *Grid) partOne() int64 {
//	dfs(startx, starty, 0)
//
//	maxVal := 0
//	for _, d := range distances {
//		if d > maxVal {
//			maxVal = d
//		}
//	}
//
//	return int64(maxVal/2 + 1)
//}

//func partOne(input []byte) int {
//	var b = bufio.NewReader(bytes.NewBuffer(input))
//
//	var startx, starty int
//
//	grid = grid[:0]
//
//	var y int
//	for {
//		line, err := b.ReadBytes('\n')
//		if err != nil {
//			break
//		}
//
//		if startx == 0 && starty == 0 {
//			for x, c := range line {
//				if c == 'S' {
//					startx = x
//					starty = y
//					line[x] = '-'
//					break
//				}
//			}
//		}
//
//		grid = append(grid, line[:len(line)-1])
//
//		y++
//	}
//
//	// Implement BFS to find the farthest node from the start
//	// We start from the start node, and we assign all nodes we can go to a distance of 1
//	// Then we go to those nodes, and assign all nodes we can go to a distance of 2
//	// We can determine neighbors this way :
//	//   If the node is a -, we can go to the node on the left or right
//	//   If the node is a |, we can go to the node above or below
//	//   If the node is a L, we can go to the node on the right or above
//	//   If the node is a J, we can go to the node on the left or above
//	//   If the node is a 7, we can go to the node on the left or below
//	//   If the node is a F, we can go to the node on the right or below
//
//	// Reset the arrays
//	// clear(visited)
//	// clear(distances)
//
//	dfs(startx, starty, 0)
//
//	// Find the farthest node
//	maxValues := 0
//	for _, d := range distances {
//		if d > maxValues {
//			maxValues = d
//		}
//	}
//
//	return maxValues/2 + 1
//}

//func coordToIndex(x, y int) int {
//	return y*len(grid[0]) + x
//}
//
//func dfs(x, y, d int) {
//	// If we have already visited this node, return
//	if visited[coordToIndex(x, y)] {
//		return
//	}
//
//	// Mark this node as visited
//	visited[coordToIndex(x, y)] = true
//	distances[coordToIndex(x, y)] = d
//
//	// If this node is a -, we can go to the node on the left or right
//	if grid[y][x] == '-' {
//		// Go to the node on the left
//		if x > 0 {
//			dfs(x-1, y, d+1)
//		}
//
//		// Go to the node on the right
//		if x < len(grid[y])-1 {
//			dfs(x+1, y, d+1)
//		}
//	}
//
//	// if this node is  |, we go to the node above or below
//	if grid[y][x] == '|' {
//		//go to node above
//		if y > 0 {
//			dfs(x, y-1, d+1)
//		}
//
//		//go to node below
//		if y < len(grid)-1 {
//			dfs(x, y+1, d+1)
//		}
//	}
//
//	// If this node is a L, we can go to the node on the right or above
//	if grid[y][x] == 'L' {
//		//go to the node on the right
//		if x < len(grid[y]) {
//			dfs(x+1, y, d+1)
//		}
//
//		//go to node above
//		if y > 0 {
//			dfs(x, y-1, d+1)
//		}
//	}
//
//	// If this node is a J, we can go to node on left or above
//	if grid[y][x] == 'J' {
//		//go left
//		if x > 0 {
//			dfs(x-1, y, d+1)
//		}
//
//		//go above
//		if y > 0 {
//			dfs(x, y-1, d+1)
//		}
//	}
//
//	//if this node is a 7, we can go to node below or left
//	if grid[y][x] == '7' {
//		//go left
//		if x > 0 {
//			dfs(x-1, y, d+1)
//		}
//
//		//go below
//		if y < len(grid)-1 {
//			dfs(x, y+1, d+1)
//		}
//	}
//
//	// if this node is F, we can go to node right or below
//	if grid[y][x] == 'F' {
//		//go right
//		if x < len(grid[y])-1 {
//			dfs(x+1, y, d+1)
//		}
//
//		//go below
//		if y < len(grid)-1 {
//			dfs(x, y+1, d+1)
//		}
//	}
//
//}

type maze []string

func (m maze) Print() {
	for _, line := range m {
		fmt.Println(line)
	}
}

func (m maze) findStart() (x, y int) {
	for y, line := range m {
		for x, char := range line {
			if char == start {
				return x, y
			}
		}
	}

	panic("no start found")
}

func (m maze) startDirection() rune {
	startX, startY := m.findStart()
	bounds := len(m) - 1

	for _, direction := range []rune{'N', 'E', 'S', 'W'} {
		newX, newY := startX, startY
		switch direction {
		case 'N':
			newY--
		case 'E':
			newX++
		case 'S':
			newY++
		case 'W':
			newX--
		}

		if newX < 0 || newY < 0 || newX > bounds || newY > bounds {
			fmt.Println("out of bounds", newX, newY, bounds)
			continue
		}

		pipe := m[newY][newX]

		if directionChange[direction][rune(pipe)] != 0 {
			return direction
		}
	}

	panic("no start direction")
}

func (m maze) traverse(startX, startY int, direction rune) (newX, newY int, newDirection rune) {
	newX, newY = startX, startY
	switch direction {
	case 'N':
		newY--
	case 'E':
		newX++
	case 'S':
		newY++
	case 'W':
		newX--
	}

	bounds := len(m) - 1
	if newX < 0 || newY < 0 || newX > bounds || newY > bounds {
		panic("OUT OF BOUNDS")
	}

	pipe := m[newY][newX]
	if directionChange[direction][rune(pipe)] == 0 {
		panic("invalid puzzle!!!!")
	}
	return newX, newY, directionChange[direction][rune(pipe)]
}

func partOne(input []string) int {
	m := maze(input)

	startX, startY := m.findStart()
	direction := m.startDirection()

	x, y := startX, startY
	steps := 0
	for {
		steps++
		x, y, direction = m.traverse(x, y, direction)
		if x == startX && y == startY {
			break
		}
	}
	return steps / 2
}

func partTwo(input []string) int64 {
	m := maze(input)

	startX, startY := m.findStart()
	direction := m.startDirection()

	visited := make(map[[2]int]bool)
	visited[[2]int{startX, startY}] = true

	x, y := startX, startY
	for {
		x, y, direction = m.traverse(x, y, direction)
		visited[[2]int{x, y}] = true
		if x == startX && y == startY {
			break
		}
	}

	m[startY] = strings.ReplaceAll(m[startY], "S", "J")
	var count int64 = 0

	for y, line := range m {
		var inside bool
		var previousChar rune

		for x, c := range line {
			v := visited[[2]int{x, y}]

			if !v {
				previousChar = '.'
				if inside {
					count++
				}
				continue
			}

			switch c {
			case '|':
				inside = !inside
			case '-':
			default:
				if previousChar == '.' {
					previousChar = c
					inside = !inside
					continue
				}

				if isHairpin(previousChar, c) {
					inside = !inside
				}

				previousChar = '.' //corner doesn't exist
			}
		}
	}
	return count
}

func Start() {
	lines := settings.GetLines("./cmd/day-ten/pipe.txt")
	for _, line := range lines {
		println(line)
	}

	partOneStart := time.Now()
	log.Print("Result: ", partOne(lines))
	log.Print("Day seven part one took: ", time.Since(partOneStart))

	partTwoStart := time.Now()
	log.Print("Result: ", partTwo(lines))
	log.Print("Day seven part two took: ", time.Since(partTwoStart))
}
