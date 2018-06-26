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
			_, err := poker.Hand(x.Hand...)
			if err == nil && !x.IsValid {
				t.Errorf("%v should have failed to create a hand, but did not", x.Hand)
			} else if err != nil && x.IsValid {
				t.Errorf("%v should have successfully created a hand, but instead errored: %s", x.Hand, err)
			}
		})
	}
}
