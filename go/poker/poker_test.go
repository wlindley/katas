package poker_test

import (
	"strconv"
	"testing"

	"github.com/wlindley/katas/go/poker"
)

func TestNewHand(t *testing.T) {
	tt := []handTestCase{
		validHand("Valid Hand", c("CJ"), c("HJ"), c("DQ"), c("SK"), c("CA")),
		invalidHand("Empty Hand"),
		invalidHand("Too Small Hand", c("C2"), c("C3"), c("C4"), c("C5")),
		invalidHand("Too Large Hand", c("C2"), c("C3"), c("C4"), c("C5"), c("C6"), c("C7")),
		invalidHand("Too Low Card", c("C1"), c("C2"), c("C3"), c("C4"), c("C5")),
		invalidHand("Too High Card", c("C15"), c("C2"), c("C3"), c("C4"), c("C5")),
		invalidHand("Duplicate Cards", c("H2"), c("H2"), c("C3"), c("S4"), c("C5")),
	}

	for _, x := range tt {
		t.Run(x.Name, func(t *testing.T) {
			_, err := poker.NewHand(x.Hand...)
			if err == nil && !x.IsValid {
				t.Errorf("%v should have failed to create a hand, but did not", x.Hand)
			} else if err != nil && x.IsValid {
				t.Errorf("%v should have successfully created a hand, but instead errored: %s", x.Hand, err)
			}
		})
	}
}

func TestCompare(t *testing.T) {
	tt := []compareTestCase{
		firstWins("First Player High Card",
			[]poker.Card{c("C2"), c("C5"), c("S7"), c("S9"), c("SA")},
			[]poker.Card{c("HK"), c("H9"), c("D7"), c("D5"), c("H3")}),
		secondWins("Second Player High Card",
			[]poker.Card{c("C2"), c("C5"), c("S7"), c("S9"), c("SK")},
			[]poker.Card{c("HA"), c("H9"), c("D7"), c("D5"), c("H3")}),
		tie("Tie High Card",
			[]poker.Card{c("C2"), c("C5"), c("S7"), c("S9"), c("SA")},
			[]poker.Card{c("HA"), c("H9"), c("D7"), c("D5"), c("H2")}),
		firstWins("First Player Second Highest Card",
			[]poker.Card{c("C2"), c("C5"), c("S7"), c("SJ"), c("SK")},
			[]poker.Card{c("HK"), c("H9"), c("D7"), c("D5"), c("H3")}),
	}

	for _, x := range tt {
		t.Run(x.Name, func(t *testing.T) {
			actual := x.FirstHand.Compare(x.SecondHand)
			if actual != x.Expected {
				t.Errorf("compared %v and %v, got %s, but expected %s", x.FirstHand, x.SecondHand, actual, x.Expected)
			}
		})
	}
}

func c(def string) poker.Card {
	v, err := strconv.Atoi(def[1:])
	value := uint(v)
	if err != nil {
		switch def[1:] {
		case "J", "j":
			value = poker.Jack
		case "Q", "q":
			value = poker.Queen
		case "K", "k":
			value = poker.King
		case "A", "a":
			value = poker.Ace
		}
	}
	switch def[0:1] {
	case "C":
		return poker.Club(value)
	case "S":
		return poker.Spade(value)
	case "H":
		return poker.Heart(value)
	case "D":
		return poker.Diamond(value)
	}
	return poker.Card{}
}

type handTestCase struct {
	Name    string
	Hand    []poker.Card
	IsValid bool
}

func validHand(name string, cards ...poker.Card) handTestCase {
	return handTestCase{
		Name:    name,
		Hand:    cards,
		IsValid: true,
	}
}

func invalidHand(name string, cards ...poker.Card) handTestCase {
	return handTestCase{
		Name:    name,
		Hand:    cards,
		IsValid: false,
	}
}

type compareTestCase struct {
	Name       string
	FirstHand  *poker.Hand
	SecondHand *poker.Hand
	Expected   poker.Winner
}

func firstWins(name string, firstCards, secondCards []poker.Card) compareTestCase {
	return createCompareTestCase(name, firstCards, secondCards, poker.First)
}

func secondWins(name string, firstCards, secondCards []poker.Card) compareTestCase {
	return createCompareTestCase(name, firstCards, secondCards, poker.Second)
}

func tie(name string, firstCards, secondCards []poker.Card) compareTestCase {
	return createCompareTestCase(name, firstCards, secondCards, poker.Tie)
}

func createCompareTestCase(name string, firstCards, secondCards []poker.Card, winner poker.Winner) compareTestCase {
	firstHand, _ := poker.NewHand(firstCards...)
	secondHand, _ := poker.NewHand(secondCards...)
	return compareTestCase{
		Name:       name,
		FirstHand:  firstHand,
		SecondHand: secondHand,
		Expected:   winner,
	}
}
