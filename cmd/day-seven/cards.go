package dayseven

import (
	"github.com/FACorreiaa/aoc-2023/cmd/common"
	"log"
	"sort"
	"strconv"
	"strings"
	"time"
)

var strength = map[byte]int{
	'A': 14,
	'K': 13,
	'Q': 12,
	'J': 11,
	'T': 10,
	'9': 9,
	'8': 8,
	'7': 7,
	'6': 6,
	'5': 5,
	'4': 4,
	'3': 3,
	'2': 2,
}

type hand struct {
	cards    string
	bid      int
	bestHand int
}

const (
	FiveOfAKind  = 7
	FourOfAKind  = 6
	FullHouse    = 5
	ThreeOfAKind = 4
	TwoPair      = 3
	OnePair      = 2
	HighCard     = 1
)

func GetBestCards(cards string) int {
	m := map[rune]int{}
	for _, r := range cards {
		m[r] += 1
	}

	if len(m) == 1 {
		return FiveOfAKind
	}

	if len(m) == 2 {
		for _, v := range m {
			if v == 4 {
				return FourOfAKind
			}
		}
		return FullHouse
	}

	if len(m) == 3 {
		for _, v := range m {
			if v == 3 {
				return ThreeOfAKind
			}
			return TwoPair
		}
	}

	if len(m) == 5 {
		return HighCard
	}

	return OnePair
}

func partOne(s string, jokerIsWild bool) int64 {
	var score int64 = 0
	hands := []hand{}

	if jokerIsWild {
		strength['J'] = -1
	} else {
		strength['J'] = 11
	}

	for _, line := range strings.Split(s, "\n") {
		split := strings.Split(line, ` `)
		bid, err := strconv.Atoi(split[1])

		if err != nil {
			common.HandleError(err, "Error converting string")
		}

		h := hand{cards: split[0], bid: bid}
		h.bestHand = GetBestCards(h.cards)

		if jokerIsWild && strings.Contains(h.cards, `J`) {
			h.bestHand = playWilds(h)
		}
		hands = append(hands, h)
	}

	sort.Slice(hands, func(i, j int) bool {
		if hands[i].bestHand == hands[j].bestHand {
			for k := range hands[i].cards {
				if hands[i].cards[k] == hands[j].cards[k] {
					continue
				}

				return strength[hands[i].cards[k]] < strength[hands[j].cards[k]]
			}
			
		}
		return hands[i].bestHand < hands[j].bestHand
	})

	//fmt.Printf("%#v\n", hands)

	for i, hand := range hands {
		score += int64((i + 1) * hand.bid)
	}

	return score
}

func playWilds(h hand) int {
	m := map[rune]int{}
	for _, r := range h.cards {
		m[r] += 1
	}

	if m['J'] >= 4 {
		return FiveOfAKind
	}

	if m['J'] == 3 {
		if len(m) == 2 {
			return FiveOfAKind
		}
		return FourOfAKind
	}

	if m['J'] == 2 {
		if h.bestHand == TwoPair {
			return FourOfAKind
		}

		if h.bestHand == OnePair {
			return ThreeOfAKind
		}

		if h.bestHand == FullHouse {
			return FiveOfAKind
		}
	}

	if m['J'] == 1 {
		if h.bestHand == ThreeOfAKind {
			return FourOfAKind
		}

		if h.bestHand == TwoPair {
			return FullHouse
		}

		if h.bestHand == OnePair {
			return ThreeOfAKind
		}

		if h.bestHand == FourOfAKind {
			return FiveOfAKind
		}
	}

	return OnePair
}

func Start() {
	lines := common.GetLines("./cmd/day-seven/cards.txt")
	for _, line := range lines {
		println(line)
	}
	partOneStart := time.Now()
	partOneResult := partOne(strings.Join(lines, "\n"), false)
	log.Print("Result: ", partOneResult)
	log.Print("Day seven part one took: ", time.Since(partOneStart))
	partTwoStart := time.Now()
	partTwoResult := partOne(strings.Join(lines, "\n"), true)
	log.Print("Result: ", partTwoResult)
	log.Print("Day seven part two took: ", time.Since(partTwoStart))

}
