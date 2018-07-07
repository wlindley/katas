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
	//TODO: test high card fallback when differing card is not highest card
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
		secondWins("Second Player Pair",
			[]poker.Card{c("SA"), c("HK"), c("DJ"), c("C10"), c("S9")},
			[]poker.Card{c("C2"), c("H2"), c("S3"), c("D4"), c("C5")}),
		firstWins("First Player Pair",
			[]poker.Card{c("C2"), c("H2"), c("S3"), c("D4"), c("C5")},
			[]poker.Card{c("SA"), c("HK"), c("DJ"), c("C10"), c("S9")}),
		firstWins("Tied Pairs Falls Back to High Card",
			[]poker.Card{c("C2"), c("S3"), c("CA"), c("S9"), c("C9")},
			[]poker.Card{c("H2"), c("D3"), c("HK"), c("D9"), c("H9")}),
		firstWins("First Player Two Pair",
			[]poker.Card{c("C2"), c("H2"), c("S3"), c("D3"), c("C5")},
			[]poker.Card{c("SA"), c("HA"), c("DJ"), c("C10"), c("S9")}),
		secondWins("Second Player Two Pair",
			[]poker.Card{c("SA"), c("HA"), c("DJ"), c("C10"), c("S9")},
			[]poker.Card{c("C2"), c("H2"), c("S3"), c("D3"), c("C5")}),
		secondWins("Tied Two Pair Falls Back to High Card",
			[]poker.Card{c("C2"), c("H2"), c("C3"), c("H3"), c("C5")},
			[]poker.Card{c("S2"), c("D2"), c("S3"), c("D3"), c("SJ")}),
		firstWins("Tied Two Pair Falls Back to Remaining Card Even if it's Not the Highest",
			[]poker.Card{c("CA"), c("HA"), c("CK"), c("HK"), c("CQ")},
			[]poker.Card{c("SA"), c("DA"), c("SK"), c("DK"), c("SJ")}),
		firstWins("First Player Three of a Kind",
			[]poker.Card{c("C2"), c("H2"), c("S2"), c("D3"), c("C5")},
			[]poker.Card{c("SA"), c("HA"), c("DJ"), c("CJ"), c("S9")}),
		secondWins("Second Player Three of a Kind",
			[]poker.Card{c("SA"), c("HA"), c("DJ"), c("CJ"), c("S9")},
			[]poker.Card{c("C2"), c("H2"), c("S2"), c("D3"), c("C5")}),
		firstWins("First Player Straight",
			[]poker.Card{c("C2"), c("H3"), c("S4"), c("D5"), c("C6")},
			[]poker.Card{c("SA"), c("HA"), c("DA"), c("C10"), c("S9")}),
		secondWins("Second Player Straight",
			[]poker.Card{c("SA"), c("HA"), c("DA"), c("C10"), c("S9")},
			[]poker.Card{c("C2"), c("H3"), c("S4"), c("D5"), c("C6")}),
		secondWins("Both Straights Falls Back to High Card",
			[]poker.Card{c("C2"), c("H3"), c("C4"), c("H5"), c("C6")},
			[]poker.Card{c("S3"), c("D4"), c("S5"), c("D6"), c("S7")}),
		firstWins("First Player Flush",
			[]poker.Card{c("SA"), c("S2"), c("S8"), c("SQ"), c("S9")},
			[]poker.Card{c("C2"), c("H3"), c("S4"), c("D5"), c("C6")}),
		secondWins("Second Player Flush",
			[]poker.Card{c("C2"), c("H3"), c("S4"), c("D5"), c("C6")},
			[]poker.Card{c("SA"), c("S2"), c("S8"), c("SQ"), c("S9")}),
		secondWins("Both Flushes Falls Back to High Card",
			[]poker.Card{c("SK"), c("S2"), c("S8"), c("SQ"), c("S9")},
			[]poker.Card{c("HA"), c("H2"), c("H8"), c("HQ"), c("H9")}),
		firstWins("First Player Full House",
			[]poker.Card{c("C3"), c("H3"), c("S3"), c("CA"), c("HA")},
			[]poker.Card{c("C2"), c("C5"), c("C8"), c("CK"), c("C9")}),
		secondWins("Second Player Full House",
			[]poker.Card{c("C2"), c("C5"), c("C8"), c("CK"), c("C9")},
			[]poker.Card{c("C3"), c("H3"), c("S3"), c("CA"), c("HA")}),
		secondWins("Both Full Houses Uses Three of a Kind",
			[]poker.Card{c("C2"), c("H2"), c("S2"), c("SA"), c("DA")},
			[]poker.Card{c("C3"), c("H3"), c("S3"), c("CA"), c("HA")}),
		firstWins("First Player Four of a Kind",
			[]poker.Card{c("C2"), c("H2"), c("D2"), c("S2"), c("C9")},
			[]poker.Card{c("C3"), c("H3"), c("S3"), c("CA"), c("HA")}),
		secondWins("Second Player Four of a Kind",
			[]poker.Card{c("C3"), c("H3"), c("S3"), c("CA"), c("HA")},
			[]poker.Card{c("C2"), c("H2"), c("D2"), c("S2"), c("C9")}),
		secondWins("Both Four of a Kind Breaks Tie on Value",
			[]poker.Card{c("C2"), c("H2"), c("D2"), c("S2"), c("C9")},
			[]poker.Card{c("C3"), c("H3"), c("D3"), c("S3"), c("C4")}),
		firstWins("First Player Straight Flush",
			[]poker.Card{c("C2"), c("C3"), c("C4"), c("C5"), c("C6")},
			[]poker.Card{c("CA"), c("HA"), c("DA"), c("SA"), c("C9")}),
		secondWins("Second Player Straight Flush",
			[]poker.Card{c("CA"), c("HA"), c("DA"), c("SA"), c("C9")},
			[]poker.Card{c("C2"), c("C3"), c("C4"), c("C5"), c("C6")}),
		secondWins("Both Straight Flush Breaks Tie on High Card",
			[]poker.Card{c("H2"), c("H3"), c("H4"), c("H5"), c("H6")},
			[]poker.Card{c("C7"), c("C3"), c("C4"), c("C5"), c("C6")}),
		tie("Both Straight Flush Tie on Same High Card",
			[]poker.Card{c("H2"), c("H3"), c("H4"), c("H5"), c("H6")},
			[]poker.Card{c("C2"), c("C3"), c("C4"), c("C5"), c("C6")}),
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
