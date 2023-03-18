package main

import (
//	"encoding/binary"
	"log"
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
	BETTINGACTION_RESET = iota // 'r'
	BETTINGACTION_INCREASE     // 'i'
	BETTINGACTION_DECREASE     // 'd'
	BETTINGACTION_STAND        // 's'
)

const (
	OUTCOME_ABORT = iota
	OUTCOME_PUSH
	OUTCOME_WIN
	OUTCOME_LOSS
	OUTCOME_INIT // DAVB - added to initialize the wager for 1st bet
)

// DAVB added The betting action a player takes.
type BettingAction int

// The action a player takes.
type Action int

// The result of a game
type Outcome int

type Round struct {
	// The deck we are all playing with.
	deck Deck

	// The dealer's hand
	Dealer Hand

	// The player's hand.
	// DAVB - @TODO make this an array of hands (to handle splits)
	Player Hand
}
/*
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

	// First set of cards...
	round.dealToDealer()
	round.dealToPlayer()

	// Second set of cards...
	round.dealToDealer()
	round.dealToPlayer()

	if verbose {
		log.Printf("Round starts. Dealer: %s, Player: %s", round.Dealer, round.Player)
	}

	// TODO: Add betting in here.

	// If the player has blackjack, he wins!
	if round.Player.Sum() == BUST_LIMIT {
		return OUTCOME_WIN
	}

	for {
		// DAVB - looking into the guts of this function is going to get interesting ...
		action := determineAction(*round)

		if action == ACTION_STAND {
			if verbose {
				log.Println("Player stands.")
			}

			// The user wants to stand so let's see what the dealer
			// does.
			break
		} else if action == ACTION_HIT {
			// Deal a card to the player and go around again.
			round.dealToPlayer()

			if verbose {
				log.Printf("Player hits. Hand: %s Total: %d", round.Player, round.Player.Sum())
			}

			// If the player busts, that's a problem.
			if round.Player.IsBusted() {
				break
			}
		} else if action == ACTION_DOUBLE {
			round.dealToPlayer()

			if verbose {
				log.Printf("Player doubles. Hand: %s Total: %d", round.Player, round.Player.Sum())
			}

			break
		}
	}

	if round.Player.IsBusted() {
		if verbose {
			log.Printf("Player busted!")
		}

		return OUTCOME_LOSS
	}

	// Now for the dealer: While the sum is less than 17, we hit.
	for round.Dealer.Sum() < 17 {
		round.dealToDealer()

		if verbose {
			log.Printf("Dealer hits. Hand: %s Total: %d", round.Dealer, round.Dealer.Sum())
		}
	}

	// Okay, if the dealer busted, you win. If the dealer is greater, you
	// win.
	if round.Dealer.IsBusted() {
		if verbose {
			log.Printf("Dealer busted! Hand: %s", round.Dealer)
		}

		return OUTCOME_WIN
	} else if round.Dealer.Sum() > round.Player.Sum() {
		if verbose {
			log.Printf("Dealer wins. Dealer: %s, Player: %s", round.Dealer, round.Player)
		}

		return OUTCOME_LOSS
	} else if round.Player.Sum() == round.Dealer.Sum() {
		if verbose {
			log.Printf("Round pushes. Dealer: %s, Player: %s", round.Dealer, round.Player)
		}

		return OUTCOME_PUSH
	}

	if verbose {
		log.Printf("Player wins! Dealer: %s, Player: %s", round.Dealer, round.Player)
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
