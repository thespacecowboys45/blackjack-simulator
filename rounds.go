package main

import (
//	"encoding/binary"
//	"log"
	"math/rand"
//	"os"
	"fmt"
//	"time"
	crypto_rand "crypto/rand"
	"encoding/binary"
)

// The minimum number of cards that must be in the deck.
const MINIMUM_SHOE_SIZE = 15

const (
	ACTION_HIT = iota
	ACTION_STAND
	ACTION_DOUBLE
)

// DAVB
const (
	BETTINGACTION_RESET = iota +5 // 'r'
	BETTINGACTION_INCREASE     // 'i'
	BETTINGACTION_DECREASE     // 'd'
	BETTINGACTION_STAND        // 's'
)

func bettingActionToString(action BettingAction) string {
	switch action {
		case BETTINGACTION_RESET:
			return "BETTINGACTION_RESET"
			break
		case BETTINGACTION_INCREASE:
			return "BETTINGACTION_INCREASE"
			break
		case BETTINGACTION_DECREASE:
			return "BETTINGACTION_DECREASE"
			break
		case BETTINGACTION_STAND:
			return "BETTINGACTION_STAND"
			break
		default:
			return fmt.Sprintf("[bettingActionToString] unknown action: %d" , action)
			break
	}
	return ""
}

const (
	OUTCOME_ABORT = iota
	OUTCOME_PUSH
	OUTCOME_WIN
	OUTCOME_WIN_BLACKJACK
	OUTCOME_LOSS
	OUTCOME_INIT // DAVB - added to initialize the wager for 1st bet
)

func outcomeToString(outcome int) string {
	switch (outcome) {
		case OUTCOME_ABORT:
			return "OUTCOME_ABORT"
			break
		case OUTCOME_PUSH:
			return "OUTCOME_PUSH"
			break
		case OUTCOME_WIN:
			return "OUTCOME_WIN"
			break
		case OUTCOME_WIN_BLACKJACK:
			return "OUTCOME_WIN_BLACKJACK"
			break
		case OUTCOME_LOSS:
			return "OUTCOME_LOSS"
			break
		case OUTCOME_INIT:
			return "OUTCOME_INIT"
			break
		default:
			return fmt.Sprintf("unknown outcome: %d", outcome)
			break
	}
	return ""
}

// DAVB added The betting action a player takes.
type BettingAction int

// The action a player takes.
type Action int

// The result of a game
type Outcome int

/**
 * dxb - future thought, add structure for Player
 *
 * to add multiple-players variable into the mix
 *
 * type Player struct {
 *     seatNumber int  // a thought to track the order of the player
 *     Hands []Hand
 * }
 */

type Round struct {
	// The deck we are all playing with.
	deck Deck

	// The dealer's hand
	Dealer Hand

	// The player's hand.
	// DAVB - @TODO make this an array of hands (to handle splits)
	Player Hand
	
	// Implement multiple hands (possible) for a player
	// @TODO - splits1
	// PlayersHand []Hand
}
/*
do not know why this is stubbed out or even in here really

maybe as an example of how to use the init() function for a module ???


func init() {
	fmt.Printf("rounds.go [init][entry]\n")
    var b [8]byte
    _, err := crypto_rand.Read(b[:])
    if err != nil {
        panic("cannot seed math/rand package with cryptographically secure random number generator")
    }
    math_rand.Seed(int64(binary.LittleEndian.Uint64(b[:])))
}
*/

func (round *Round) dealToDealer() {
	// Create the initial hand...
	var tmpCard Card

	// Get the dealer's card first...
	tmpCard, round.deck = round.deck.Draw()
	round.Dealer = round.Dealer.AddCard(tmpCard)
}

func (round *Round) dealToPlayer() {
	// Create the initial hand...
	var tmpCard Card

	// Get the dealer's card first...
	tmpCard, round.deck = round.deck.Draw()
	round.Player = round.Player.AddCard(tmpCard)
	
	// @TODO - splits1
	// add logic to handle multiple hands
	//
	// 2nd option: put this functionality below and
	// pass a parameter to this function which is 
	// *which hand* to deal for the player
	// - only issue is going to be that we do not want
	// to alternate between the two hands - we need to
	// play out one new hand entirely, and then go to
	// the next hand
	//
	// for i = 0; i < len(round.PlayersHand); i++
	//
	//    round.PlayersHand[i] = round.PlayersHand[i].AddCard(tmpCard)
}

