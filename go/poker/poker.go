package poker

import (
	"fmt"
	"sort"
)

// Winner represents which player won the hand
type Winner int

// First, Second, and Tie denote which player's hand was higher value
const (
	First  Winner = -1
	Tie    Winner = 0
	Second Winner = 1
)

func (w Winner) String() string {
	switch w {
	case First:
		return "First player won"
	case Second:
		return "Second player won"
	}
	return "Tie"
}

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

// Constants for the value of face cards
const (
	Jack, J  uint = 11, 11
	Queen, Q uint = 12, 12
	King, K  uint = 13, 13
	Ace, A   uint = 14, 14
)

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

// Hand represents the cards a player has been dealt
type Hand struct {
	cards [handSize]Card
}

func (h *Hand) Len() int           { return len(h.cards) }
func (h *Hand) Swap(i, j int)      { h.cards[i], h.cards[j] = h.cards[j], h.cards[i] }
func (h *Hand) Less(i, j int) bool { return h.cards[i].value < h.cards[j].value }

func (h *Hand) any(predicate func(Card) bool) (Card, bool) {
	for _, card := range h.cards {
		if predicate(card) {
			return card, true
		}
	}
	return invalidCard, false
}

func isInvalid(card Card) bool {
	return card == invalidCard
}

func (h *Hand) containsDuplicates() (Card, bool) {
	for i := 0; i < len(h.cards)-1; i++ {
		for j := i + 1; j < len(h.cards); j++ {
			if h.cards[i] == h.cards[j] {
				return h.cards[i], true
			}
		}
	}
	return invalidCard, false
}

// Compare compares two poker hands and returns an instance of Winner informing the caller who won the hand
func (h *Hand) Compare(other *Hand) Winner {
	for i := handSize - 1; i >= 0; i-- {
		result := h.cards[i].compare(other.cards[i])
		if result != Tie {
			return result
		}
	}
	return Tie
}

// NewHand returns a Hand with the 5 cards making a player's hand
func NewHand(cards ...Card) (*Hand, error) {
	if len(cards) != handSize {
		return nil, fmt.Errorf("incorrect number of cards, expected %d", handSize)
	}
	hand := Hand{}
	copy(hand.cards[:], cards)
	sort.Sort(&hand)
	if card, hasInvalid := hand.any(isInvalid); hasInvalid {
		return nil, fmt.Errorf("hand contains invalid card %v", card)
	}
	if card, hasDuplicates := hand.containsDuplicates(); hasDuplicates {
		return nil, fmt.Errorf("hand contains duplicate cards %v", card)
	}
	return &hand, nil
}

const minValue = 2
const maxValue = 14
const handSize = 5
