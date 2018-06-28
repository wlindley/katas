package poker

import (
	"fmt"
	"sort"
)

// Winner represents which player won the hand
type Winner string

// First means the first player won the hand
const First Winner = "first player won"

// Second means the second player won the hand
const Second Winner = "second player won"

// Tie means the two players tied
const Tie Winner = "tie"

// Card stores the attributes of a single playing card
type Card struct {
	suit  string
	value int
}

func (card Card) compare(other Card) Winner {
	if card.value > other.value {
		return First
	} else if card.value < other.value {
		return Second
	}
	return Tie
}

// Jack is the value of a jack card
const Jack uint = 11

// J is an alias for Jack
const J = Jack

// Queen is the value of a queen card
const Queen uint = 12

// Q is an alias for Queen
const Q = Queen

// King is the value of a king card
const King uint = 13

// K is an alias for King
const K = King

// Ace is the value of an ace card
const Ace uint = 14

// A is an alias for Ace
const A = Ace

// Heart creates a card with the suit of hearts
func Heart(value uint) Card {
	return createCard("H", value)
}

// Diamond creates a card with the suit of diamonds
func Diamond(value uint) Card {
	return createCard("D", value)
}

// Spade creates a card with the suit of spades
func Spade(value uint) Card {
	return createCard("S", value)
}

// Club creates a card with the suit of clubs
func Club(value uint) Card {
	return createCard("C", value)
}

var invalidCard = Card{
	suit:  "INVALID",
	value: 0,
}

func createCard(suit string, value uint) Card {
	if value < minValue || value > maxValue {
		return invalidCard
	}
	return Card{
		suit:  suit,
		value: int(value),
	}
}

// Cards represents the cards a player has been dealt
type Cards []Card

func (cards Cards) Len() int           { return len(cards) }
func (cards Cards) Swap(i, j int)      { cards[i], cards[j] = cards[j], cards[i] }
func (cards Cards) Less(i, j int) bool { return cards[i].value < cards[j].value }

func (cards Cards) any(predicate func(Card) bool) (Card, bool) {
	for _, card := range cards {
		if predicate(card) {
			return card, true
		}
	}
	return invalidCard, false
}

func isInvalid(card Card) bool {
	return card == invalidCard
}

func (cards Cards) containsDuplicates() (Card, bool) {
	for i := 0; i < len(cards)-1; i++ {
		for j := i + 1; j < len(cards); j++ {
			if cards[i] == cards[j] {
				return cards[i], true
			}
		}
	}
	return invalidCard, false
}

// NewHand returns a Cards with the 5 cards making a player's hand
func NewHand(cards ...Card) (Cards, error) {
	if len(cards) != handSize {
		return nil, fmt.Errorf("incorrect number of cards, expected %d", handSize)
	}
	hand := Cards(cards)
	sort.Sort(hand)
	if card, hasInvalid := hand.any(isInvalid); hasInvalid {
		return nil, fmt.Errorf("hand contains invalid card %v", card)
	}
	if card, hasDuplicates := hand.containsDuplicates(); hasDuplicates {
		return nil, fmt.Errorf("hand contains duplicate cards %v", card)
	}
	return hand, nil
}

// Compare compares two poker hands and returns an instance of Winner informing the caller who won the hand
func (cards Cards) Compare(other Cards) Winner {
	for i := handSize - 1; i >= 0; i-- {
		result := cards[i].compare(other[i])
		if result != Tie {
			return result
		}
	}
	return Tie
}

const minValue = 2
const maxValue = 14
const handSize = 5