func (round *Round) Play(determineAction func(round Round) Action) Outcome {
	// If there are less than (some number) cards in the deck, we'll abort
	// this round.
	if len(round.deck) < MINIMUM_SHOE_SIZE {
		return OUTCOME_ABORT
	}

	// Clear our both hands!
	round.Dealer = Hand{}
	round.Player = Hand{}
	
	// @TODO - split1
	// add:
	// round.PlayersHand = []
	//
	// or something
	

	// First set of cards...
	round.dealToDealer()
	round.dealToPlayer()

	// Second set of cards...
	round.dealToDealer()
	round.dealToPlayer()

	if verbose {
		fmt.Printf("[rounds.go] Round starts. Dealer: %s, Player: %s", round.Dealer, round.Player)
	}

	// TODO: Add betting in here.

	// If the player has blackjack, he wins!
	if round.Player.Sum() == BUST_LIMIT {
		//
		// @TODO - add OUTCOME_BLACKJACK as possibility
		//
		// return OUTCOME_WIN_BLACKJACK
		//
		return OUTCOME_WIN
	}
	
	// @TODO - splits1
	// it is possible to use this as insertion point - perhaps
	// 1) determine if player SPLITS - this is initial decision, just like doubledown
	//
	// 2) loop through each hand to play it out.  Variant comes to play if house
	// allows more than one split opportunity for the player after the first one.
	//    Maybe that is the option here - deal with all splits, then the player
	//    has all the hands they will ever have to begin with.
	//
	//    Exception to this is: the order that the cards come out - may allow first
	//    split hand to be played out, then we move on to the second one.  The first
	//    card for the second split hand is Ace.  So, again, we have to decide if we
	//    split or not.
	
	// test for split opportunity
	if round.Player[0] == round.Player[1]  {
		fmt.Printf("[rounds.go][play()] There is a SPLIT OPPORTUNITY.  We drew the same value cards: %d!", round.Player[0])
	}
	

	// dxb - we are in a for loop, always HITTING until we either stand, double, or bust
	for {
		// DAVB - looking into the guts of this function is going to get interesting ...
		// So, in the betting module, instead of this (passed in function call), then
		// call it here.  I get it - this was done this way so that the 
		// strategy employed can be passed in as a function, one which was
		// determined previously, outside this function call
		// 
		// and, therefore, is available to this code block
		//
		action := determineAction(*round)

		if action == ACTION_STAND {
			if verbose {
				fmt.Println("[rounds.go] Player stands.")
			}

			// The user wants to stand so let's see what the dealer
			// does.
			break
		} else if action == ACTION_HIT {
			// Deal a card to the player and go around again.
			round.dealToPlayer()
			
			// @TODO - splits1
			// put some logic in here - how do we deal with the players second (or third) hands?
			//
			// for i=0; i< len(round.PlayersHand); i++ {
			//     round.PlayersHand[i].dealToPlayer()
			

			if verbose {
				fmt.Printf("[rounds.go] Player hits. Hand: %s Total: %d", round.Player, round.Player.Sum())
			}

			// If the player busts, that's a problem.
			if round.Player.IsBusted() {
				break
			}
		} else if action == ACTION_DOUBLE {
			round.dealToPlayer()
			
			// @TODO - doubledown1
			//
			// We need to impact / affect the current wager - is it available here ?
			// Perhaps always return, from this function, the players

			if verbose {
				fmt.Printf("[rounds.go] Player doubles. Hand: %s Total: %d", round.Player, round.Player.Sum())
			}

			break
		}
		// @TODO - splits1
		//
		// add new case
		// else if action == ACTION_SPLIT {
		//
		// -----> need new function to take the cards and create two hands from one
		// and then deal two more cards on top of the first two, and
		// THEN
		// re-evaluate the above code!
		//
		//
		//
	}

	// @TODO - split1
	// How do we handle different outcomes for multiple hands?
	//
	// Do we *need* to check this here, or is it simply a short-circuit?
	// *CAN* we do this after the dealer - or is it a short circuit because
	// we *do not want* the dealer to continue if we busted out completely.
	if round.Player.IsBusted() {
		if verbose {
			fmt.Printf("[rounds.go] Player busted!")
		}

		return OUTCOME_LOSS
	}

	// @TODO - put in boolean to check if dealer hits soft-17,
	//  IF SO
	//  THEN
	// this becaomse:
	//  for round.Dealer.Sum() < 18
	//
	
	// Now for the dealer: While the sum is less than 17, we hit.
	for round.Dealer.Sum() < 17 {
		round.dealToDealer()

		if verbose {
			fmt.Printf("[rounds.go] Dealer hits. Hand: %s Total: %d", round.Dealer, round.Dealer.Sum())
		}
	}

	// Okay, if the dealer busted, you win. If the dealer is greater, you
	// win.
	if round.Dealer.IsBusted() {
		if verbose {
			fmt.Printf("[rounds.go] Dealer busted! Hand: %s", round.Dealer)
		}

		return OUTCOME_WIN
	} else if round.Dealer.Sum() > round.Player.Sum() {
		if verbose {
			fmt.Printf("[rounds.go] Dealer wins. Dealer: %s, Player: %s", round.Dealer, round.Player)
		}

		return OUTCOME_LOSS
	} else if round.Player.Sum() == round.Dealer.Sum() {
		if verbose {
			fmt.Printf("[rounds.go] Round pushes. Dealer: %s, Player: %s", round.Dealer, round.Player)
		}

		return OUTCOME_PUSH
	}

	if verbose {
		fmt.Printf("[rounds.go] Player wins! Dealer: %s, Player: %s", round.Dealer, round.Player)
	}

	return OUTCOME_WIN
}

func seedRand() {

    var b [8]byte
    _, err := crypto_rand.Read(b[:])
    if err != nil {
        panic("cannot seed math/rand package with cryptographically secure random number generator")
    }
    theSeed := int64(binary.LittleEndian.Uint64(b[:]))
	fmt.Printf("theSeed: %v\n", theSeed)
	rand.Seed(theSeed)

/*
 don't use time to seed 
 
	now := time.Now().UnixNano()
	fmt.Printf("NOW: %v\n", now)
	//rand.Seed(seed)
	rand.Seed(now)
*/	
}

func NewRound(deck Deck) *Round {
	round := new(Round)
	round.deck = deck
	return round
}
