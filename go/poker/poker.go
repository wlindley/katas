package poker

import (
	"fmt"
	"sort"
)

const minValue = 2
const maxValue = 14
const handSize = 5

// Winner represents which player won the hand
type Winner int

// First, Second, and Tie denote which player's hand was higher value, none means there's no valid outcome
const (
	First  Winner = -1
	Tie    Winner = 0
	Second Winner = 1
	none   Winner = 100
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

func (w Winner) conclusive() bool {
	return w == First || w == Second
}

// Card stores the attributes of a single playing card
type Card struct {
	suit  string
	value int
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

var invalidCard = Card{}

func createCard(suit string, value uint) Card {
	if value < minValue || value > maxValue {
		return invalidCard
	}
	return Card{
		suit:  suit,
		value: int(value),
	}
}

type cardCollection []Card

func (cards cardCollection) any(predicate func(Card) bool) (Card, bool) {
	for _, card := range cards {
		if predicate(card) {
			return card, true
		}
	}
	return invalidCard, false
}

func (cards cardCollection) all(predicate func(Card) bool) bool {
	for _, card := range cards {
		if !predicate(card) {
			return false
		}
	}
	return true
}

func isInvalid(card Card) bool {
	return card == invalidCard
}

func valueEquals(value int) func(Card) bool {
	return func(card Card) bool {
		return card.value == value
	}
}

func suitEquals(suit string) func(Card) bool {
	return func(card Card) bool {
		return card.suit == suit
	}
}

// Hand represents the cards a player has been dealt
type Hand struct {
	cards [handSize]Card
}

// NewHand returns a Hand with the 5 cards making a player's hand
func NewHand(cards ...Card) (*Hand, error) {
	if len(cards) != handSize {
		return nil, fmt.Errorf("incorrect number of cards, expected %d", handSize)
	}
	hand := Hand{}
	copy(hand.cards[:], cards)
	sort.Sort(&hand)
	if card, hasInvalid := cardCollection(hand.cards[:]).any(isInvalid); hasInvalid {
		return nil, fmt.Errorf("hand contains invalid card %v", card)
	}
	if card, hasDuplicates := hand.containsDuplicates(); hasDuplicates {
		return nil, fmt.Errorf("hand contains duplicate cards %v", card)
	}
	return &hand, nil
}

func (h *Hand) Len() int           { return len(h.cards) }
func (h *Hand) Swap(i, j int)      { h.cards[i], h.cards[j] = h.cards[j], h.cards[i] }
func (h *Hand) Less(i, j int) bool { return h.cards[i].value > h.cards[j].value } // Descending order by value

func (h *Hand) highCard() Card {
	return h.cards[0]
}

func (h *Hand) containsDuplicates() (Card, bool) {
	for i := 0; i < len(h.cards)-1; i++ {
		if h.cards[i] == h.cards[i+1] {
			return h.cards[i], true
		}
	}
	return invalidCard, false
}

func (h *Hand) valueOfSetSize(size int) int {
	return h.valueOfSetSizeIgnore(size, 0)
}

func (h *Hand) valueOfSetSizes(firstSize, secondSize int) (int, int) {
	firstValue := h.valueOfSetSizeIgnore(firstSize, 0)
	secondValue := h.valueOfSetSizeIgnore(secondSize, firstValue)
	return firstValue, secondValue
}

func (h *Hand) valueOfSetSizeIgnore(size, ignore int) int {
	for i := 0; i <= len(h.cards)-size; i++ {
		currentValue := h.cards[i].value
		if currentValue == ignore {
			continue
		}
		if cardCollection(h.cards[i : i+size]).all(valueEquals(currentValue)) {
			return currentValue
		}
	}
	return 0
}

func (h *Hand) valueOfStraight() int {
	for i := 0; i < len(h.cards)-1; i++ {
		if h.cards[i].value-1 != h.cards[i+1].value {
			return 0
		}
	}
	return h.highCard().value
}

func (h *Hand) valueOfFlush() int {
	cards := cardCollection(h.cards[:])
	if cards.all(suitEquals(h.highCard().suit)) {
		return h.highCard().value
	}
	return 0
}

var comparators = []func(*Hand, *Hand) Winner{
	compareStraightFlushes,
	compareFourOfAKind,
	compareFullHouses,
	compareFlushes,
	compareStraights,
	compareThreeOfAKind,
	compareTwoPairs,
	comparePairs,
	compareHighCards,
}

// Compare compares two poker hands and returns an instance of Winner informing the caller who won the hand
func (h *Hand) Compare(other *Hand) Winner {
	for _, comparator := range comparators {
		result := comparator(h, other)
		if result != Tie && result != none {
			return result
		}
	}
	return Tie
}

func compareStraightFlushes(first, second *Hand) Winner {
	firstFlush := first.valueOfFlush()
	firstStraight := first.valueOfStraight()
	secondFlush := second.valueOfFlush()
	secondStraight := second.valueOfStraight()
	result := compareMultiSetValidity(firstFlush, firstStraight, secondFlush, secondStraight)
	if result == Tie {
		result = compareValues(first.highCard().value, second.highCard().value)
	}
	return result
}

func compareFourOfAKind(first, second *Hand) Winner {
	return compareSetOfSize(first, second, 4)
}

func compareFullHouses(first, second *Hand) Winner {
	firstThree, firstTwo := first.valueOfSetSizes(3, 2)
	secondThree, secondTwo := second.valueOfSetSizes(3, 2)
	result := compareMultiSetValidity(firstThree, firstTwo, secondThree, secondTwo)
	if result == Tie {
		result = compareValues(firstThree, secondThree)
	}
	return result
}

func compareFlushes(first, second *Hand) Winner {
	firstFlush := first.valueOfFlush()
	secondFlush := second.valueOfFlush()
	return compareValues(firstFlush, secondFlush)
}

func compareStraights(first, second *Hand) Winner {
	firstStraight := first.valueOfStraight()
	secondStraight := second.valueOfStraight()
	return compareValues(firstStraight, secondStraight)
}

func compareThreeOfAKind(first, second *Hand) Winner {
	return compareSetOfSize(first, second, 3)
}

func compareTwoPairs(first, second *Hand) Winner {
	firstHigh, firstLow := first.valueOfSetSizes(2, 2)
	secondHigh, secondLow := second.valueOfSetSizes(2, 2)
	result := compareMultiSetValidity(firstHigh, firstLow, secondHigh, secondLow)
	if !result.conclusive() {
		result = compareValues(firstHigh, secondHigh)
	}
	if !result.conclusive() {
		result = compareValues(firstLow, secondLow)
	}
	return result
}

func comparePairs(first, second *Hand) Winner {
	return compareSetOfSize(first, second, 2)
}

func compareHighCards(first, second *Hand) Winner {
	for i := 0; i < handSize; i++ {
		result := compareValues(first.cards[i].value, second.cards[i].value)
		if result.conclusive() {
			return result
		}
	}
	return Tie
}

func compareValues(first, second int) Winner {
	if first == 0 && second == 0 {
		return none
	}
	if first > second {
		return First
	} else if first < second {
		return Second
	}
	return Tie
}

func compareSetOfSize(first, second *Hand, size int) Winner {
	firstSet := first.valueOfSetSize(size)
	secondSet := second.valueOfSetSize(size)
	return compareValues(firstSet, secondSet)
}

func compareMultiSetValidity(firstOne, firstTwo, secondOne, secondTwo int) Winner {
	firstValid := allFound(firstOne, firstTwo)
	secondValid := allFound(secondOne, secondTwo)
	switch {
	case firstValid && secondValid:
		return Tie
	case firstValid && !secondValid:
		return First
	case !firstValid && secondValid:
		return Second
	}
	return none
}

func allFound(values ...int) bool {
	for _, value := range values {
		if value == 0 {
			return false
		}
	}
	return true
}
