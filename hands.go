package main

import(
	//"fmt"
	dlog "bitbucket.org/thespacecowboys45/dlogger"
	
)

// A hand represents a slice of cards
type Hand []Card

// Adds a card to the current hand, creating a new hand. This doesn't work in
// place for obvious reasons.
func (hand Hand) AddCard(card Card) Hand {
	return append(hand, card)
}

// dxb - see if hand is splittable
func (hand Hand) CanSplit(allowsplits bool) bool {
	dlog.LogEvent("[hands.go][CanSplit()][entry]", "trace")
	dlog.Always("Hand: %v has %d cards. allowsplits is %t", hand, len(hand), allowsplits)
	
	if !allowsplits {
		dlog.Always("Splits not allowed according to rules.")
		return false
	}
	
	// if the hand has more than 2-cards, no splitting
	if len(hand) > 2 {
		dlog.Always("Hand has more than 2 cards.  Cannot split.")
		return false
	}
	
	dlog.Always("Hand values: %d, %d", hand[0].Value, hand[1].Value)
	
	if hand[0].Value == hand[1].Value {
		dlog.Info("Yes, can split")
		return true
	}
	dlog.Info("No, can't split")
	return false
}

// dxb - see if a splittable hand is supposed to be split
// ( based on a split strategy )
func (hand Hand) DoesSplit() bool {
	dlog.Always("[hands.go][CanSplit()][entry]")
	dlog.Always("Hand: %v", hand)
	dlog.Always("Hand values: %d, %d", hand[0].Value, hand[1].Value)
	
	// dxb - potentially we can *only* split if it is the same
	// card.  Only applies for 10-value cards, 10, J, Q, K
	// This function may change also depending on house-rules.
	//
	// For example: Some houses may not allow splitting certain cards (???)
	//  --> tbd
	
	// Do we need to check if the same symbol (for 10-value cards)?
	//dlog.Always("Hand symbols: %d, %d", hand[0].Symbol, hand[1].Symbol)

	// For now, always split like handed cards	
	if hand[0].Value == hand[1].Value {
		dlog.Info("Yes, choose to split")
		return true
	}
	dlog.Info("No, choose NOT to split")
	return false
}

// dxb - see if hand is splittable
func (hand Hand) Split() (Hand, Hand) {
	// HERE WE ARE< finally.  After many coding hours.  This func
	dlog.Always("[hands.go][Split()][entry]")
	dlog.Always("Hand: %v", hand)
	dlog.Always("Hand values: %d, %d", hand[0].Value, hand[1].Value)

	// create two new hands, return both as "the split" and discard the existing Hand
	newHand1 := Hand{}
	newHand2 := Hand{}
	
	// Divy up existing cards
	newHand1 = newHand1.AddCard(hand[0])
	newHand2 = newHand2.AddCard(hand[1])
	
	return newHand1, newHand2
}


// Recursively optimizes the hand for busting. Given the number of alternatives
// allowed to use, determines if we can make a sum with that number of
// alternatives. If it can't, it will try again with ANOTHER number of
// alternatives.
func (hand Hand) sumWithAlternates(alternates int) int {
	accum := 0
	// dxb - not sure what this variable is used for
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
