package main

import(
	//"fmt"
)

// Represents a set of cards...obviously...
type Hand []Card

// Adds a card to the current hand, creating a new hand. This doesn't work in
// place for obvious reasons.
func (hand Hand) AddCard(card Card) Hand {
	return append(hand, card)
}

// Recursively optimizes the hand for busting. Given the number of alternatives
// allowed to use, determins if we can make a sum with that number of
// alternatives. If it can't, it will try again with ANOTHER number of
// alternatives.
func (hand Hand) sumWithAlternates(alternates int) int {
	accum := 0
	// DAVB - not sure what this variable is used for
	alternatesUsed := 0

	for _, card := range hand {
		// If first iteration and the card we examine has two unique total values assigned
		if alternatesUsed < alternates && card.HasUsefulAlternate() {
			// We used 1 alternate value and the total points is that alternate value (11)
			alternatesUsed += 1
			accum += card.AlternateValue
		} else {
			accum += card.Value
		}
	}

	// If we're still busted and the alternates is less than the number of
	// cards in the hand, we should try a different approach. Otherwise,
	// there's nothing we can do.
	if accum > BUST_LIMIT && alternates < len(hand) {
		return hand.sumWithAlternates(alternates + 1)
	}

	return accum
}

// Get the current total of the hand.
func (hand Hand) Sum() int {
	return hand.sumWithAlternates(0)
}

// Returns true if the hand is busted, false otherwise.
func (hand Hand) IsBusted() bool {
	return hand.Sum() > BUST_LIMIT
}

// If the hand has an ace that is counting as it's 11 value, it's considered a
// soft hand. Different strategies are applied in that scenario.
func (hand Hand) IsSoft() bool {
	aces := 0
	otherSum := 0

	// Let's see if the hand actually *has* an ace anyway.
	for _, card := range hand {
		if card.Symbol == CARD_ACE {
			aces += 1
		} else {
			otherSum += card.Value
		}
	}

	// No ace, so this hand can't be soft.
	if aces < 1 {
		return false
	}
	
	// DAVB - Player has an ACE at this point in code

	// If any number of aces can be added in at their primary value then the hand
	// is indeed soft!
	singles := (aces - 1)

	//
	// ^^^^ Is this right?
	// DAVB - Count the hand total if we consider the ACE to have value of 1
	// helps determine if we have "wiggle room" to take a hit.
	// Rather - it determines which "strategy" to use in the strategies file: soft or hard
	//
	/*
	@TODO - figure out later
	t := BUST_LIMIT-(otherSum+singles)
	msg := fmt.Sprintf("-------------------> t: %d , otherSum: %d, singles: %d\n", t, otherSum, singles)
	dlog(msg)
	*/
	
	
	return BUST_LIMIT-(otherSum+singles) >= 11
}

func (hand Hand) IsHard() bool {
	return !hand.IsSoft()
}
