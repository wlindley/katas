package poker

import (
	"fmt"
)

// Winner represents which player won the hand
type Winner int

// First means the first player won the hand
const First Winner = 1

// Second means the second player won the hand
const Second Winner = -1

// Tie means the two players tied
const Tie Winner = 0

// Card stores the attributes of a single playing card
type Card struct {
	suit  string
	value int
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

// Hand returns a Cards with the 5 cards making a player's hand
func Hand(cards ...Card) (Cards, error) {
	if len(cards) != handSize {
		return nil, fmt.Errorf("incorrect number of cards, expected %d", handSize)
	}
	hand := Cards(cards)
	if card, hasInvalid := hand.any(isInvalid); hasInvalid {
		return nil, fmt.Errorf("hand contains invalid card %v", card)
	}
	if card, hasDuplicates := hand.containsDuplicates(); hasDuplicates {
		return nil, fmt.Errorf("hand contains duplicate cards %v", card)
	}
	return nil, nil
}

const minValue = 2
const maxValue = 14
const handSize = 5
