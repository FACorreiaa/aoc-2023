package dayeight

import (
	"github.com/FACorreiaa/aoc-2023/cmd/settings"
	"log"
	"regexp"
	"strings"
	"time"
)

var items = regexp.MustCompile(`(\w+) = \((\w+), (\w+)\)`)

func partOne(s string) int64 {
	var score int64 = 0

	lines := strings.Split(s, "\n\n")

	//directions := lines[0]

	//fmt.Printf("%#v", lines)

	m := map[string][]string{}
	for _, line := range strings.Split(lines[1], "\n") {
		match := items.FindStringSubmatch(line)

		m[match[1]] = []string{match[2], match[3]}
	}

	local := "AAA"

	for local != `ZZZ` {
		for _, d := range lines[0] {
			score++
			if d == 'R' {
				local = m[local][1]
			} else {
				local = m[local][0]
			}

			if local == `ZZZ` {
				break
			}
		}
	}

	return score
}

type node struct {
	Name      string
	LeftNode  *node
	RightNode *node
	z         bool
	a         bool
}

// partThree same logic as part two but optimised
func partThree(s string) int64 {
	lines := strings.Split(s, "\n\n")
	m := map[string][]string{}
	nodes := map[string]*node{}

	for _, line := range strings.Split(lines[1], "\n") {
		match := items.FindStringSubmatch(line)
		m[match[1]] = []string{match[2], match[3]}
	}

	for k, _ := range m {
		nodes[k] = &node{Name: k, z: strings.Contains(k, `Z`), a: strings.Contains(k, `A`)}
	}

	for k, _ := range nodes {
		nodes[k].LeftNode = nodes[m[k][0]]
		nodes[k].RightNode = nodes[m[k][1]]

	}

	activeNodes := make([]*node, 0)
	for k, v := range nodes {
		if strings.Contains(k, `A`) {
			activeNodes = append(activeNodes, v)
		}
	}

	iterations := make([][]int, 0)
	for _, n := range activeNodes {
		dir := lines[0]
		iteration := make([]int, 0)
		count := 0

		var first *node
		for {
			for count == 0 || !strings.Contains(n.Name, `Z`) {
				count += 1

				if dir[0] == 'R' {
					n = n.RightNode
				} else {
					n = n.LeftNode
				}

				dir = dir[1:] + string(dir[0])
			}
			iteration = append(iteration, count)
			if first == nil {
				first = n
				count = 0
			} else if n == first {
				break
			}
		}
		iterations = append(iterations, iteration)
	}
	nums := make([]int, 0)

	for _, iteration := range iterations {
		nums = append(nums, iteration[0])
	}

	lcm := nums[0]

	for _, n := range nums {
		lcm = lcm * n / gcd(lcm, n)
	}

	return int64(lcm)
}

func gcd(a, b int) int {
	for b != 0 {
		a, b = b, a%b
	}
	return a
}

func Start() {
	lines := settings.GetLines("./cmd/day-eight/wasteland.txt")
	for _, line := range lines {
		println(line)
	}
	partOneStart := time.Now()
	partOneResult := partOne(strings.Join(lines, "\n"))
	loPrint("Result: ", partOneResult)
	loPrint("Day seven part one took: ", time.Since(partOneStart))

	partThreeStart := time.Now()
	partThreeResult := partThree(strings.Join(lines, "\n"))
	loPrint("Result: ", partThreeResult)
	loPrint("Day seven part two took: ", time.Since(partThreeStart))
}
