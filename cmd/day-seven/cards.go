package dayseven

import (
	"cmp"
	"github.com/FACorreiaa/aoc-2023/cmd/common"
	"log"
	"slices"
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

func compareHands(a, b hand) int {
	// Compare bestHand first
	if a.bestHand != b.bestHand {
		return cmp.Compare(a.bestHand, b.bestHand)
	}

	// If bestHand is the same, compare cards element-wise
	for k := range a.cards {
		if a.cards[k] != b.cards[k] {
			return cmp.Compare(strength[a.cards[k]], strength[b.cards[k]])
		}
	}

	// If all elements are the same, return 0
	return 0
}

// GetBestHand check the best bid
func GetBestHand(hand string) int {
	var handType = map[string]int{
		"highCard":     1,
		"onePair":      2,
		"twoPair":      3,
		"threeOfAKind": 4,
		"fullHouse":    5,
		"fourOfAKind":  6,
		"fiveOfAKind":  7,
	}

	var cards []string
	for _, char := range hand {
		cards = append(cards, string(char))
	}

	var uniqueCards []string

	for _, card := range cards {
		if !slices.Contains(uniqueCards, card) {
			uniqueCards = append(uniqueCards, card)
		}
	}

	if len(uniqueCards) == 1 {
		return handType["fiveOfAKind"]
	}

	if len(uniqueCards) == 2 {
		for _, card := range uniqueCards {
			if strings.Count(hand, card) == 4 {
				return handType["fourOfAKind"]
			} else if strings.Count(hand, card) == 3 {
				return handType["fullHouse"]
			}
		}
	}

	if len(uniqueCards) == 3 {
		for _, card := range uniqueCards {
			if strings.Count(hand, card) == 3 {
				return handType["threeOfAKind"]
			} else if strings.Count(hand, card) == 2 {
				return handType["twoPair"]
			}
		}
	}

	if len(uniqueCards) == 4 {
		for _, card := range uniqueCards {
			if strings.Count(hand, card) == 2 {
				return handType["onePair"]
			}
		}
	}

	if len(uniqueCards) == 5 {
		return handType["highCard"]
	}

	return -1
}

func partOne(s string, jokerIsWild bool) int64 {
	var score int64 = 0
	hands := make([]hand, 0)

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
		h.bestHand = GetBestHand(h.cards)

		if jokerIsWild && strings.Contains(h.cards, `J`) {
			h.bestHand = playWilds(h)
		}
		hands = append(hands, h)
	}

	slices.SortFunc(hands, compareHands)

	for i, h := range hands {
		score += int64((i + 1) * h.bid)
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

// partOne MUST REFACTOR WITH slices.SortFunc() and cmp
//func partOne(s string, jokerIsWild bool) int64 {
//	var score int64 = 0
//	hands := make([]hand, 0)
//
//	if jokerIsWild {
//		strength['J'] = -1
//	} else {
//		strength['J'] = 11
//	}
//
//	for _, line := range strings.Split(s, "\n") {
//
//		split := strings.Split(line, ` `)
//		bid, err := strconv.Atoi(split[1])
//
//		if err != nil {
//			common.HandleError(err, "Error converting string")
//		}
//
//		h := hand{cards: split[0], bid: bid}
//		h.bestHand = GetBestHand(h.cards)
//
//		if jokerIsWild && strings.Contains(h.cards, `J`) {
//			h.bestHand = playWilds(h)
//		}
//		hands = append(hands, h)
//	}
//
//	sort.Slice(hands, func(i, j int) bool {
//		if hands[i].bestHand == hands[j].bestHand {
//			for k := range hands[i].cards {
//				if hands[i].cards[k] == hands[j].cards[k] {
//					continue
//				}
//
//				return strength[hands[i].cards[k]] < strength[hands[j].cards[k]]
//			}
//
//		}
//
//		return hands[i].bestHand < hands[j].bestHand
//	})
//
//	//fmt.Printf("%#v\n", hands)
//
//	for i, hand := range hands {
//		score += int64((i + 1) * hand.bid)
//	}
//
//	return score
//}

//study later
//func GetBestCards(cards string) int {
//	m := map[rune]int{}
//	for _, r := range cards {
//		m[r] += 1
//	}
//
//	if len(m) == 1 {
//		return FiveOfAKind
//	}
//
//	if len(m) == 2 {
//		for _, v := range m {
//			if v == 4 {
//				return FourOfAKind
//			} else if v == 3 {
//				return FullHouse
//			}
//		}
//
//	}
//
//	if len(m) == 3 {
//		for _, v := range m {
//			if v == 3 {
//				return ThreeOfAKind
//			}
//			return TwoPair
//		}
//	}
//
//	if len(m) == 5 {
//		return HighCard
//	}
//
//	return OnePair
//}

func Start() {
	lines := common.GetLines("./cmd/day-seven/cards.txt")
	for _, line := range lines {
		println(line)
	}
	partOneWithSlicesStart := time.Now()
	partOneWithSlicesResult := partOne(strings.Join(lines, "\n"), false)
	log.Print("Result: ", partOneWithSlicesResult)
	log.Print("Day seven part one with slices took with no Joker: ", time.Since(partOneWithSlicesStart))
	partTwoWithSlicesStart := time.Now()
	partTwoWithSlicesResult := partOne(strings.Join(lines, "\n"), true)
	log.Print("Result: ", partTwoWithSlicesResult)
	log.Print("Day seven part one with slices took with Joker: ", time.Since(partTwoWithSlicesStart))
}
