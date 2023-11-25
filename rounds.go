package main

import (
//	"encoding/binary"
//	"log"
	"math/rand"
//	"os"
	"fmt"
	"log"
//	"time"
	crypto_rand "crypto/rand"
	"encoding/binary"
	dlog "bitbucket.org/thespacecowboys45/dlogger"
)

// The minimum number of cards that must be in the deck.
const MINIMUM_SHOE_SIZE = 30

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

func outcomeToString(outcome Outcome) string {
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
	Player Hand

	// Implement multiple players sitting for a round	
	num_players int
	Players []Hand
	Outcomes []Outcome
	
	// @TODO - splits1
	// Implement multiple hands (possible) for a player
	// probably make this a player object (NOT A PLAYA object!)

	// implement Players as objects
	PlayersObj []Player
	
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

func (round *Round) dealToMultiPlayer(player_num int) {
	// Create the initial hand...
	var tmpCard Card

	// Get the dealer's card first...
	tmpCard, round.deck = round.deck.Draw()
	
	// swap out for multi-player
	//round.Player = round.Player.AddCard(tmpCard)
	round.Players[player_num] = round.Players[player_num].AddCard(tmpCard)
	
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


// creating new function so I don't pollute the original working function
func (round *Round) PlayMultiPlayer(determineAction func(round Round, player_number int) Action) ([]Outcome, int) {
	log.Printf("[rounds.go][PlayMultiPlayer][entry]")
	
	// Total number of hands played this round
	total_hands_played_this_round := 0
	dlog.Debug("[rounds.go][PlayMultiPlayer][initialize total_hands_played_this_round to %d]", total_hands_played_this_round)
	
	// If there are less than (some number) cards in the deck, we'll abort
	// this round.fvgbdxsefr
	if len(round.deck) < MINIMUM_SHOE_SIZE {
		// @TODO - refactor for multiplayer
		// hack attack
		for i:=0; i<round.num_players; i++ {
			round.Outcomes[i] = OUTCOME_ABORT
		}
		
		//return OUTCOME_ABORT, aand that we played 0 hands (no one gets to play this round)
		return round.Outcomes, total_hands_played_this_round
	}

	// Clear out all hands!
	//round.Player = Hand{} // refactor for multiplayer
	for i:=0; i < round.num_players; i++ {
		log.Printf("[rounds.go][PlayMultiPlayer()][ process player, initialize empty hand: %d]", i)
		round.Players[i] = Hand{}
	}

	round.Dealer = Hand{}

	
	/////////
	// LET'S PLAY!
	/////////
	
	/**
	 * Oct '23
	 * implement multi-player
	 */

	// First card for players...
	for i:=0; i < round.num_players; i++ {
		log.Printf("[rounds.go][PlayMultiPlayer()][ process player first card: %d]", i)
		round.dealToMultiPlayer(i)
	}

	// First card to dealer...
	log.Printf("[rounds.go][PlayMultiPlayer()][ deal first card to dealer ]")
	round.dealToDealer()
	
	// Second card for players...
	// Everyone receives a second card, and then we evaluate the players action choice.
	// We do not evaluate (splits) as the second card is dealt to each player.
	// In other words - player 2 cannot split her hand before player 3 is dealt a second
	// initial card.
	for i:=0; i < round.num_players; i++ {
		log.Printf("[rounds.go][PlayMultiPlayer()][ process player second card: %d]", i)
		round.dealToMultiPlayer(i)
	}

	// Second card to dealer...
	log.Printf("[rounds.go][PlayMultiPlayer()][ deal second card to dealer ]")
	round.dealToDealer()

	if verbose {
		log.Printf("[rounds.go][PlayMultiPlayer()] Round starts. Dealer: %s", round.Dealer)
		for i:=0; i < round.num_players; i++ {
			log.Printf("[rounds.go][PlayMultiPlayer()][ PLAYER %d starts with: %s]", i, round.Players[i])
		}
	}
	
	/**
	 * Main player loop - deal for all players 
	 */
	for i:=0; i< round.num_players; i++ {
		// Play
		// TODO: Add betting in here.
		log.Printf("[rounds.go][player loop][player #%d playing][Current hand total: %d", i, round.Players[i].Sum())
		
		// Assume we played another hand for this player.  Count it.
		total_hands_played_this_round++		
		dlog.Debug("[rounds.go][PlayMultiPlayer][incremented total_hands_played_this_round == %d]", total_hands_played_this_round)
		
	
		// If the player has blackjack, he wins!
		if round.Players[i].Sum() == BUST_LIMIT {
			//
			// @TODO - add OUTCOME_BLACKJACK as possibility
			//
			// return OUTCOME_WIN_BLACKJACK
			//
			
			
			//return OUTCOME_WIN
			log.Printf("[rounds.go][player loop][ Player %d got blackjack! ]", i)
			round.Outcomes[i] = OUTCOME_WIN
			continue
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
		if round.Players[i][0] == round.Players[i][1]  {
			log.Printf("[rounds.go][PlayMultiPlayer()][player #%d] There is a SPLIT OPPORTUNITY.  We drew the same value cards: %d!", i, round.Players[i][0])
			
			
			// Add a new hand for this player
			round.PlayersObj[i].Hands = round.PlayersObj[i].AddHand(Hand{}) 
			round.PlayersObj[i].activeHand = 0
			
			dlog.Always("[rounds.go][PlayMultiPlayer()][player #%d SPLIT this hand and now has %d hands, %d h-cap]",
				i, len(round.PlayersObj[i].Hands), cap(round.PlayersObj[i].Hands))			
		}
		
	
		// dxb - we are in a for loop, always HITTING until we either stand, double, or bust
		for {
			// dxb - looking into the guts of this function is going to get interesting ...
			// So, in the betting module, instead of this (passed in function call), then
			// call it here.  I get it - this was done this way so that the 
			// strategy employed can be passed in as a function, one which was
			// determined previously, outside this function call
			// 
			// and, therefore, is available to this code block
			//
			
			// determineAction is actually looking at a players hand 
			// to compare against the dealer
			
			
			
			//action := determineAction(*round)
			action := determineAction(*round, i)


			// dxb - had to add additional check here.  With multiple players we can
			// run out of cards in middle of round and that is no es bueno
			// ...			
			// If there are less than (some number) cards in the deck, we'll abort
			// this round.
			if len(round.deck) < MINIMUM_SHOE_SIZE {
				// @TODO - refactor for multiplayer
				// hack attack
				for i:=0; i<round.num_players; i++ {
					round.Outcomes[i] = OUTCOME_ABORT
				}
				//return OUTCOME_ABORT, aand that we played 0 hands (no one gets to play this round)
				dlog.Always("[rounds.go][PlayMultiPlayer()][ran out of cards.  deck length==%d, min=%d]", len(round.deck), MINIMUM_SHOE_SIZE)
				return round.Outcomes, 0
			}	
			
			if action == ACTION_STAND {
				if verbose {
					log.Printf("[rounds.go][PlayMultiPlayer()][player #%d] Player stands.", i)
				}
	
				// The user wants to stand so let's see what the dealer does.
				break
			} else if action == ACTION_HIT {
				// Deal a card to the player and go around again.
				//round.dealToPlayer()
				if verbose {
					dlog.Debug("[rounds.go][PlayMultiPlayer()][player #%d] Player hits. Hand: %s Total: %d", i, round.Players[i], round.Players[i].Sum())
				}
				
				round.dealToMultiPlayer(i)

				// @TODO - splits1
				// put some logic in here - how do we deal with the players second (or third) hands?
				//
				// for i=0; i< len(round.PlayersHand); i++ {
				//     round.PlayersHand[i].dealToPlayer()
				
	
				// If the player busts, that's a problem.
				//if round.Player.IsBusted() {
				if round.Players[i].IsBusted() {
					break
				}
			} else if action == ACTION_DOUBLE {
				if verbose {
					dlog.Debug("[rounds.go][PlayMultiPlayer()][player #%d] Player doubles. Hand: %s Total: %d", i, round.Player, round.Player.Sum())
				}

				//round.dealToPlayer()
				round.dealToMultiPlayer(i)

				
				// @TODO - doubledown1
				//
				// We need to impact / affect the current wager - is it available here ?
				// Perhaps always return, from this function, the players
	
				// player doubled and gets no more cards
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
		//if round.Player.IsBusted() {
		if round.Players[i].IsBusted() {
			if verbose {
				log.Printf("[rounds.go][PlayMultiPlayer()][player #%d] Player busted!", i)
			}
	
			//return OUTCOME_LOSS
			round.Outcomes[i] = OUTCOME_LOSS
		}
	} // end for-loop of players

	
	/**
	 * -----------------------
	 * If any players are still in lets move on to the dealer
	 * -----------------------
	 */
	everyoneBusted := true
	for j:=0; j<num_players; j++ {
		if !round.Players[j].IsBusted() {
			log.Printf("[rounds.go][PlayerMultiPlayer()][player #%d][everyoneBusted?][ player %d did not bust this round ]", j, j)
			everyoneBusted = false
			break
		} else {
			log.Printf("[rounds.go][PlayerMultiPlayer()][player #%d][everyoneBusted?][ player %d busted this round ]", j, j)
		}
	}
	
	// short-circuit and do not deal to the dealer
	if everyoneBusted {
		log.Printf("[rounds.go][PlayerMultiPlayer()][everyone busted.  skip dealer.]")
		return round.Outcomes, total_hands_played_this_round
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
			log.Printf("[rounds.go][PlayMultiPlayer()] Dealer hits. Hand: %s Total: %d", round.Dealer, round.Dealer.Sum())
		}
	}

	/**
	 * dealer is done drawing cards
	 *
	 * check player results
	 */
	for i:=0; i<num_players; i++ {
		// Okay, if the dealer busted, you win. 
		// If the dealer is lesser, you win.
		if round.Dealer.IsBusted() {
			if verbose {
				log.Printf("[rounds.go][PlayMultiPlayer()] Dealer busted! Hand: %s", round.Dealer)
			}
	
			log.Printf("[rounds.go][PlayMultiPlayer][player #%d][Check if player has busted already, hand total: %d]", i, round.Players[i].Sum())
			if round.Players[i].IsBusted() {
				// this is redundant and should have been set already, but be explicit
				log.Printf("[rounds.go][PlayMultiPlayer][player #%d][Player already busted.  Sorry!]", i)
				round.Outcomes[i] = OUTCOME_LOSS
				continue
				
			} else {
				log.Printf("[rounds.go][PlayMultiPlayer][player #%d][Player did not bust.  It is a win.]", i)
				round.Outcomes[i] = OUTCOME_WIN
				continue
			}
		} else if round.Dealer.Sum() > round.Players[i].Sum() {
			if verbose {
				log.Printf("[rounds.go][PlayMultiPlayer()][player %d] Dealer wins. Dealer: %s, Player: %s", i, round.Dealer, round.Players[i])
			}
			round.Outcomes[i] = OUTCOME_LOSS
			continue
			
		} else if round.Players[i].Sum() == round.Dealer.Sum() {
			if verbose {
				log.Printf("[rounds.go][PlayMultiPlayer()] Round pushes. Dealer: %s, Player: %s", round.Dealer, round.Players[i])
			}
			round.Outcomes[i] = OUTCOME_PUSH
			continue
		}
	
		// We get here in the case the player is still in the game and beat the dealer's total
		if verbose {
			log.Printf("[rounds.go][PlayMultiPlayer()] Player %d wins! Dealer: %s, Player: %s", i, round.Dealer, round.Players[i])
			log.Printf(" -------------- HOW DID WE GET HERE ??? -----------------------------")
		}
		round.Outcomes[i] = OUTCOME_WIN
	} // end for loop checking players card totals against the dealer

	for i:=0; i<len(round.Outcomes); i++ {
		log.Printf("[rounds.go][PlayerMultiPlayer()][outcome for player #%d=%s", i, outcomeToString(round.Outcomes[i]))	
	}
	
	return round.Outcomes, total_hands_played_this_round
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
		log.Printf("[rounds.go] Round starts. Dealer: %s, Player: %s", round.Dealer, round.Player)
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
		log.Printf("[rounds.go][play()] There is a SPLIT OPPORTUNITY.  We drew the same value cards: %d!", round.Player[0])
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
				log.Printf("[rounds.go] Player stands.")
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
				log.Printf("[rounds.go] Player hits. Hand: %s Total: %d", round.Player, round.Player.Sum())
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
				log.Printf("[rounds.go] Player doubles. Hand: %s Total: %d", round.Player, round.Player.Sum())
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
			log.Printf("[rounds.go] Player busted!")
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
			log.Printf("[rounds.go] Dealer hits. Hand: %s Total: %d", round.Dealer, round.Dealer.Sum())
		}
	}

	// Okay, if the dealer busted, you win. If the dealer is greater, you
	// win.
	if round.Dealer.IsBusted() {
		if verbose {
			log.Printf("[rounds.go] Dealer busted! Hand: %s", round.Dealer)
		}

		return OUTCOME_WIN
	} else if round.Dealer.Sum() > round.Player.Sum() {
		if verbose {
			log.Printf("[rounds.go] Dealer wins. Dealer: %s, Player: %s", round.Dealer, round.Player)
		}

		return OUTCOME_LOSS
	} else if round.Player.Sum() == round.Dealer.Sum() {
		if verbose {
			log.Printf("[rounds.go] Round pushes. Dealer: %s, Player: %s", round.Dealer, round.Player)
		}

		return OUTCOME_PUSH
	}

	if verbose {
		log.Printf("[rounds.go] Player wins! Dealer: %s, Player: %s", round.Dealer, round.Player)
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

func NewRound(deck Deck, num_players int) *Round {
	log.Printf("[rounds.go][NewRound()][num_players=%d]", num_players)
	round := new(Round)
	round.deck = deck
	round.num_players = num_players
	round.Players = make([]Hand, num_players)
	round.Outcomes = make([]Outcome, num_players)
	
	// ^^^ do not modify, working code for multi-player
	
	// phase 2 - code uses an object (as opposed to just a single Hand) per player
	round.PlayersObj = make([]Player, num_players)
	
	// dev - learning (workd)
//	round.PlayersObj[0].Hand = Hand{}
//	round.PlayersObj[0].Hands = make([]Hand, 1)
	
	// dev - learning try 2
	//round.PlayersObj[0].Hands = append(round.PlayersObj[0].Hands, Hand{})
	
	// try 2- works (??)
	for j:=0; j<num_players; j++ {
/*

//work through this -> preferred way
		round.PlayersObj[j] = round.PlayersObj[j].NewPlayer()
		
*/
		
		round.PlayersObj[j].Hands = round.PlayersObj[j].AddHand(Hand{}) 
		round.PlayersObj[j].activeHand = 0
		
		dlog.Always("[rounds.go][NewRound()][player #%d has %d hands, %d h-cap]",
			j, len(round.PlayersObj[j].Hands), cap(round.PlayersObj[j].Hands))
					
	}		
	
	
	//round.PlayersObj[0].AddHand(Hand{}) 
	
	/*
	for i:=0; i<num_players; i++ {
		log.Printf("[rounds.go][NewRound()][create new player #%d]", i)
		player := NewPlayer()
		//round.PlayersObj[i] = player
		round.PlayersObj = append(round.PlayersObj, player)
	}
	*/
	
	return round
}
