package main

import (
	"bytes"
	"math/rand"
	"log"
)

// The number of decks to play with.
const DEFAULT_DECKS = 6

// Well, obiously!
const BUST_LIMIT = 21

// Represents a deck of cards.
type Deck []Card

// Shuffles the deck, creating a new deck.
func (deck Deck) Shuffle() Deck {
	// Just to make sure we're sufficiently random.
	seedRand()

	// get a range back of number in 'random' order to determine how to
	// re-order the cards in the deck
	perm := rand.Perm(len(deck))
	newDeck := make(Deck, len(deck))

	for j, i := range perm {
		newDeck[i] = deck[j]
	}

	return newDeck
}

func (deck Deck) String() string {
	var buf bytes.Buffer

	for _, card := range deck {
		buf.WriteString(card.String())
		buf.WriteString(" ")
	}

	return buf.String()
}

// Draws a card from the deck and removes it from the deck.
func (deck Deck) Draw() (Card, Deck) {
	log.Printf("[decks.go][Draw()][length of deck: %d]", len(deck))
	return deck[0], deck[1:len(deck)]
}

// DAVB - generates a suit of cards in ascending order
func generateSuit(suit rune, deck Deck) Deck {
	for i := 1; i < 14; i += 1 {
		deck = append(deck, NewCard(i, suit))
	}

	return deck
}

// Returns a new set of cards. Whew.
func NewDeck() Deck {
	deck := Deck{}
	deck = generateSuit(SUIT_HEARTS, deck)
	deck = generateSuit(SUIT_DIAMONDS, deck)
	deck = generateSuit(SUIT_CLUBS, deck)
	deck = generateSuit(SUIT_SPADES, deck)
	return deck
}

// Returns a set of decks all as a single deck.
func NewMultipleDeck(decks int) Deck {
	deck := Deck{}

	// DAVB - creates a deck of decks
	for i := 0; i < decks; i++ {
		deck = append(deck, NewDeck()...)
	}

	//
	// In the end, we end up with an array of decks and inside each
	// deck are four suits (heart,diamond,club,spade) 
	// comprised of fourteen (14) cards in ascending
	// order numerically, starting at 2 ... J, Q, K, A
	//
	return deck
}
