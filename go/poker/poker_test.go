package poker_test

import (
	"testing"

	"github.com/wlindley/katas/go/poker"
)

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

func TestHand(t *testing.T) {
	tt := []handTestCase{
		validHand("Valid Hand", poker.Club(poker.J), poker.Heart(poker.Jack), poker.Diamond(poker.Queen), poker.Spade(poker.King), poker.Club(poker.Ace)),
		invalidHand("Empty Hand"),
		invalidHand("Too Small Hand", poker.Club(2), poker.Club(3), poker.Club(4), poker.Club(5)),
		invalidHand("Too Large Hand", poker.Club(2), poker.Club(3), poker.Club(4), poker.Club(5), poker.Club(6), poker.Club(7)),
		invalidHand("Too Low Card", poker.Club(1), poker.Club(2), poker.Club(3), poker.Club(4), poker.Club(5)),
		invalidHand("Too High Card", poker.Club(15), poker.Club(2), poker.Club(3), poker.Club(4), poker.Club(5)),
		invalidHand("Duplicate Cards", poker.Heart(2), poker.Heart(2), poker.Club(3), poker.Spade(4), poker.Club(5)),
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

func TestCompare(t *testing.T) {
	tt := []compareTestCase{
		firstWins("First Player High Card", []poker.Card{poker.Club(2), poker.Club(5), poker.Spade(7), poker.Spade(9), poker.Spade(poker.Ace)}, []poker.Card{poker.Heart(poker.King), poker.Heart(9), poker.Diamond(7), poker.Diamond(5), poker.Heart(3)}),
		secondWins("Second Player High Card", []poker.Card{poker.Club(2), poker.Club(5), poker.Spade(7), poker.Spade(9), poker.Spade(poker.King)}, []poker.Card{poker.Heart(poker.Ace), poker.Heart(9), poker.Diamond(7), poker.Diamond(5), poker.Heart(3)}),
		tie("Tie High Card", []poker.Card{poker.Club(2), poker.Club(5), poker.Spade(7), poker.Spade(9), poker.Spade(poker.Ace)}, []poker.Card{poker.Heart(poker.Ace), poker.Heart(9), poker.Diamond(7), poker.Diamond(5), poker.Heart(2)}),
		firstWins("First Player Second Highest Card", []poker.Card{poker.Club(2), poker.Club(5), poker.Spade(7), poker.Spade(poker.Jack), poker.Spade(poker.King)}, []poker.Card{poker.Heart(poker.King), poker.Heart(9), poker.Diamond(7), poker.Diamond(5), poker.Heart(3)}),
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
