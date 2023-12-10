package p1

import (
	"regexp"
	"strconv"
)

type hand struct {
	hand string
	Bid  int
}

var priority = map[string]int{
	"A": 13,
	"K": 12,
	"Q": 11,
	"J": 10,
	"T": 9,
	"9": 8,
	"8": 7,
	"7": 6,
	"6": 5,
	"5": 4,
	"4": 3,
	"3": 2,
	"2": 1,
}

func (h hand) value() int {
	ofAKind := map[rune]int{}

	var (
		fiveOfAKind  bool
		fourOfAKind  bool
		fullhouse    bool
		threeOfAKind bool
		twoPair      bool
		onePair      bool
	)

	for _, char := range h.hand {
		ofAKind[char]++

		switch {
		case !onePair && ofAKind[char] == 2:
			onePair = true
		case onePair && ofAKind[char] == 2:
			twoPair = true
		case ofAKind[char] == 3:
			threeOfAKind = true
		case ofAKind[char] == 4:
			fourOfAKind = true
		case ofAKind[char] == 5:
			fiveOfAKind = true
		}
	}

	if threeOfAKind && twoPair {
		fullhouse = true
	}

	switch {
	case fiveOfAKind:
		return 7
	case fourOfAKind:
		return 6
	case fullhouse:
		return 5
	case threeOfAKind:
		return 4
	case twoPair:
		return 3
	case onePair:
		return 2
	}

	return 1
}

func (a hand) less(b hand) bool {
	if a == b {
		return false
	}

	aVal := a.value()
	bVal := b.value()
	switch {
	case aVal < bVal:
		return true
	case aVal > bVal:
		return false
	}

	for idx := range a.hand {
		switch {
		case a.hand[idx:idx+1] == b.hand[idx:idx+1]:
			continue
		case priority[a.hand[idx:idx+1]] < priority[b.hand[idx:idx+1]]:
			return true
		case priority[a.hand[idx:idx+1]] > priority[b.hand[idx:idx+1]]:
			return false
		}
	}

	return false
}

type ByValue []hand

func (a ByValue) Len() int           { return len(a) }
func (a ByValue) Swap(i, j int)      { a[i], a[j] = a[j], a[i] }
func (a ByValue) Less(i, j int) bool { return a[i].less(a[j]) }

func ParseHandsForPartOne(input string) []hand {
	pattern := regexp.MustCompile(`(\w+) (\d+)`)

	matches := pattern.FindAllStringSubmatch(input, -1)
	result := make([]hand, len(matches))

	for idx, match := range matches {
		result[idx].hand = match[1]
		result[idx].Bid, _ = strconv.Atoi(match[2])
	}

	return result
}
